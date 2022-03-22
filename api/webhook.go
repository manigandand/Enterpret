package api

import (
	"enterpret/errors"
	"enterpret/ingester"
	"enterpret/response"
	"enterpret/schema"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// generic api handler, process all the webhooks calls
func handleEvent(w http.ResponseWriter, r *http.Request) *errors.AppError {
	slug := chi.URLParam(r, "slug")
	apiKey := chi.URLParam(r, "apiKey")
	version := chi.URLParam(r, "version")

	alertConfig, err := store.AlertConfig().GetBySlug(slug, version)
	if err != nil {
		return err
	}
	integration, err := store.GetIntegrationByAPIKey(apiKey)
	if err != nil {
		return err
	}
	org, err := store.GetOrgByID(integration.OrganizationID)
	if err != nil {
		return err
	}

	ingestion := ingester.NewIngestion(nil, nil, alertConfig)
	headers := make(map[string]string)
	for k := range r.Header {
		headers[k] = r.Header.Get(k)
	}
	ingestion.Set("headers", headers)

	defer r.Body.Close()
	data, ierr := ingestion.IngestFromReq(r)
	if err != nil {
		return ierr
	}

	alert := &schema.Alert{
		CreatedAt:      time.Now(),
		AlertConfigID:  alertConfig.ID,
		IntegrationID:  integration.ID,
		OrganizationID: org.ID,
		Source:         alertConfig.Slug,
		Type:           alertConfig.Type,
		Subject:        data.Subject,
		Message:        data.Message,
		Language:       alertConfig.Language,
		Metadata:       data.Metadata,
		Raw:            data.Raw,
	}
	_, _ = store.Alerts().Save(alert)

	response.Created(w, alert)
	return nil
}
