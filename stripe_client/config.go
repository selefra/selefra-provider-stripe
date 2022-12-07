package stripe_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	APIKey string `yaml:"api_key"  mapstructure:"api_key"`
}
