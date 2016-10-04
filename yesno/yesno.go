package yesno

import (
	"fmt"
	"net/http"

	"github.com/go-chat-bot/bot"
)

func yesno(command *bot.Cmd) (msg string, err error) {
	res, err := http.Get("https://yesno.wtf/api")
	if err != nil {
		return "shit happened", err
	}

	msg = fmt.Sprintf("all signs point to %s", res.Request.URL.String())
	return
}

func init() {
	bot.RegisterCommand(
		"yesno",
		"Returns a random yes/no/maybe gif",
		"",
		yesno)
}
