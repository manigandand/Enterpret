package api

import (
	"enterpret/errors"
	"enterpret/ingester"
	"enterpret/response"
	"net/http"

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

	ingestion := ingester.NewIngestion(org, integration, alertConfig)
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

	response.Created(w, data)
	return nil
}
