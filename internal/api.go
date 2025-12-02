package api

import (
	"net/http"
	"log"
	"encoding/json"
)

func getLocations() []string{

	url := "https://pokeapi.co/api/v2/location-area/"
	
	req, err:=http.NewRequest("GET", url, nil)
	if err!=nil {
		log.Fatal(err)
	}

	client:=http.Client{}
	res, err:= client.Do(req)
	if err!=nil {
		log.Fatal(err)
	}

	data := []string{}
	decoder:=json.NewDecoder(res.Body)
	err=decoder.Decode(&data)
	if err!=nil {
		log.Fatal(err)
	}

	return data

}
