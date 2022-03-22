package api

import (
	"enterpret/errors"
	"enterpret/response"
	"net/http"
)

func getAllAlertSources(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ac, err := store.AlertConfig().All()
	if err != nil {
		return err
	}
	response.OK(w, ac)
	return nil
}

func getAllAlerts(w http.ResponseWriter, r *http.Request) *errors.AppError {
	alerts, err := store.Alerts().All()
	if err != nil {
		return err
	}
	response.OK(w, alerts)
	return nil
}
