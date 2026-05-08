package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/route"
	"github.com/anugrahsputra/portfolio-backend/pkg/logger"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

func initDatabase(cfg *config.Config) *config.Database {
	db, err := config.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Optional: Initialize schema if needed (can be toggled via env or flag)
	if os.Getenv("INIT_SCHEMA") == "true" {
		if err := db.InitSchema("sql/schema/schema.sql"); err != nil {
			log.Errorf("Failed to initialize schema: %v", err)
		}
	}
	return db
}

func initMail(cfg *config.Config) *config.Mail {
	m, err := config.NewMail(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize mailer: %v", err)
	}

	return m
}

func main() {
	// Initialize Logger
	logger.ConfigureLogger()

	// Load Configuration
	cfg := config.Load()

	// Initialize Database
	db := initDatabase(cfg)
	defer db.Close()

	// initialize mailer
	mail := initMail(cfg)

	// Initialize router
	app := route.SetupRouter(db, mail, cfg)

	// Graceful Shutdown
	go func() {
		log.Infof("Server starting on port %s", cfg.Port)
		if err := app.Listen(":" + cfg.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Info("Server exiting")
}

