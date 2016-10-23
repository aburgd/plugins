package about

import (
    "fmt"

    "github.com/go-chat-bot/bot"
)

var email = "comfelliot@gmail.com"
var subjectLine = "[BRIGHT]"
func about(command *bot.Cmd) (msg string, err error) {
    if err != nil {
      return "shit happened", err
    }
    if command.User.Nick == "ghostpage" || command.User.Nick == "TyGently" {
      msg = fmt.Sprintf("if I'm broken, please contact my owner with the subject line %s at %s", subjectLine, email)
    } else {
      msg = fmt.Sprint("sorry, you're not allowed to do that!")
    }
    return
}

func init() {
    bot.RegisterCommand(
        "about",
        "gives information about bot maintenance",
        "",
        about)
}
