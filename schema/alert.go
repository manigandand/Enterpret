package schema

import "time"

type Alert struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`

	AlertConfigID  uint              `json:"alert_config_id"`
	IntegrationID  uint              `json:"integration_id"`
	OrganizationID uint              `json:"organization_id"`
	Source         string            `json:"source"`
	Type           string            `json:"type"`
	Subject        string            `json:"subject"`
	Message        string            `json:"message"`
	Language       string            `json:"language"`
	Metadata       map[string]string `json:"metadata"`
	Raw            interface{}       `json:"raw"`
}
