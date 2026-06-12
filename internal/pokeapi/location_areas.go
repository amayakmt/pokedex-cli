package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type RespShallowLocation struct {
	Next     *string           `json:"next"`
	Previous *string           `json:"previous"`
	Results  []ShallowLocation `json:"results"`
}

type ShallowLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (c *Client) GetLocationAreas(pageURL *string) (RespShallowLocation, error) {
	baseUrl := "https://pokeapi.co/api/v2/location-area"

	var useURL string
	if pageURL == nil {
		useURL = baseUrl
	} else {
		useURL = *pageURL
	}

	// check cache first
	if val, ok := c.cache.Get(useURL); ok {
		var respStruct RespShallowLocation
		err := json.Unmarshal(val, &respStruct)
		return respStruct, err
	}

	// cache miss - make the request
	req, err := http.NewRequest("GET", useURL, nil)
	if err != nil {
		return RespShallowLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocation{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocation{}, err
	}

	var respStruct RespShallowLocation
	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		return RespShallowLocation{}, err
	}

	c.cache.Add(useURL, body)
	return respStruct, nil
}
