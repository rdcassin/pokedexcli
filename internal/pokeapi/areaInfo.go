package pokeapi

type AreaInfoAPIResults struct {
	EncounterMethodsAndRates []EncounterMethodAndRate `json:"encounter_method_rates"`
	GameIndex            	 int                      `json:"game_index"`
	ID                   	 int                      `json:"id"`
	Location             	 Location                 `json:"location"`
	AreaName             	 string                   `json:"name"`
	AreaNameVariations   	 []AreaNameVariation      `json:"names"`
	PokemonEncounters    	 []PokemonEncounter       `json:"pokemon_encounters"`
}

type EncounterMethodAndRate struct {
	Method			 EncounterMethod  				 `json:"encounter_method"`
	VersionsAndRates []EncounterMethodVersionAndRate `json:"version_details"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodVersionAndRate struct {
	Rate    int     	   `json:"rate"`
	Version VersionDetails `json:"version"`
}

type VersionDetails struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type AreaNameVariation struct {
	Language Language `json:"language"`
	AreaName string   `json:"name"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon   				  		   `json:"pokemon"`
	Details []PokemonEncounterAndVersionDetail `json:"version_details"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounterAndVersionDetail struct {
	EncounterDetails []PokemonEncounterDetail `json:"encounter_details"`
	MaxChance		 int                      `json:"max_chance"`
	VersionDetails   VersionDetails  		  `json:"version"`
}

type PokemonEncounterDetail struct {
	Chance          int   				   `json:"chance"`
	ConditionValues []any 				   `json:"condition_values"`
	MaxLevel        int   				   `json:"max_level"`
	EncounterMethod EncounterMethodDetails `json:"method"`
	MinLevel	 	int   				   `json:"min_level"`
}

type EncounterMethodDetails struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}