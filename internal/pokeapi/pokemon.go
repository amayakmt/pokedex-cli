package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Pokemon struct {
	BaseExperience int     `json:"base_experience"`
	Name           string  `json:"name"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

type Stats struct {
	Stat     Stat `json:"stat"`
	BaseStat int  `json:"base_stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type Types struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	// check cache first
	if val, ok := c.cache.Get(url); ok {
		var respStruct Pokemon
		err := json.Unmarshal(val, &respStruct)
		return respStruct, err
	}

	// cache miss - make the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var respStruct Pokemon
	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)
	return respStruct, nil
}
