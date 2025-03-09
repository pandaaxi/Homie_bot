package main

import (
    "log"
    "net/http"
    "sync"

    "github.com/pandaaxi/homie_bot/pkg/bot"
    "github.com/pandaaxi/homie_bot/pkg/api"
    "github.com/pandaaxi/homie_bot/pkg/database"
    "github.com/pandaaxi/homie_bot/pkg/utils/config"
    "github.com/pandaaxi/homie_bot/pkg/utils/logger"
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
