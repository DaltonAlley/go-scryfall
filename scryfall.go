package scryfall

import (
	"net/http"
)

const APIBaseURL = "https://api.scryfall.com/cards/"

func New() *http.Client {
	return &http.Client{}
}
