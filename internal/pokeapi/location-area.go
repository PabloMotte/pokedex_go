package pokeapi

import "encoding/json"


// https://pokeapi.co/docs/v2#location-areas
type LocationAreas struct {
	Count    int    `json:"count"`
	NextUrl     *string `json:"next"`
	PreviousUrl *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// https://pokeapi.co/docs/v2#location-area
type LocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []any `json:"condition_values"`
				PokemonDetails  any   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationAreas(data []byte) (LocationAreas, error) {
	// decode the json data from the response
	// into a new locationAreas struct containing a list of areas
	var location_areas LocationAreas
	err := json.Unmarshal(data, &location_areas)
	if err != nil {
		return LocationAreas{}, err
	}

	return location_areas, nil
}

func GetLocationArea(data []byte) (LocationArea, error) {
	// decode the json data from the response
	// into a new locationArea struct containing details of a single area
	var location_area LocationArea
	err := json.Unmarshal(data, &location_area)
	if err != nil {
		return LocationArea{}, err
	}

	return location_area, nil
}
