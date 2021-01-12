package server

type Config struct {
	// Listen port
	//
	// Default :11111
	Port string
	// Serving folder with images
	//
	// Default '/resources'
	ServeFolder string
}

func NewConfig() Config {
	return Config{
		Port:        ":11111",
		ServeFolder: "/resources",
	}
}
