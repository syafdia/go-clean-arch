package config

type HTTPConfig struct {
	BaseURL string
}

type CronConfig struct {
	// TODO
}

type PresentationConfig struct {
	HTTP HTTPConfig
	Cron CronConfig
}
