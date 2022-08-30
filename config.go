package server

type Config struct {
	hostname string
	port     int
	protocol string
}

var defaultConfig = Config{
	hostname: "0.0.0.0",
	port:     9000,
	protocol: "http",
}
