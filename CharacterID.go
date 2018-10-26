package ffxiv

import (
	"fmt"
	"strings"

	"github.com/aerogo/http/client"
)

// GetCharacterID fetches the first found character ID for the given nick and server.
func GetCharacterID(nick string, server string) (string, error) {
	// Replace spaces with plus signs
	nick = strings.Replace(nick, " ", "+", -1)

	url := fmt.Sprintf("https://na.finalfantasyxiv.com/lodestone/character/?q=%s&worldname=%s", nick, server)
	_, err := client.Get(url).End()

	if err != nil {
		return "", err
	}

	return "", nil
}
