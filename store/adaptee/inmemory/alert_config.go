package inmemory

import (
	"enterpret/errors"
	"enterpret/schema"
)

// AlertConfig implements AlertConfig adapter interface
type AlertConfig struct {
	*Client
}

// NewAlertConfigStore ...
func NewAlertConfigStore(client *Client) *AlertConfig {
	return &AlertConfig{client}
}

func (a *AlertConfig) GetBySlug(slug, version string) (*schema.AlertConfig, *errors.AppError) {
	ac, ok := a.Client.alertConfigs[slug]
	if !ok {
		return nil, errors.NotFound("invalid alert slug")
	}
	if ac.Version != version {
		return nil, errors.BadRequest("unsupported alert version")
	}

	return ac, nil
}

func (a *AlertConfig) All() (map[string]*schema.AlertConfig, *errors.AppError) {
	return a.Client.alertConfigs, nil
}
