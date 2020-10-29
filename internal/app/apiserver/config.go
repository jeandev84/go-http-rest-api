package apiserver


// Config ...
type Config struct {

	// Адрес на котором будет запустится наш веб сервер
	BindAddr string  `toml:"bind_addr"`
	LogLevel string  `toml:"log_level"`
}


// NewConfig
func NewConfig() *Config {
	return &Config {
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}