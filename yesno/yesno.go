package yesno

import (
	"fmt"
  "log"
	"net/http"
  "encoding/json"

	"github.com/go-chat-bot/bot"
)

type yesnoresp struct {
	Answer string `json:"answer"`
	Forced bool `json:"forced"`
	Image string `json:"image"`
}

func yesno(command *bot.Cmd) (msg string, err error) {
  req, err := http.NewRequest("Get", "https://yesno.wtf/api", nil)
  if err != nil {
    return "shit happened", err
  }
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return "shit happened", err
  }
  defer resp.Body.Close()

  var record yesnoresp

  if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
    log.Println(err)
  }

  msg = fmt.Sprintf("all signs point to %s %s", record.Answer, record.Image)
  return
}

func init() {
	bot.RegisterCommand(
		"yesno",
		"Returns a random yes/no/maybe gif",
		"",
		yesno)
}
