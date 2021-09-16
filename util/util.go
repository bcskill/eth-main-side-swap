package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var tgAlerter TgAlerter

type TgAlerter struct {
	BotId  string
	ChatId string
}

func InitTgAlerter(cfg AlertConfig) {
	tgAlerter = TgAlerter{
		BotId:  cfg.TelegramBotId,
		ChatId: cfg.TelegramChatId,
	}
}

type MessageInfo struct {
	MsgType  string `json:"msgtype"`
	Text   TextInfo    `json:"text"`
}

type TextInfo struct {
	Content  string `json:"content"`
	MentionedList []string `json:"mentioned_list"`
	MentionedMobileList    []string `json:"mentioned_mobile_list"`
}

func SendTelegramMessage(msg string) {
	if tgAlerter.BotId == "" || tgAlerter.ChatId == "" || msg == "" {
		return
	}
	msg = fmt.Sprintf("eth-main-side-swap-backend alert: %s", msg)
	endPoint := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", tgAlerter.BotId)

	textInfo := TextInfo{ Content: msg, MentionedList: []string{"@all"}, MentionedMobileList: []string{"@all"}}
	postInfo := MessageInfo {
		MsgType: "text",
		Text: textInfo,
	}
	jsonBytes, err := json.Marshal(postInfo)
	if err != nil {
		Logger.Errorf("send message error=%s",  err)
		return
	}
	Logger.Infof("send tg message, msg=%s",  msg)
	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonBytes))
	if err != nil {
		Logger.Errorf("send message error=%s",  err)
		return
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		Logger.Errorf("send message error=%s",  err)
		return
	}
	resBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Logger.Errorf("send message error=%s",  err)
		return
	}
	Logger.Infof("response, msg=%s",  string(resBytes))

	defer resp.Body.Close()
}
