package inmemory

import (
	"enterpret/errors"
	"enterpret/schema"
)

// Alert implements Alert adapter interface
type Alert struct {
	*Client
}

// NewAlertStore ...
func NewAlertsStore(client *Client) *Alert {
	return &Alert{client}
}

func (a *Alert) All() ([]*schema.Alert, *errors.AppError) {
	return a.Client.alerts, nil
}

func (a *Alert) Save(alert *schema.Alert) (*schema.Alert, *errors.AppError) {
	alert.ID = uint(len(a.Client.alerts) + 1)
	a.Client.alerts = append(a.Client.alerts, alert)
	return alert, nil
}
