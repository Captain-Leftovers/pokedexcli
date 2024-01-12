package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName

	fullURL := baseURL + endpoint

	// check cache here

	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache found

		fmt.Println("cache found :)")

		pokemon := Pokemon{}

		err := json.Unmarshal(data, &pokemon)

		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	fmt.Println("cache not found :( -> -> fetching ... <- <-")

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("pokeapi returned status code %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)

	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}

	err = json.Unmarshal(data, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}
