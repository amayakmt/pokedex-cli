package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type RespLocationArea struct {
	PokemonEncounter []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonName `json:"pokemon"`
}

type PokemonName struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (c *Client) GetLocationArea(name string) (RespLocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + name

	// check cache first
	if val, ok := c.cache.Get(url); ok {
		var respStruct RespLocationArea
		err := json.Unmarshal(val, &respStruct)
		return respStruct, err
	}

	// cache miss - make the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	var respStruct RespLocationArea
	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, body)
	return respStruct, nil
}
