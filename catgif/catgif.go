package catgif

import (
	"github.com/go-chat-bot/bot"
	"net/http"
  "fmt"
)

func gif(command *bot.Cmd) (msg string, err error) {
	res, err := http.Get("http://thecatapi.com/api/images/get?format=src&type=gif")
	if err != nil {
		return "", err
	}
  msg = fmt.Sprintf("here's your freaking cat gif: %s", res.Request.URL.String())
	return
}

func init() {
	bot.RegisterCommand(
		"catgif",
		"Returns a random cat gif.",
		"",
		gif)
}
