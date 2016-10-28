package support

import (
  "fmt"

  "github.com/go-chat-bot/bot"
)

var kofiLink = "http://ko-fi.com/A642DU2"

func support(command *bot.Cmd) (msg string, err error) {
  if err != nil {
    return "shit happened", err
  }
  msg = fmt.Sprintf("Please support my creator, the almighty eng13, by buying them a coffee: %s", kofiLink)

  return
}

func init() {
  bot.RegisterCommand(
    "support",
    "tell user how to support creator",
    "",
    support)
}
