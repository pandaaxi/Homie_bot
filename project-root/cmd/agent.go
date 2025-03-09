package main

import (
	"flag"
	"log"
	"time"

	"project-root/pkg/vnstat"
	// Use net/http or other packages to send the data
)

func main() {
	token := flag.String("t", "", "Secret token for agent registration")
	serverURL := flag.String("n", "", "Main bot server URL")
	flag.Parse()

	if *token == "" || *serverURL == "" {
		log.Fatal("Token and server URL must be provided")
	}

	log.Printf("Agent started with token: %s and server: %s", *token, *serverURL)

	// Main loop: read vnStat data and send to the main server
	for {
		usage, err := vnstat.GetUsage()
		if err != nil {
			log.Printf("Error reading vnStat data: %v", err)
		} else {
			log.Printf("Usage data: %+v", usage)
			// TODO: Send usage data to the server via HTTPS API (e.g., HTTP POST)
		}
		time.Sleep(5 * time.Minute)
	}
}
