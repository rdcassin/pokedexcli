package pokeapi

type AreaInfoAPIResults struct {
	EncounterMethodsAndRates []EncounterMethodAndRate `json:"encounter_method_rates"`
	GameIndex            	 int                      `json:"game_index"`
	ID                   	 int                      `json:"id"`
	Location             	 Location                 `json:"location"`
	AreaName             	 string                   `json:"name"`
	AreaNameVariations   	 []AreaNameVariation      `json:"names"`
	PokemonEncounters    	 []PokemonEncounter      `json:"pokemon_encounters"`
}

type EncounterMethodAndRate struct {
	EncounterMethod 					EncounterMethod  `json:"encounter_method"`
	EncounterMethodsVersionsAndRates  	[]EncounterMethodsVersionAndRate `json:"version_details"`
}

type EncounterMethod struct {
	MethodName string `json:"name"`
	MethodURL  string `json:"url"`
}

type EncounterMethodsVersionAndRate struct {
	EncounterMethodRate    int     	   	  `json:"rate"`
	EncounterMethodVersion VersionDetails `json:"version"`
}

type VersionDetails struct {
	EncounterMethodVersionName string `json:"name"`
	EncounterMethodVersionURL  string `json:"url"`
}

type Location struct {
	LocationName string `json:"name"`
	LocationURL  string `json:"url"`
}

type AreaNameVariation struct {
	Language LanguageDetails `json:"language"`
	AreaName     string   `json:"name"`
}

type LanguageDetails struct {
	LanguageName string `json:"name"`
	LanguageURL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon        				   	  PokemonDetails   					  `json:"pokemon"`
	PokemonEncounterAndVersionDetails []PokemonEncounterAndVersionDetail  `json:"version_details"`
}

type PokemonDetails struct {
	PokemonName string `json:"name"`
	PokemonURL  string `json:"url"`
}

type PokemonEncounterAndVersionDetail struct {
	PokemonEncounterDetails 			[]PokemonEncounterDetail 		`json:"encounter_details"`
	MaxChance		  					int                             `json:"max_chance"`
	PokemoneEncounterVersionDetails     VersionDetails  				`json:"version"`
}

type PokemonEncounterDetail struct {
	Chance          int   				   `json:"chance"`
	ConditionValues []any 				   `json:"condition_values"`
	MaxLevel        int   				   `json:"max_level"`
	EncounterMethod EncounterMethodDetails `json:"method"`
	MinLevel	 	int   				   `json:"min_level"`
}

type EncounterMethodDetails struct {
	EncounterMethodName string `json:"name"`
	EncounterMethodURL  string `json:"url"`
}