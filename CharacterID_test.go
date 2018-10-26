package ffxiv_test

import (
	"testing"

	"github.com/animenotifier/ffxiv"
	"github.com/stretchr/testify/assert"
)

func TestGetCharacterID(t *testing.T) {
	id, err := ffxiv.GetCharacterID("Aky Otara", "Asura")

	assert.NoError(t, err)
	assert.Equal(t, "9015414", id)
}

func TestGetCharacterIDFail(t *testing.T) {
	id, err := ffxiv.GetCharacterID("asdfasdfghjklasdf", "Asura")

	assert.Error(t, err)
	assert.Empty(t, id)
}
