package main

import (
	"log"
	"net/http"
	"sync"

	"project-root/pkg/bot"
	"project-root/pkg/api"
	"project-root/pkg/database"
	"project-root/pkg/utils/config"
	"project-root/pkg/utils/logger"
)

func main() {
	// Load configuration and initialize logger
	cfg := config.LoadConfig()
	logger.Init(cfg.LogLevel)

	// Initialize SQLite database
	err := database.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Use WaitGroup to run both bot and API server concurrently
	var wg sync.WaitGroup

	// Start Telegram bot in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		bot.StartBot(cfg.TelegramToken)
	}()

	// Start HTTPS API server for agent communication
	wg.Add(1)
	go func() {
		defer wg.Done()
		api.StartServer(cfg.APIAddr, cfg.TLSCert, cfg.TLSKey)
	}()

	wg.Wait()
}
