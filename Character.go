package ffxiv

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aerogo/http/client"
)

var levelRegEx = regexp.MustCompile(`level\s(\d+)`)

// Character represents a Final Fantasy XIV character.
type Character struct {
	Level int
}

// GetCharacter fetches character data for a given character ID.
func GetCharacter(id string) (*Character, error) {
	// Fetch the page
	url := fmt.Sprintf("https://na.finalfantasyxiv.com/lodestone/character/%s", id)
	response, err := client.Get(url).End()

	if err != nil {
		return nil, err
	}

	page := response.Bytes()
	reader := bytes.NewReader(page)
	document, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		return nil, err
	}

	character := &Character{}
	characterClassData := document.Find(".character__class__data")

	if characterClassData.Length() == 0 {
		return nil, errors.New("Error parsing character class data")
	}

	levelInfo := strings.ToLower(characterClassData.Text())
	matches := levelRegEx.FindStringSubmatch(levelInfo)

	if len(matches) >= 2 {
		levelText := matches[1]
		level, err := strconv.Atoi(levelText)

		if err != nil {
			return nil, err
		}

		character.Level = level
	}

	return character, nil
}
