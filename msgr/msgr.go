package msgr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	ChatID           string `json:"chat_id"`
	ReplyToMessageID string `json:"reply_to_message_id"`
}

type TextMessage struct {
	Message
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type VideoMessage struct {
	Message
	VideoID string `json:"video"`
}

type TgMsgr struct {
	BotToken string
}

func (msgr *TgMsgr) SendTextMessage(message TextMessage) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", msgr.BotToken)

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("message failed to send: %s", body)
	}

	return nil
}

func (msgr *TgMsgr) SendVideoMessage(message VideoMessage) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendVideo", msgr.BotToken)

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("message failed to send: %s", body)
	}

	return nil
}
