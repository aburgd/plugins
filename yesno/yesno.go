package yesno

import (
	"fmt"
  "log"
	"net/http"
  "encoding/json"

	"github.com/go-chat-bot/bot"
)

type yesnorecord struct {
	Answer string `json:"answer"`
	Forced bool `json:"forced"`
	Image string `json:"image"`
}

func yesno(command *bot.Cmd) (msg string, err error) {
  req, err := http.NewRequest("GET", "https://yesno.wtf/api", nil)
  if err != nil {
    return "shit happened", err
  }
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return "shit happened", err
  }
  defer resp.Body.Close()

  var record yesnorecord
  msg = fmt.Sprintf("%+v", record)

  if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
    log.Println(err)
  }

  msg = fmt.Sprintf("all signs point to %s", record.Image)
  return
}

func init() {
	bot.RegisterCommand(
		"yesno",
		"Returns a random yes/no/maybe gif",
		"",
		yesno)
}
