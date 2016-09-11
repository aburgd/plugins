package botgod

import (
  "github.com/go-chat-bot/bot"
  "regexp"
  "math/rand"
)

const (
  pattern = "((I|i) (am|AM)|(eng13|ENG13|eng|ENG) (is|IS)|((you|YOU|You)|(they|They|THEY)) (are|ARE)) ((the|THE)|(a|A)) (botgod|bot god|BOTGOD|BOT GOD)"
)

var (
  re = regexp.MustCompile(pattern)
)

func botgod(command *bot.PassiveCmd) (string, error) {
  issuer := command.User.Nick
  possible := []string{
    "you've gotta be shitting me " + issuer,
    issuer + " i swear to myself",
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
