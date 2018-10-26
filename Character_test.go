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
}

func TestGetCharacterFail(t *testing.T) {
	id := "404"
	character, err := ffxiv.GetCharacter(id)

	assert.Error(t, err)
	assert.Nil(t, character)
}
