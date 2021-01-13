package server

var (
	DefaultPort       = ":11111"
	DefaultSaveFolder = "./resources"
	DefaultWorkers    = 5
)

type Config struct {
	// Listen port
	//
	// Default :11111
	Port string
	// Serving folder with images
	//
	// Default '/resources'
	ServeFolder string
	// The number of workers who download content in parallel
	//
	// Default 5
	ParallelWorkers int
}

func NewConfig() Config {
	return Config{
		Port:            DefaultPort,
		ServeFolder:     DefaultSaveFolder,
		ParallelWorkers: DefaultWorkers,
	}
}
