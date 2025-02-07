package Bot

import (
	"MyBot/consts"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type TGBot struct {
	URL url.URL
}
type GetMeResponse struct {
	OK     bool `json:"ok"`
	Result User `json:"result"`
}
type GetUpdatesResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}
type User struct {
	ID       int    `json:"id"`
	IsBot    bool   `json:"is_bot"`
	UserName string `json:"username"`
}
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}
type BotMessage struct {
	MessageID int    `json:"message_id"`
	Text      string `json:"text"`
}
type Message struct {
	MessageID int    `json:"message_id"`
	Text      string `json:"text"`
	From      User   `json:"from"`
	Chat      Chat   `json:"chat"`
}
type Chat struct {
	ID int `json:"id"`
}

func (t *TGBot) GetMe() error {
	getMePath, err := url.JoinPath(t.URL.String(), consts.MethodGetMe)
	resp, err := http.Get(getMePath)
	if err != nil {
		return consts.CantReachBot
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return consts.CantReachBot
	}

	var getMeResponse GetMeResponse

	if err := json.Unmarshal(body, &getMeResponse); err != nil {
		return consts.CantReachBot
	}

	fmt.Printf("\nName:%v\n", getMeResponse.Result.UserName)

	return nil
}
func (t *TGBot) GetUpdates(offset int) ([]Update, error) {
	getUpdatesPath, err := url.JoinPath(t.URL.String(), consts.MethodGetUpdates)

	query := url.Values{}
	if offset != 0 {
		query.Add("offset", strconv.Itoa(offset))
	}

	req, err := http.NewRequest(http.MethodGet, getUpdatesPath, nil)
	if err != nil {
		return nil, consts.CantGetUpdates
	}

	var client http.Client
	req.URL.RawQuery = query.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, consts.CantGetUpdates
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, consts.CantReachBot
	}

	var getUpdatesResponse GetUpdatesResponse
	if err := json.Unmarshal(body, &getUpdatesResponse); err != nil {
		return nil, consts.CantGetUpdates
	}

	return getUpdatesResponse.Result, nil
}
func (t *TGBot) SendMessage(chatID int, text string) (Message, error) {
	sendMessagePath, err := url.JoinPath(t.URL.String(), consts.MethodSendMessage)
	if err != nil {
		return Message{}, err
	}

	query := url.Values{}
	query.Add("chat_id", strconv.Itoa(chatID))
	query.Add("text", text)

	req, err := http.NewRequest(http.MethodGet, sendMessagePath, nil)
	if err != nil {
		return Message{}, err
	}

	var client http.Client
	req.URL.RawQuery = query.Encode()
	res, err := client.Do(req)
	if err != nil {
		return Message{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Message{}, err
	}

	var message Message
	if err := json.Unmarshal(body, &message); err != nil {
		return Message{}, err
	}

	return message, err
}

func New(token string) TGBot {
	return TGBot{
		URL: url.URL{
			Scheme: "https",
			Host:   consts.TgBotHost,
			Path:   "bot" + token,
		},
	}
}
func PrintUpdates(updates []Update) int {
	var res int
	if len(updates) == 0 {
		log.Printf("Message: no message")
		return 0
	}
	for _, update := range updates {
		log.Printf("Message: %s\n", update.Message.Text)
		res = update.UpdateID
	}
	return res + 1
}
