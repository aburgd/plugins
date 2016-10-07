package duckduckgo

import (
  "github.com/go-chat-bot/bot"
  "github.com/pkg/errors"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
  "strings"
	// "net/url"
)

type ddgrecord struct {
	Abstract string `json:"Abstract"`
	AbstractSource string `json:"AbstractSource"`
	AbstractText string `json:"AbstractText"`
	AbstractURL string `json:"AbstractURL"`
	Answer int `json:"Answer"`
	AnswerType string `json:"AnswerType"`
	Definition string `json:"Definition"`
	DefinitionSource string `json:"DefinitionSource"`
	DefinitionURL string `json:"DefinitionURL"`
	Heading string `json:"Heading"`
	Image string `json:"Image"`
	Redirect string `json:"Redirect"`
	RelatedTopics []struct {
		FirstURL string `json:"FirstURL"`
		Icon struct {
			Height int `json:"Height"`
			URL string `json:"URL"`
			Width int `json:"Width"`
		} `json:"Icon"`
		Result string `json:"Result"`
		Text string `json:"Text"`
	} `json:"RelatedTopics"`
	Results []struct {
		FirstURL string `json:"FirstURL"`
		Icon struct {
			Height int `json:"Height"`
			URL string `json:"URL"`
			Width int `json:"Width"`
		} `json:"Icon"`
		Result string `json:"Result"`
		Text string `json:"Text"`
	} `json:"Results"`
	Type string `json:"Type"`
}

func duckgo(command *bot.Cmd) (msg string, err error) {
  args := command.Args[0:]
  query := strings.Join(args, " ")
	url := fmt.Sprintf("http://api.duckduckgo.com/?q=" + query + "&format=json&pretty=1")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = errors.Wrap(err, "HTTP/GET failed")
		return "problem", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "Do failed")
		return "problem", err
	}
	// defer resp.Body.Close()

	var record ddgrecord

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	msg = fmt.Sprintf("here's the gist: %s", record.AbstractURL)
  defer resp.Body.Close()
	return
}

func init() {
	bot.RegisterCommand(
		"ddg",
		"queries a search on DuckDuckGo",
		"",
		duckgo)
}
