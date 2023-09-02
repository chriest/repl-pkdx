package pokeapi

import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
)

func (c *Client) ListAreas(URL *string) (LocationsFromJSON, error) {

	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if URL != nil{
		fullURL = *URL
	}

	data, o := c.cache.Get(fullURL)
	if o {
		
		locationFromJSON := LocationsFromJSON{}
		err := json.Unmarshal(data, &locationFromJSON)
		if err!= nil {
			return LocationsFromJSON{}, err
		}
		//fmt.Println("hit")
		return locationFromJSON, err

	}
	//fmt.Println("miss")
	// richiesta
	req, err := http.NewRequest("GET", fullURL, nil)

	if err!= nil {
		return LocationsFromJSON{}, err
	}
	// esecuzione richiesta
	res, e := c.httpClient.Do(req)

	if e!= nil {
		return LocationsFromJSON{}, err
	}
	// controllo risposta
	 if res.StatusCode > 299 {
		return LocationsFromJSON{}, fmt.Errorf("error: %v", res.StatusCode)
	 }

	 reqData, err := io.ReadAll(res.Body)
	 defer res.Body.Close()

	 if err!= nil {
		return LocationsFromJSON{}, err
	}

	// init struttura derivata da GO>JSON
	locationFromJSON := LocationsFromJSON{}

	err = json.Unmarshal(reqData, &locationFromJSON)
	if err!= nil {
		return LocationsFromJSON{}, err
	}
	c.cache.Add(fullURL, reqData)

	return locationFromJSON, err
}

func (c *Client) GetAreas(locationName string) (LocationAreas, error) {

	endpoint := "/location-area/" + locationName
	fullURL := baseURL + endpoint

	data, o := c.cache.Get(fullURL)
	if o {
		
		locationAreas := LocationAreas{}
		err := json.Unmarshal(data, &locationAreas)
		if err!= nil {
			return LocationAreas{}, err
		}
		//fmt.Println("hit")
		return locationAreas, err

	}
	//fmt.Println("miss")
	// richiesta
	req, err := http.NewRequest("GET", fullURL, nil)

	if err!= nil {
		return LocationAreas{}, err
	}
	// esecuzione richiesta
	res, e := c.httpClient.Do(req)

	if e!= nil {
		return LocationAreas{}, err
	}
	// controllo risposta
	 if res.StatusCode > 299 {
		return LocationAreas{}, fmt.Errorf("error: %v", res.StatusCode)
	 }

	 reqData, err := io.ReadAll(res.Body)
	 defer res.Body.Close()

	 if err!= nil {
		return LocationAreas{}, err
	}

	// init struttura derivata da GO>JSON
	locationAreas := LocationAreas{}

	err = json.Unmarshal(reqData, &locationAreas)
	if err!= nil {
		return LocationAreas{}, err
	}
	c.cache.Add(fullURL, reqData)

	return locationAreas, err
}