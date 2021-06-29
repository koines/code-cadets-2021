package pokemon

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
)

const pokemonData = "https://pokeapi.co/api/v2/pokemon/"

type LocationArea struct {
	Name string
}

type LocationAreaEncounter struct {
	Location LocationArea `json:"location_area""`
}

type Pokemon struct {
	Name string `json:"name"`
}

type Output struct {
	Name      string
	Locations []string
}

func GetData(url string) ([]byte, error) {
	httpClient := pester.New()

	httpResponse, err := httpClient.Get(url)
	if err != nil {
		return nil, errors.WithMessage(err, "HTTP get towards pokemon API")
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "reading body of pokemon API response")
	}

	return bodyContent, nil
}

func FindPokemonLocations(input string) error {
	bodyContent, error := GetData(pokemonData + input)
	if error != nil {
		return error
	}

	var decodedContent Pokemon
	err := json.Unmarshal(bodyContent, &decodedContent)
	if err != nil {
		return errors.WithMessage(err, "unmarshal: json content")
	}

	bodyContent, error = GetData(pokemonData + input + "/encounters")
	if err != nil {
		return err
	}

	var areas []LocationAreaEncounter
	err = json.Unmarshal(bodyContent, &areas)
	if err != nil {
		return errors.WithMessage(err, "unmarshal: json content")
	}

	var output Output
	output.Name = decodedContent.Name
	for _, value := range areas {
		output.Locations = append(output.Locations, value.Location.Name)
	}

	result, _ := json.MarshalIndent(output, "", "\t")
	fmt.Println(string(result))

	return nil
}
