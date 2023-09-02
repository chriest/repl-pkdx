package pokeapi

import (
	"net/http"
)

func (c *Client) ListAreas() (LocationsFromJSON, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)

	if err!= nil {
		return LocationsFromJSON{}, err

	}

	res, e := 
}