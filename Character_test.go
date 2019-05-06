package ffxiv_test

import (
	"testing"

	"github.com/animenotifier/ffxiv"
	"github.com/stretchr/testify/assert"
)

func TestGetCharacter(t *testing.T) {
	id := "9015414"
	character, err := ffxiv.GetCharacter(id)

	assert.NoError(t, err)
	assert.NotNil(t, character)
	assert.NotZero(t, character.Level)
	assert.Equal(t, "Aky Otara", character.Nick)
	assert.Equal(t, "Asura", character.Server)
	assert.Equal(t, "Mana", character.DataCenter)
}

func TestGetCharacterFail(t *testing.T) {
	id := "404"
	character, err := ffxiv.GetCharacter(id)

	assert.Error(t, err)
	assert.Nil(t, character)
}
