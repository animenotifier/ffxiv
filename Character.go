package ffxiv

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/aerogo/http/client"
)

var digit = regexp.MustCompile("[0-9]+")

// Character represents a Final Fantasy XIV character.
type Character struct {
	Nick       string
	Server     string
	DataCenter string
	Class      string
	Level      int
	ItemLevel  int
}

// GetCharacter fetches character data for a given character ID.
func GetCharacter(id string) (*Character, error) {
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

	characterName := document.Find(".frame__chara__name").Text()

	if characterName == "" {
		return nil, errors.New("Error parsing character name")
	}

	// This will look like: "Asura (Mana)"
	characterServerAndDataCenter := document.Find(".frame__chara__world").Text()

	if characterServerAndDataCenter == "" {
		return nil, errors.New("Error parsing character server")
	}

	// Normalize whitespace characters
	characterServerAndDataCenter = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return ' '
		}

		return r
	}, characterServerAndDataCenter)

	// Split the server and data center
	serverInfo := strings.Split(characterServerAndDataCenter, " ")

	if len(serverInfo) < 2 {
		return nil, errors.New("Character server info does not seem to include the data center")
	}

	characterServer := serverInfo[0]
	characterDataCenter := serverInfo[1]
	characterDataCenter = strings.TrimPrefix(characterDataCenter, "(")
	characterDataCenter = strings.TrimSuffix(characterDataCenter, ")")

	if characterDataCenter == "" {
		return nil, errors.New("Error parsing character data center")
	}

	// Level
	characterLevel := document.Find(".character__class__data").Text()

	if characterLevel == "" {
		return nil, errors.New("Error parsing character level")
	}

	// Weapon
	characterWeapon := document.Find(".db-tooltip__item__category").Text()

	if characterWeapon == "" {
		return nil, errors.New("Error parsing character class")
	}

	level, err := strconv.Atoi(digit.FindStringSubmatch(characterLevel)[0])

	if err != nil {
		return nil, err
	}

	itemLevel := calculateItemLevel(document)
	className := getJobName(document)

	if className == "" {
		// Determine class name using the weapon info
		className = strings.Split(characterWeapon, "'")[0]
		className = strings.Replace(className, "Two-handed", "", -1)
		className = strings.Replace(className, "One-handed", "", -1)
		className = strings.TrimSpace(className)
	}

	character := &Character{
		Nick:       characterName,
		Class:      className,
		Server:     characterServer,
		DataCenter: characterDataCenter,
		Level:      level,
		ItemLevel:  itemLevel,
	}

	return character, nil
}

// calculateItemLevel will try to return the average of all item levels.
func calculateItemLevel(document *goquery.Document) int {
	items := document.Find(".item_detail_box")

	itemCount := 0
	itemLevelSum := 0

	items.Each(func(i int, item *goquery.Selection) {
		itemCategory := strings.ToLower(item.Find(".db-tooltip__item__category").Text())

		// Ignore soul crystals
		if itemCategory == "soul crystal" {
			return
		}

		// Two-handed weapons are counted twice
		weight := 1

		if strings.Contains(itemCategory, "two-handed") {
			weight = 2
		}

		// Find item level
		itemLevelText := item.Find(".db-tooltip__item__level").Text()
		itemLevelText = digit.FindStringSubmatch(itemLevelText)[0]
		itemLevel, err := strconv.Atoi(itemLevelText)

		if err != nil {
			return
		}

		itemLevelSum += itemLevel * weight
		itemCount += weight
	})

	if itemCount == 0 {
		return 0
	}

	return itemLevelSum / itemCount
}

// getJobName finds the job name by looking at the soul crystal text.
func getJobName(document *goquery.Document) string {
	const soulCrystalPrefix = "Soul of the "

	jobName := ""
	itemNames := document.Find(".db-tooltip__item__name")

	itemNames.EachWithBreak(func(i int, item *goquery.Selection) bool {
		itemName := item.Text()

		if strings.HasPrefix(itemName, soulCrystalPrefix) {
			jobName = itemName[len(soulCrystalPrefix):]
			return false
		}

		return true
	})

	return jobName
}
