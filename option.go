package server

type OptionFunc func(cfg *Config)

func HostnameOpt(hostname string) OptionFunc {
	return func(cfg *Config) {
		cfg.hostname = hostname
	}
}

func PortOpt(port int) OptionFunc {
	return func(cfg *Config) {
		cfg.port = port
	}
}

func ProtocolOpt(name string) OptionFunc {
	return func(cfg *Config) {
		cfg.protocol = name
	}
}
