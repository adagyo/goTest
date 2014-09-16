package utils

type Config struct {
	// Mongo URL
	MgoURL string

	// Database name
	MgoDB string
}

func LoadConfig(conf *Config) {
	conf.MgoURL = "localhost"
	conf.MgoDB	= "myApi"
}
