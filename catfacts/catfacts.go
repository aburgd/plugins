package catfacts

import (
	"fmt"
	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
	"regexp"
)

const (
	pattern   = "(?i)\\b(cat|gato|miau|meow|garfield|lolcat|missekat|pusekatt|kissekatt|kisu|kisa|mjau|mja|mjá|nau|mau|kitty)[s|z]{0,1}\\b"
	msgPrefix = "here's a damn fact about cats: %s"
)

type facts struct {
	Facts   []string `json:"facts"`
	Success string   `json:"success"`
}

var (
	re          = regexp.MustCompile(pattern)
	catFactsURL = "http://catfacts-api.appspot.com/api/facts?number=1"
)

func catFacts(command *bot.PassiveCmd) (string, error) {
	if !re.MatchString(command.Raw) {
		return "", nil
	}
	data := &facts{}
	err := web.GetJSON(catFactsURL, data)
	if err != nil {
		return "", err
	}

	if len(data.Facts) == 0 {
		return "", nil
	}

	return fmt.Sprintf(msgPrefix, data.Facts[0]), nil
}

func init() {
	bot.RegisterPassiveCommand(
		"catfacts",
		catFacts)
}
