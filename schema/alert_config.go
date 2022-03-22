package schema

// AlertConfig - Alert schema containts all the config to parse the alert data
type AlertConfig struct {
	ID         uint              `json:"id,omitempty" yaml:"-"`
	Version    string            `json:"version" yaml:"version"`
	Name       string            `json:"name" yaml:"name"`
	Slug       string            `json:"slug" yaml:"slug"`
	Type       string            `json:"type" yaml:"type"`
	SupportDoc string            `json:"support_doc" yaml:"supportDoc"`
	IsValid    bool              `json:"is_valid" yaml:"isValid"`
	Subject    string            `json:"subject" yaml:"subject"`
	Message    string            `json:"message" yaml:"message"`
	Language   string            `json:"language" yaml:"language"`
	Metadata   map[string]string `json:"metadata" yaml:"metadata"`
	// IsBatched       bool              `json:"is_batched" yaml:"isBatched"`
	// ArraySelector   string            `json:"array_selector" yaml:"arraySelector"`
}
