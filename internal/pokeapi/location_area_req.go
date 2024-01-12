package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"

	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	// check cache here

	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache found

		fmt.Println("cache found :)")

		locationAreasResp := LocationAreasResp{}

		err := json.Unmarshal(data, &locationAreasResp)

		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}

	fmt.Println("cache not found :( -> -> fetching ... <- <-")

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasResp{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("pokeapi returned status code %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)

	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}

	err = json.Unmarshal(data, &locationAreasResp)

	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName

	fullURL := baseURL + endpoint

	// check cache here

	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache found

		fmt.Println("cache found :)")

		locationArea := LocationArea{}

		err := json.Unmarshal(data, &locationArea)

		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	fmt.Println("cache not found :( -> -> fetching ... <- <-")

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationArea{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("pokeapi returned status code %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)

	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}

	err = json.Unmarshal(data, &locationArea)

	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationArea, nil
}
