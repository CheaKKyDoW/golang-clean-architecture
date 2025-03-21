package main

import (
	"golang-clean-architecture/internal/infrastructure/server"
	"golang-clean-architecture/pkg/config"
	"log"
)

var (
	version string
	build   string
)

// @title Open API Server
// @version 0.0.1
// @description Open API Server
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
//
// @BasePath /api/v1
func main() {

	cfg, err := config.LoadEnv(".env")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting app (version: %s build: %s, env: %s)", version, build, cfg.App.ENV)

	if cfg.App.Debug {
		log.Printf("config: %+v", cfg)
	}

	srv, err := server.NewServer(&cfg)

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(srv.Run())
}
