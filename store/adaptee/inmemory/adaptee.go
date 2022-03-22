package inmemory

import (
	"enterpret/schema"
	"enterpret/store/adapter"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

// NewAdapter returns store inmemory adapter(*Client)
func NewAdapter() adapter.Store {
	// Load Data
	c := &Client{
		alertConfigs: make(map[string]*schema.AlertConfig),
	}
	c.AlertConfigConn = NewAlertConfigStore(c)
	c.AlertsConn = NewAlertsStore(c)

	c.loadAlertConfigs()

	return c
}

func (c *Client) loadAlertConfigs() {
	if err := filepath.Walk("./manifest", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".yaml" {
			// parse yaml
			b, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal("can't read yaml file " + err.Error())
			}
			var alertConfig schema.AlertConfig
			if err := yaml.Unmarshal(b, &alertConfig); err != nil {
				log.Fatal("can't parse yaml file " + err.Error())
			}
			alertConfig.ID = uint(len(c.alertConfigs) + 1)
			c.alertConfigs[alertConfig.Slug] = &alertConfig
		}
		return nil
	}); err != nil {
		panic("can't seed data " + err.Error())
	}

	c.organization = &schema.Organization{
		ID:           1,
		Name:         "Enterpret Inc",
		Slug:         "enterpret",
		Subscription: "enterprise-monthly",
	}

	apiKey := uuid.New().String()
	c.integration = &schema.Integration{
		ID:             1,
		Name:           "Enterpret Push Integrations",
		Slug:           "enterpret-intg",
		OrganizationID: 1,
		AlertConfigID:  1,
		APIKey:         apiKey,
	}

	log.Println("Available Alert Configs Test Endpoints:")
	for slug, ac := range c.alertConfigs {
		log.Printf("/%s/alert/%s/%s\n", ac.Version, slug, apiKey)
	}
}
