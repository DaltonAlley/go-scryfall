package scryfall

import (
	"net/http"
)

const APIBaseURL = "https://api.scryfall.com/cards/"

type Scryfall struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Scryfall {
	return &Scryfall{
		baseURL:    APIBaseURL,
		httpClient: httpClient,
	}
}
