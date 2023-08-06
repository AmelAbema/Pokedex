package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseUrl + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	do, err := c.httpClient.Do(request)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(do.Body)
	if do.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", do.StatusCode)
	}
	all, err := io.ReadAll(do.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(all, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}
