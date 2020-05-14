package config

// Configuration is the exported object containing all configurations
type Configuration struct {
	API APIConfiguration
}

type APIConfiguration struct {
	Port string
}
