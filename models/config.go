package models

type Config struct {
	Server serverConfig
}

type serverConfig struct {
	IP       string
	Port     int
	TickRate int
	Slots    int
	CORS     bool
	Debug    bool
}
