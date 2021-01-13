package server

import (
	"ResourceServer/service"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app           *fiber.App
	port          string
	servingFolder string
	service       service.Service
}

// Нужно доделать фичу чтоб можно было искать
// изображение только по названию без расщирения
func (s *Server) InitRoutes() error {
	s.app.Use(cors.New())
	s.app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))

	s.app.Static("/", s.loadServingRoot())

	s.app.Get("/docs/*", swagger.Handler)
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
