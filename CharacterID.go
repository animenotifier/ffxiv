package ffxiv

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/aerogo/http/client"
)

// GetCharacterID fetches the first found character ID for the given nick and server.
func GetCharacterID(nick string, server string) (string, error) {
	// Replace spaces with plus signs
	nick = strings.Replace(nick, " ", "+", -1)

	// Fetch the page
	url := fmt.Sprintf("https://na.finalfantasyxiv.com/lodestone/character/?q=%s&worldname=%s", nick, server)
	response, err := client.Get(url).End()

	if err != nil {
		return "", err
	}

	page := response.Bytes()
	reader := bytes.NewReader(page)
	document, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		return "", err
	}

	href := document.Find(".entry__link").First().AttrOr("href", "")

	if !strings.HasPrefix(href, "/lodestone/character/") {
		return "", nil
	}

	id := strings.TrimPrefix(href, "/lodestone/character/")
	id = strings.TrimSuffix(id, "/")

	return id, nil
}
