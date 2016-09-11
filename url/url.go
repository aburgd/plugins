package url

import (
	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
	"html"
	"net/url"
	"regexp"
	"strings"
  "log"
)

const (
	minDomainLength = 3
)

var (
	re = regexp.MustCompile("<title>\\n*?(.*?)\\n*?<\\/title>")
)

func canBeURLWithoutProtocol(text string) bool {
  matched, err := regexp.MatchString("[a-z0-9]\\.(com|edu|gov|net|org|co|me)", text)
  if err != nil{
    log.Fatal("shit happened ", err)
  }

  return len(text) > minDomainLength &&
    !strings.HasPrefix(text, "http") &&
    matched
}

func extractURL(text string) string {
	extractedURL := ""
	for _, value := range strings.Split(text, " ") {
		if canBeURLWithoutProtocol(value) {
			value = "http://" + value
		}

		parsedURL, err := url.Parse(value)
		if err != nil {
			continue
		}
		if strings.HasPrefix(parsedURL.Scheme, "http") {
			extractedURL = parsedURL.String()
			break
		}
	}
	return extractedURL
}

func urlTitle(cmd *bot.PassiveCmd) (string, error) {
	URL := extractURL(cmd.Raw)

	if URL == "" {
		return "", nil
	}

	body, err := web.GetBody(URL)
	if err != nil {
		return "", err
	}

	title := re.FindString(string(body))
	if title == "" {
		return "", nil
	}

	title = strings.Replace(title, "\n", "", -1)
	title = title[strings.Index(title, ">")+1 : strings.LastIndex(title, "<")]

	return html.UnescapeString(title), nil
}

func init() {
	bot.RegisterPassiveCommand(
		"url",
		urlTitle)
}
