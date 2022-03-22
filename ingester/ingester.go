package ingester

import (
	"bytes"
	"encoding/json"
	"enterpret/errors"
	"enterpret/schema"
	"enterpret/types"

	// "errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"

	"github.com/google/uuid"
)

// Ingester contains all the variables and templates for ingesting the incoming JSON data
type Ingester struct {
	renderTemplate *template.Template
	delimiters     Delims
	integration    *schema.Integration
	organization   *schema.Organization

	logger *log.Logger

	variables map[string]interface{}

	*schema.AlertConfig
}

var baseRenderTempl = template.New("renderTemplate")

// NewIngestion creates a new ingestion object
func NewIngestion(org *schema.Organization, integ *schema.Integration, ac *schema.AlertConfig) *Ingester {
	s := &Ingester{
		renderTemplate: template.Must(baseRenderTempl.Clone()),
		organization:   org,
		delimiters: Delims{
			Left:  DefaultDelimiters.Left,
			Right: DefaultDelimiters.Right,
		},
		integration: integ,
		logger:      log.New(os.Stdout, fmt.Sprintf("[%s] ", uuid.New().String()), log.Lshortfile),

		variables: make(map[string]interface{}),

		AlertConfig: ac,
	}

	return s
}

func (i *Ingester) Set(name string, val interface{}) {
	i.variables[name] = val
}

// IngestFromReq ingests the incoming http.Request(JSON) data
func (i *Ingester) IngestFromReq(r *http.Request) (*types.Ingester, *errors.AppError) {
	body, err := i.decodeBody(r)
	if err != nil {
		return nil, err
	}

	return i.Ingest(body), nil
}

func (i *Ingester) Ingest(data interface{}) *types.Ingester {
	message := strings.TrimSpace(i.parseTmpl(i.Message, data))
	subject := strings.TrimSpace(i.parseTmpl(i.Subject, data))

	meta := make(map[string]string)
	for k, v := range i.Metadata {
		meta[k] = strings.TrimSpace(i.parseTmpl(v, data))
	}

	return &types.Ingester{
		Subject:  subject,
		Message:  message,
		Metadata: meta,
		Raw:      data,
	}
}

// IngestFromJSON ingests the incoming JSON data
func (i *Ingester) decodeBody(r *http.Request) (interface{}, *errors.AppError) {
	i.logger.Println(r.RequestURI)
	contentType := strings.ToLower(r.Header.Get("Content-Type"))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.BadRequest("error while reading the request body")
	}
	i.logger.Println("Incoming Payload: ", string(body))

	switch contentType {
	case "application/json":
		return i.decodeBodyJSON(body)
	case "application/x-www-form-urlencoded":
		return i.decodeFormData(body)
	default:
		return nil, errors.BadRequest("unsupported content-type")
	}
}

func (i *Ingester) decodeBodyJSON(body []byte) (interface{}, *errors.AppError) {
	var req interface{}

	body = bytes.TrimSpace(body)
	if len(body) < 1 {
		return nil, errors.BadRequest("error while readin the request body string")
	}
	switch body[0] {
	case '[': // rune 91 "["
		// req = []map[string]interface{}{}
		return nil, errors.BadRequest("json array not supported")
	case '{': // rune 123 "{"
		req = types.JSON{}
	default:
		return nil, errors.BadRequest("invalid request type not an array/object")
	}
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, errors.BadRequest("invalid request body " + err.Error())
	}

	return req, nil
}

func (i *Ingester) decodeFormData(body []byte) (interface{}, *errors.AppError) {
	req := types.JSON{}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, errors.BadRequest("error while parsing url query " + err.Error())
	}

	for key, value := range values {
		if len(value) == 1 {
			req[key] = value[0]
			continue
		}
		req[key] = value
	}

	return req, nil
}

// parseTemplate parses the template and returns the rendered string
func (i *Ingester) parseTmpl(gotmpl string, data interface{}) string {
	enclose := i.delimiters.Enclose

	gohtmlTemplate := enclose("with .data") + gotmpl + enclose("end")
	// i.logger.Println("Parsing template 1: ", gohtmlTemplate)

	tmpl := template.Must(i.renderTemplate.Clone())
	tmpl.Delims(i.delimiters.Left, i.delimiters.Right)
	tmpl.Funcs(template.FuncMap{
		"tojson": func(v interface{}) string {
			b, err := json.Marshal(v)
			if err != nil {
				return ""
			}
			return string(b)
		},
	})
	tmpl = template.Must(tmpl.Parse(gohtmlTemplate))

	word := bytes.NewBuffer(nil)
	if err := tmpl.Execute(word, types.JSON{
		"data":      data,
		"variables": i.variables,
	}); err != nil {
		i.logger.Println("Error while parsing template: ", err)
	}

	return word.String()
}
