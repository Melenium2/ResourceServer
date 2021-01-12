package server

import (
	"ResourceServer/service"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app           *fiber.App
	port          string
	servingFolder string
	service       service.Service
}

func (s *Server) InitRoutes() error {
	s.app.Static("/", s.servingFolder)
	s.app.Get("/load", s.loadRoute)
	s.app.Post("/load/batch", s.loadBatchRoute)

	return nil
}

func (s *Server) Listen() error {
	return s.app.Listen(s.port)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}

func New(app *fiber.App, service service.Service, config ...Config) *Server {
	var conf Config
	if len(config) > 0 {
		conf = config[0]
	}

	return &Server{
		app:           app,
		port:          conf.Port,
		servingFolder: conf.ServeFolder,
		service:       service,
	}
}
