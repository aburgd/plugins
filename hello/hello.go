package hello

import (
    "fmt"
    "github.com/go-chat-bot/bot"
)

func hello(command *bot.Cmd) (msg string, err error) {
    if err != nil {
      return "shit happened", err
    }
    msg = fmt.Sprintf("what the hell do you want?")
    return
}

func init() {
    bot.RegisterCommand(
        "hi",
        "Sends a 'Hello' message to you on the channel.",
        "",
        hello)
}
