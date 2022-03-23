package inmemory

import (
	"enterpret/errors"
	"enterpret/schema"
	"enterpret/store/adapter"
)

// Client struct implements the store adapter interface
type Client struct {
	organization    *schema.Organization
	integration     *schema.Integration
	alertConfigs    map[string]*schema.AlertConfig
	alerts          []*schema.Alert
	AlertConfigConn adapter.AlertConfig
	AlertsConn      adapter.Alerts
}

func (c *Client) Close() {
	c.alertConfigs = nil
}

func (c *Client) GetOrgByID(id uint) (*schema.Organization, *errors.AppError) {
	if c.organization.ID != id {
		return nil, errors.NotFound("invalid organization id")
	}
	return c.organization, nil
}

func (c *Client) GetIntegrationByAPIKey(apiKey string) (*schema.Integration, *errors.AppError) {
	if c.integration.APIKey != apiKey {
		return nil, errors.NotFound("invalid integration api key")
	}
	return c.integration, nil
}

func (c *Client) GetIntegrationByID(id uint) (*schema.Integration, *errors.AppError) {
	if c.integration.ID != id {
		return nil, errors.NotFound("invalid integration id")
	}
	return c.integration, nil
}

// Topic ...
func (c *Client) AlertConfig() adapter.AlertConfig {
	return c.AlertConfigConn
}

// Alerts ...
func (c *Client) Alerts() adapter.Alerts {
	return c.AlertsConn
}
