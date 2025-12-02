package pokeapi

import (
	"net/http"
	"encoding/json"
)

type LocationAreaResponse struct {
	Results []struct {
		Name string
		URL string
	}
	Next *string
	Previous *string
}

func GetLocations(pageURL *string) (LocationAreaResponse, error) {

	url  := "https://pokeapi.co/api/v2/location-area/"

	if pageURL != nil {
		url = *pageURL
	}
	
	req, err:=http.NewRequest("GET", url, nil)
	if err!=nil {
		return LocationAreaResponse{}, err
	}

	client:=http.Client{}
	res, err:= client.Do(req)
	if err!=nil {
		return LocationAreaResponse{}, err
	}

	var data LocationAreaResponse
	decoder:=json.NewDecoder(res.Body)
	err=decoder.Decode(&data)
	if err!=nil {
		return LocationAreaResponse{}, err
	}


	return data, nil

}
