package server

import (
	"fmt"
	"golang-clean-architecture/internal/infrastructure/client"
	"golang-clean-architecture/internal/infrastructure/database"
	"golang-clean-architecture/internal/infrastructure/validate"
	"golang-clean-architecture/internal/infrastructure/websocket"

	"golang-clean-architecture/pkg/config"
	pkg_model "golang-clean-architecture/pkg/model"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	// swagger
	// _ "golang-clean-architecture/docs"
)

type Resource struct {
	Cfg         *config.Cfg
	DBConn      *gorm.DB
	RedisClient *redis.Client
	App         *fiber.App
	HTTPClient  *client.HTTPClient
	Validator   *validate.Validator
	WSManager   *websocket.WebSocketManager
}

func NewServer(cfg *config.Cfg) (resource *Resource, err error) {

	resource = &Resource{
		Cfg:       cfg,
		Validator: validate.NewValidator(validator.WithRequiredStructEnabled()),
	}

	if cfg.App.Debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	resource.DBConn, err = database.NewPostgresDB(cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Database.SearchPath, cfg.Database.SSLMode, cfg.Database.TZ)
	if err != nil {
		return
	}
	if cfg.App.Debug {
		resource.DBConn = resource.DBConn.Debug()
	}

	return
}

func (r *Resource) ErrorHandler(c *fiber.Ctx, err error) error {
	respForm := pkg_model.Response{
		Header:      "Internal Server Error",
		Description: err.Error(),
	}
	return c.Status(http.StatusInternalServerError).JSON(respForm)
}

func (r *Resource) Run() (err error) {

	// init app
	r.App = fiber.New(fiber.Config{
		AppName:      r.Cfg.App.Name,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		// support 8k header size
		ReadBufferSize: 8 * 1024,
		BodyLimit:      4 * 1024 * 1024,
		// https://docs.gofiber.io/guide/faster-fiber
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		// if multiple cpus can prefork
		// Prefork:      runtime.NumCPU() > 1,
		ErrorHandler: r.ErrorHandler,
	})

	// Use global middlewares.
	r.App.Use(cors.New())
	r.App.Use(helmet.New())
	r.App.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe:     r.Liveness,
		LivenessEndpoint:  "/livez",
		ReadinessProbe:    r.Readiness,
		ReadinessEndpoint: "/readyz",
	}))
	r.App.Use(logger.New())
	r.App.Use(recover.New())
	r.App.Use(requestid.New(
		requestid.Config{
			ContextKey: "requestID",
		},
	))
	r.App.Use(limiter.New(limiter.Config{
		Max:               r.Cfg.App.RateLimitConn,
		Expiration:        r.Cfg.App.RateLimitWindow,
		KeyGenerator:      func(c *fiber.Ctx) string { return c.IP() },
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, slow down duckyou.",
			})
		},
	},
	))
	// ✅ Global WebSocket Upgrade Check

	r.App.Get("/swagger/*", swagger.HandlerDefault)

	// init module by go:build tags
	r.InitModule()

	// listen in other go routine
	go func() {
		addr := fmt.Sprintf("%s:%d", r.Cfg.App.Host, r.Cfg.App.Port)
		if err := r.App.Listen(addr); err != nil {
			log.Panic(err)
		}
	}()

	// channel to receive os signal
	closeCh := make(chan os.Signal, 1)
	signal.Notify(closeCh, os.Interrupt, syscall.SIGTERM)

	// block here and wait for os signal
	<-closeCh
	log.Println("Gracefully shutting down...")
	_ = r.App.Shutdown()

	log.Println("Running cleanup tasks...")

	// cleanup tasks
	if r.RedisClient != nil {
		log.Println("Closing redis connection...")
		if err = r.RedisClient.Close(); err != nil {
			log.Println("Redis connection was unable to closed.")
			log.Println(err)
		}
		log.Println("Redis connection was successful closed.")
	}

	if r.DBConn != nil {
		if db, err := r.DBConn.DB(); err == nil {
			log.Println("Closing database connection...")
			if err = db.Close(); err != nil {
				log.Println("Database connection was unable to closed.")
				log.Println(err)
			}
			log.Println("Database connection was successful closed.")
		}
	}
	log.Println("Fiber was successful shutdown.")
	return
}

func (r *Resource) Liveness(_ *fiber.Ctx) (live bool) {
	return true
}

func (r *Resource) Readiness(c *fiber.Ctx) (ready bool) {
	wg := &sync.WaitGroup{}

	var redisReady, dbReady bool

	// Check Redis if exist
	if r.RedisClient != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			redisReady = r.RedisClient.Ping(c.Context()).Err() == nil
		}()
	} else {
		redisReady = true
	}

	// Check DB
	if r.DBConn == nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if sqlDB, err := r.DBConn.DB(); err == nil {
				dbReady = sqlDB.PingContext(c.Context()) == nil
			}
		}()
	} else {
		dbReady = true
	}

	wg.Wait()

	ready = redisReady && dbReady
	return
}
