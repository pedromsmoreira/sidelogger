package loggers

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Server struct {
	app *fiber.App
	wg  sync.WaitGroup
	log Logger
}

func NewServer(log Logger) *Server {
	// configure middlewares
	// configure routes
	app := fiber.New(fiber.Config{
		ReadTimeout: 5 * time.Second,
	})
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Sidelogger!")
	})

	app.Get("/dashboard", monitor.New())

	return &Server{
		app: app,
		wg:  sync.WaitGroup{},
		log: log,
	}
}

func (s *Server) Start(address string, port string) error {
	s.wg.Add(1)
	var err error
	go func() {
		defer s.wg.Done()

		err := s.app.Listen(fmt.Sprintf("%v:%v", address, port))

		if err != nil {
			fmt.Errorf("API Server stopped listening due to %v", err)
		}
	}()

	return err
}

func (s *Server) Shutdown() error {
	err := s.app.Shutdown()

	if err != nil {
		s.wg.Wait()
	}

	return err
}
