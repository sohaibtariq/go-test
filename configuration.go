package swaggerpetstoreopenapi30

import (
    "github.com/apimatic/go-core-runtime/https"
    "os"
)

type ConfigurationOptions func(*Configuration)

type Configuration struct {
    environment       Environment
    httpConfiguration https.HttpConfiguration
    apiKey            string
}

func newConfiguration(options ...ConfigurationOptions) Configuration {
    config := Configuration{}
    
    for _, option := range options {
        option(&config)
    }
    return config
}

func WithEnvironment(environment Environment) ConfigurationOptions {
    return func(c *Configuration) {
        c.environment = environment
    }
}

func WithHttpConfiguration(httpConfiguration https.HttpConfiguration) ConfigurationOptions {
    return func(c *Configuration) {
        c.httpConfiguration = httpConfiguration
    }
}

func WithApiKey(apiKey string) ConfigurationOptions {
    return func(c *Configuration) {
        c.apiKey = apiKey
    }
}

func (c *Configuration) Environment() Environment {
    return c.environment
}

func (c *Configuration) HttpConfiguration() https.HttpConfiguration {
    return c.httpConfiguration
}

func (c *Configuration) ApiKey() string {
    return c.apiKey
}

func (c *Configuration) LoadFromEnvironment() {
    environment := os.Getenv("SWAGGERPETSTOREOPENAPI_30_ENVIRONMENT")
    if environment != "" {
        c.environment = Environment(environment)
    }
    apiKey := os.Getenv("SWAGGERPETSTOREOPENAPI_30_API_KEY")
    if apiKey != "" {
        c.apiKey = apiKey
    }
}

// Available Servers
type Server string

const (
    ENUMDEFAULT Server = "default"
)

// Available Environments
type Environment string

const (
    PRODUCTION Environment = "production"
)

func RetryConfigurationFactory(options ...https.RetryConfigurationOptions) https.RetryConfiguration {
    config := DefaultRetryConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}

func HttpConfigurationFactory(options ...https.HttpConfigurationOptions) https.HttpConfiguration {
    config := DefaultHttpConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}

func ConfigurationFactory(options ...ConfigurationOptions) Configuration {
    config := DefaultConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}
