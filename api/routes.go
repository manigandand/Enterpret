package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Routes - all the registered routes
func Routes(router chi.Router) {
	router.Get("/", IndexHandeler)
	router.Get("/top", HealthHandeler)
	router.Route("/", InitV1Routes)
}

func InitV1Routes(r chi.Router) {
	r.Method(http.MethodPost, "/{version}/alert/{slug}/{apiKey}", Handler(handleEvent))
	r.Method(http.MethodGet, "/v1/alert-sources", Handler(getAllAlertSources))
	r.Method(http.MethodGet, "/v1/alerts", Handler(getAllAlerts))
}
