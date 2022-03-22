package config

import "log"

const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

var (
	Env                  string
	Port                 string
	APIHost              string
	ServerRecipeEndpoint string
	DefaultAPIHost       = "http://localhost:8080"
)

// Initialize initializes all the env variables for this package.
// This function should be called only once during the application.
//
// Example:
//     LoadFromJSON("config.json", "dash.json")
//     AddEnvEntry("HOST", "localhost", &host)
//     AddEnvEntry("PORT", "8080", &port)
//     ...
//     config.Initialize()
func Initialize(files ...string) {
	if initDone {
		panic("config initialization done already")
	}
	LoadFromJSON(files...)

	addNewEnvEntry("ENV", &Env, EnvDevelopment)
	addNewEnvEntry("PORT", &Port, "8080")
	addNewEnvEntry("API_HOST", &APIHost, DefaultAPIHost)

	// load all the env variables. Must be called at the end.
	load()
	log.Println("Inited config...👍")
}
