package adapter

import (
	"enterpret/errors"
	"enterpret/schema"
)

type Store interface {
	Close()
	GetOrgByID(id uint) (*schema.Organization, *errors.AppError)
	GetIntegrationByAPIKey(apiKey string) (*schema.Integration, *errors.AppError)
	GetIntegrationByID(id uint) (*schema.Integration, *errors.AppError)
	AlertConfig() AlertConfig
	Alerts() Alerts
}

type AlertConfig interface {
	GetBySlug(slug, version string) (*schema.AlertConfig, *errors.AppError)
	All() (map[string]*schema.AlertConfig, *errors.AppError)
}

type Alerts interface {
	All() ([]*schema.Alert, *errors.AppError)
	Save(alert *schema.Alert) (*schema.Alert, *errors.AppError)
}
