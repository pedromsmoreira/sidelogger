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
	ls  *LoggerService
}

func NewServer(ls *LoggerService) *Server {
	app := fiber.New(fiber.Config{
		ReadTimeout: 5 * time.Second,
	})
	app.Use(cors.New())
	v1 := app.Group("/v1")
	v1.Get("/dashboard", monitor.New())
	v1.Post("/logs", func(c *fiber.Ctx) error {
		req := new(LogRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		if err := ls.Log(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
			"success": true,
		})
	})

	return &Server{
		app: app,
		wg:  sync.WaitGroup{},
		ls:  ls,
	}
}

func (s *Server) Start(port string) error {
	s.wg.Add(1)
	var err error
	go func() {
		defer s.wg.Done()
		err := s.app.Listen(fmt.Sprintf(":%v", port))

		if err != nil {
			s.ls.SimpleLog(fmt.Sprintf("API Server stopped listening due to %v", err))
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
