# Homie Bot

A Telegram bot-based VPS monitoring and control service.

## Features

- **VPS Monitoring & Control**: Manage VPS instances via Telegram commands.
- **Bandwidth Usage Reporting**: Uses vnStat for tracking data usage.
- **Agent Management**: Deploy lightweight agents on VPS nodes.
- **HTTPS API Communication**: Secure communication using TLS.
- **SQLite Storage**: Local storage for VPS, usage logs, and alerts.

## Setup

1. **Environment Variables**:  
   Set the following variables (or update the defaults in `pkg/utils/config.go`):
   - `TELEGRAM_TOKEN`
   - `DB_PATH`
   - `API_ADDR`
   - `TLS_CERT`
   - `TLS_KEY`
   - `LOG_LEVEL`
   - `AGENT_SECRET_TOKEN` (for API authentication)

2. **Build and Run**:
   ```bash
   go build -o homie_bot ./cmd/main.go
   ./homie_bot
