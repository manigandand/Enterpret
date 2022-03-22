package response

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"net/http"
)

// Response holds the handlerfunc response
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta Meta        `json:"meta"`
}

// Meta holds the status of the request informations
type Meta struct {
	Status  int    `json:"status_code"`
	Message string `json:"error_message,omitempty"`
}

func (r *Response) Send(w http.ResponseWriter) {
	gz := gzip.NewWriter(w)
	defer gz.Close()
	buf, err := json.Marshal(r)
	if err != nil {
		// Fail(w, )
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(r.Meta.Status)
	if r.Meta.Status != http.StatusNoContent {
		if _, err := gz.Write(buf); err != nil {
			log.Println("respond.With.error: ", err)
		}
	}
}
