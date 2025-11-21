package scryfall

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindCardByID(t *testing.T) {
	resp, err := http.Get("https://api.scryfall.com/cards/random")
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var randomCard Card
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&randomCard))

	foundCard, err := FindCardByID(randomCard.ID)
	require.NoError(t, err)

	assert.Equal(t, randomCard.ID, foundCard.ID, "Expected ID %s, got %s", randomCard.ID, foundCard.ID)
	assert.Equal(t, randomCard.Name, foundCard.Name, "Expected Name %s, got %s", randomCard.Name, foundCard.Name)
	assert.Equal(t, randomCard.TypeLine, foundCard.TypeLine, "Expected TypeLine %s, got %s", randomCard.TypeLine, foundCard.TypeLine)
}
