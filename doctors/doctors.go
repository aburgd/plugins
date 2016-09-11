package doctors

import (
  "fmt"
  "github.com/go-chat-bot/bot"
)

const hub string = "http://www.scp-wiki.net/doctors-of-the-church-hub"

func doctors(command *bot.Cmd) (msg string, err error) {
  if err != nil {
    return "shit happened", err
  }

  msg = fmt.Sprintf("ah shit, you're one of them? goddammit fine. here: %s",
    hub)
  return
}

func init() {
  bot.RegisterCommand(
    "doctors",
    "sends the url for doctors of the church canon hub",
    "",
    doctors)
}
