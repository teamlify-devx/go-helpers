package server

import (
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	mw_logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	cfg "github.com/spf13/viper"
	"os"
	"time"
)

func (s *server) NewHttpServer() (*fiber.App, error) {
	s.fiber = fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: false,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ServerHeader:  cfg.GetString("Server.SERVER_HEADER"),
		AppName:       cfg.GetString("Server.PROJECT_NAME") + " " + cfg.GetString("Server.API_VER"),
		Immutable:     true,
	})

	s.fiber.Use(cors.New())

	if cfg.GetBool("Server.ENABLE_PROFILER") == true {
		s.fiber.Use(pprof.New())
	}

	if cfg.GetBool("Server.ENABLE_METRICS") == true {
		s.fiber.Get("/metrics", monitor.New(monitor.Config{Title: cfg.GetString("Server.PROJECT_NAME")}))
	}

	if cfg.GetBool("ENABLE_LOGGER") == true {
		s.fiber.Use(mw_logger.New())
	}

	if cfg.GetBool("ENABLE_RATE_LIMIT") == true {
		s.fiber.Use(limiter.New(limiter.Config{
			Max:                cfg.GetInt("Server.RATE_LIMIT_MAX"),
			Expiration:         time.Duration(cfg.GetInt("Server.RATE_LIMIT_EXP")) * time.Second,
			SkipFailedRequests: true,
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.Get("x-forwarded-for")
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.SendFile("./toofast.html")
			},
		}))
	}

	if cfg.GetBool("ENABLE_DOCS") == true {
		s.fiber.Get("/doc/*", swagger.HandlerDefault)
	}

	s.fiber.Get("/", func(c *fiber.Ctx) error {
		s.logger.Infof("Health check !")
		return c.SendString("Everything is OK ! ;)")
	})

	return s.fiber, nil
}

func (s *server) StartListen() (err error) {
	port := os.Getenv("PORT")

	// If port not init set to default
	if port == "" {
		port = "8080"
		s.logger.Infof("Defaulting to port %s", port)
	}

	URI := fmt.Sprintf("%s:%s", "", port)
	if err = s.fiber.Listen(URI); err != nil {
		s.logger.Fatalf("Error starting server : %s", err.Error())
	}

	s.logger.Infof("Server is listening on PORT : %s", os.Getenv(port))
	return
}
