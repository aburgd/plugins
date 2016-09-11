package botgod

import (
  "github.com/go-chat-bot/bot"
  "regexp"
  "math/rand"
)

const (
  pattern = "(I am|eng13 is|you are)(the|a)(botgod)|(bot god)"
)

var (
  re = regexp.MustCompile(pattern)
)

func botgod(command *bot.PassiveCmd) (string, error) {
  issuer := command.User.Nick
  possible := []string{
    "you've gotta be shitting me " + issuer,
    issuer + "i swear to myself",
    "do you have a church? hm, " + issuer + "?",
  }

  if re.MatchString(command.Raw) {
    return possible[rand.Intn(len(possible))], nil
  }
  return "", nil
}

func init()  {
  bot.RegisterPassiveCommand(
    "botgod",
    botgod)
}
