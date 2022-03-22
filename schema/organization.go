package schema

type Organization struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Subscription string `json:"subscription"`
	Deleted      bool   `json:"-"`
}

// Integration struct contains the information about the integration of the organization
type Integration struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	OrganizationID uint   `json:"organization_id"`
	AlertConfigID  uint   `json:"alert_config_id"`
	APIKey         string `json:"api_key"`
	Deleted        bool   `json:"-"`
}
