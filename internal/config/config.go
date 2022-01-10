package config

// Cfg is the global config variable.
//nolint:gochecknoglobals
var Cfg = Config{
	APIPrefix: "/api",
}

type Config struct {
	APIPrefix string
}
