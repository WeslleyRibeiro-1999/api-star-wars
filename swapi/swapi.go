package swapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetPlanetsSW() PlanetasSW {
	response, err := http.Get("https://swapi.dev/api/planets/")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject PlanetasSW
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

func GetApparances(planetName string) int {
	planets := GetPlanetsSW()
	for _, value := range planets.Results {
		if value.Name == planetName {
			return len(value.Films)
		}
	}

	return 0
}
