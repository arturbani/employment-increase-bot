package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type telegramUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
}

type telegramSticker struct {
	ID string `json:"file_id"`
}

type telegramMessage struct {
	Chat struct {
		ID int64 `json:"id"`
	} `json:"chat"`
	NewChatMembers []telegramUser `json:"new_chat_members"`
}

type sendStickerReqBody struct {
	ChatID  int64  `json:"chat_id"`
	Sticker string `json:"sticker"`
}

type webhookReqBody struct {
	Message telegramMessage
}

func welcomeNewUser(chatID int64) error {
	employmentIncreaseSticker := "CAACAgIAAxkBAAM8X4y1bGIIshhAyJFnt-KydBdC0o8AAnwBAALhNeMIeh1Xz2U4YBQbBA"
	reqBody := &sendStickerReqBody{
		ChatID:  chatID,
		Sticker: employmentIncreaseSticker,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendSticker", os.Getenv("TELEGRAM_TOKEN"))
	res, err := http.Post(telegramURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

func handler(res http.ResponseWriter, req *http.Request) {
	body := &webhookReqBody{}

	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	if len(body.Message.NewChatMembers) == 0 {
		return
	}

	if err := welcomeNewUser(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	fmt.Println("reply sent")
}

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(handler))
}
