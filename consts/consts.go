package consts

import "errors"

const (
	TgBotHost           = "api.telegram.org"
	MethodGetMe         = "getMe"
	MethodGetUpdates    = "getUpdates"
	MethodSendMessage   = "sendMessage"
	MethodDeleteMessage = "deleteMessage"

	MsgStart = "/start"
	MsgHelp  = "/help"
	MsgRnd   = "/rnd"
	MsgRm    = "/rm"
)

var CantReachBot = errors.New("can't reach bot")
var CantGetUpdates = errors.New("can't get updates")
