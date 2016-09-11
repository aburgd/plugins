package about

import (
    "fmt"

    "github.com/go-chat-bot/bot"
)

var skip string = "http://www.scp-wiki.net/scp-963"

func about(command *bot.Cmd) (msg string, err error) {
    if err != nil {
      return "shit happened", err
    }

    msg = fmt.Sprintf("if you haven't already, here's the damn article %s", skip)
    return
}

func init() {
    bot.RegisterCommand(
        "about",
        "sends the url for jack bright's amulet",
        "",
        about)
}
