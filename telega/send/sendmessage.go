package telega

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const (
	chatID       = -1001739414659
	BOT_TOKEN    = "5874445302:AAFmH15Qo6ggFhB_a1bBWAoh-41FPzwZJGI"
	TELEGRAM_URL = "https://api.telegram.org/bot"
)

type BotSendMessage struct {
	Result struct {
		Message_id int
	}
}

func SendMessage(text string) {
	textAll := fmt.Sprintf("%s", text)
	data := []byte(fmt.Sprintf(`{"chat_id":%d , "text":"%s", "parse_mode":"HTML", "disable_web_page_previev"}`, chatID, textAll))

	tx := bytes.NewReader(data)
	_, err := http.Post(fmt.Sprintf("%s%s/sendMessage", TELEGRAM_URL, BOT_TOKEN), "application/json", tx)
	if err != nil {
		log.Fatal(err)
	}
}
