package ffxiv_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/animenotifier/ffxiv"
)

func TestGetCharacter(t *testing.T) {
	id := "9015414"
	character, err := ffxiv.GetCharacter(id)

	assert.Nil(t, err)
	assert.NotNil(t, character)
	assert.NotEqual(t, character.Level, 0)
	assert.Equal(t, "Aky Otara", character.Nick)
	assert.Equal(t, "Asura", character.Server)
	assert.Equal(t, "Mana", character.DataCenter)
}

func TestGetCharacterFail(t *testing.T) {
	id := "404"
	character, err := ffxiv.GetCharacter(id)

	assert.NotNil(t, err)
	assert.Nil(t, character)
}
