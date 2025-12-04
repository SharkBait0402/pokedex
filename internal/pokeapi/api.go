package pokeapi

import (
	"net/http"
	"encoding/json"
	"github.com/sharkbait0402/pokedex/internal/pokecache"
)

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

type LocationAreaResponse struct {
	Results []struct {
		Name string
		URL string
	}
	Next *string
	Previous *string
}

func (c *Client) GetLocations(pageURL *string) (LocationAreaResponse, error) {

	url  := "https://pokeapi.co/api/v2/location-area/"

	if pageURL != nil {
		url = *pageURL
	}

	if info, ok:=c.cache.Get(url); ok {
		var newInfo LocationAreaResponse
		err:=json.Unmarshal(info, &newInfo)
		if err!=nil {
			return LocationAreaResponse{}, err
		}
		return newInfo, nil
	}

	req, err:=http.NewRequest("GET", url, nil)
	if err!=nil {
		return LocationAreaResponse{}, err
	}

	res, err:= c.httpClient.Do(req)
	if err!=nil {
		return LocationAreaResponse{}, err
	}

	var data LocationAreaResponse
	decoder:=json.NewDecoder(res.Body)
	err=decoder.Decode(&data)
	if err!=nil {
		return LocationAreaResponse{}, err
	}

	newByte, err:=json.Marshal(data)
	if err!=nil{
		return LocationAreaResponse{}, err
	}

	c.cache.Add(url, newByte)

	return data, nil

}
