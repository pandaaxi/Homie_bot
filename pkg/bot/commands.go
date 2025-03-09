package bot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleUpdate processes incoming Telegram updates and dispatches commands.
func HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	text := update.Message.Text
	chatID := update.Message.Chat.ID

	switch {
	case strings.HasPrefix(text, "/VPS"):
		handleVPS(chatID)
	case strings.HasPrefix(text, "/Add_VPS"):
		handleAddVPS(chatID)
	case strings.HasPrefix(text, "/RM_VPS"):
		handleRMVPS(chatID)
	case strings.HasPrefix(text, "/Total"):
		handleTotal(chatID)
	case strings.HasPrefix(text, "/Alert"):
		handleAlert(chatID)
	default:
		msg := tgbotapi.NewMessage(chatID, "Unknown command")
		BotAPI.Send(msg)
	}
}

func handleVPS(chatID int64) {
	// TODO: Fetch VPS details from database and format the message accordingly.
	message := "VPS List:\nRemark - IP:Region - Provider - Price($) - Data Reset(Times) - Available Data/Usage - Reset Date"
	msg := tgbotapi.NewMessage(chatID, message)
	BotAPI.Send(msg)
}

func handleAddVPS(chatID int64) {
	// TODO: Collect VPS details from user, then generate an agent install command.
	token := "GENERATED_SECRET_TOKEN"
	installCmd := "./agent -t " + token + " -n https://your-main-bot-domain"
	message := "Add VPS:\nRun the following command on your VPS:\n" + installCmd
	msg := tgbotapi.NewMessage(chatID, message)
	BotAPI.Send(msg)
}

func handleRMVPS(chatID int64) {
	// TODO: List all VPS with an option to delete each.
	message := "Remove VPS:\n[Feature not fully implemented]"
	msg := tgbotapi.NewMessage(chatID, message)
	BotAPI.Send(msg)
}

func handleTotal(chatID int64) {
	// TODO: Calculate and display overall data consumption.
	message := "Total Data Consumption:\nLast Month: 0GB\nCurrent Month: 0GB\nYesterday: 0GB\nToday: 0GB"
	msg := tgbotapi.NewMessage(chatID, message)
	BotAPI.Send(msg)
}

func handleAlert(chatID int64) {
	// TODO: Set or display alert thresholds.
	message := "Alert configuration feature is under development."
	msg := tgbotapi.NewMessage(chatID, message)
	BotAPI.Send(msg)
}
