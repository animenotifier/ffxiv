package ffxiv_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/animenotifier/ffxiv"
)

func TestGetCharacterID(t *testing.T) {
	id, err := ffxiv.GetCharacterID("Aky Otara", "Asura")

	assert.Nil(t, err)
	assert.Equal(t, "9015414", id)
}

func TestGetCharacterIDFail(t *testing.T) {
	id, err := ffxiv.GetCharacterID("asdfasdfghjklasdf", "Asura")

	assert.NotNil(t, err)
	assert.Equal(t, id, "")
}
