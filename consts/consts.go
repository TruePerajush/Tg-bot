package consts

import "errors"

const (
	TgBotHost         = "api.telegram.org"
	MethodGetMe       = "getMe"
	MethodGetUpdates  = "getUpdates"
	MethodSendMessage = "sendMessage"

	msgStart = "/start"
	msgHelp  = "/help"
	msgRnd   = "/rnd"
	msgRm    = "/rm"
)

var CantReachBot = errors.New("can't reach bot")
var CantGetUpdates = errors.New("can't get updates")
