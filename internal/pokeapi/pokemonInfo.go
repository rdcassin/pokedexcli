package pokeapi

type PokemonInfoAPIResults struct {
	Abilities                 []Ability     `json:"abilities"`
	BaseExperience            int           `json:"base_experience"`
	Cries                     CryLinks      `json:"cries"`
	Forms                     []Form        `json:"forms"`
	GameIndices               []GameIndex   `json:"game_indices"`
	Height                    int           `json:"height"`
	HeldItems              	  []HeldItem    `json:"held_items"`
	ID                        int           `json:"id"`
	IsDefault                 bool          `json:"is_default"`
	LocationAreaEncountersURL string        `json:"location_area_encounters"`
	Moves                     []Move        `json:"moves"`
	Name                      string        `json:"name"`
	Order                     int           `json:"order"`
	PastAbilities             []PastAbility `json:"past_abilities"`
	PastTypes                 []PastType    `json:"past_types"`
	Species                   Species       `json:"species"`
	Sprites                   Sprites       `json:"sprites"`
	Stats                     []Stat       `json:"stats"`
	Types                     []Type       `json:"types"`
	Weight                    int           `json:"weight"`
}

type Ability struct {
	Details  AbilityDetails `json:"ability"`
	IsHidden bool    		`json:"is_hidden"`
	Slot     int	 		`json:"slot"`
}

type AbilityDetails struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type CryLinks struct {
	LatestURL string `json:"latest"`
	LegacyURL string `json:"legacy"`
}

type Form struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type GameIndex struct {
	IndexNum	   int      	  `json:"game_index"`
	VersionDetails VersionDetails `json:"version"`
}

// VersionDetails struct defined in areaInfo.go

type HeldItem struct {
	Item     Item      `json:"item"`
	Versions []Version `json:"version_details"`
}

type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Version struct {
	Rarity  int     	   `json:"rarity"`
	Version VersionDetails `json:"version"`
}

type Move struct {
	Details         	MoveDetails          `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type MoveDetails struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int             	`json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethod 	`json:"move_learn_method"`
	Order           int             	`json:"order"`
	VersionGroup    VersionGroupDetails `json:"version_group"`
}

type MoveLearnMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionGroupDetails struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PastAbility struct {
	Abilities  []Ability  `json:"abilities"`
	Generation Generation `json:"generation"`
}

type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PastType struct {
	Generation Generation `json:"generation"`
	Types      []Type    `json:"types"`
}

type Type struct {
	Slot 	int  		`json:"slot"`
	Details TypeDetails `json:"type"`
}

type TypeDetails struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Sprites struct {
	BackDefaultURL   string   `json:"back_default"`
	BackFemale       any      `json:"back_female"`
	BackShinyURL     string   `json:"back_shiny"`
	BackShinyFemale  any      `json:"back_shiny_female"`
	FrontDefaultURL  string   `json:"front_default"`
	FrontFemale      any      `json:"front_female"`
	FrontShinyURL    string   `json:"front_shiny"`
	FrontShinyFemale any      `json:"front_shiny_female"`
	Other            Other    `json:"other"`
	Versions         Versions `json:"versions"`
}

type Sprite struct {
	Animated			Animated `json:"animated"`
	BackDefaultURL      string   `json:"back_default"`
	BackGrayURL		    string   `json:"back_gray"`
	BackTransparentURL  string   `json:"back_transparent"`
	BackFemale          any		 `json:"back_female"`
	BackShinyURL        string   `json:"back_shiny"`
	BackShinyFemale     any		 `json:"back_shiny_female"`
	FrontDefaultURL     string	 `json:"front_default"`
	FrontGrayURL		string	 `json:"front_gray"`
	FrontTransparentURL string	 `json:"front_transparent"`
	FrontFemale         any		 `json:"front_female"`
	FrontShinyURL       string	 `json:"front_shiny"`
	FrontShinyFemale    any		 `json:"front_shiny_female"`
}

type Animated struct {
	BackDefaultURL      string   `json:"back_default"`
	BackGrayURL		    string   `json:"back_gray"`
	BackTransparentURL  string   `json:"back_transparent"`
	BackFemale          any		 `json:"back_female"`
	BackShinyURL        string   `json:"back_shiny"`
	BackShinyFemale     any		 `json:"back_shiny_female"`
	FrontDefaultURL     string	 `json:"front_default"`
	FrontGrayURL		string	 `json:"front_gray"`
	FrontTransparentURL string	 `json:"front_transparent"`
	FrontFemale         any		 `json:"front_female"`
	FrontShinyURL       string	 `json:"front_shiny"`
	FrontShinyFemale    any		 `json:"front_shiny_female"`
}

type Other struct {
	DreamWorld      Sprite `json:"dream_world"`
	Home            Sprite `json:"home"`
	OfficialArtwork Sprite `json:"official-artwork"`
	Showdown        Sprite `json:"showdown"`
}

type Versions struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationII   GenerationII   `json:"generation-ii"`
	GenerationIII  GenerationIII  `json:"generation-iii"`
	GenerationIV   GenerationIV   `json:"generation-iv"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVI   GenerationVI   `json:"generation-vi"`
	GenerationVII  GenerationVII  `json:"generation-vii"`
	GenerationVIII GenerationVIII `json:"generation-viii"`
}

type GenerationI struct {
	RedBlue Sprite `json:"red-blue"`
	Yellow  Sprite `json:"yellow"`
}

type GenerationII struct {
	Crystal Sprite `json:"crystal"`
	Gold    Sprite `json:"gold"`
	Silver  Sprite `json:"silver"`
}

type GenerationIII struct {
	Emerald          Sprite `json:"emerald"`
	FireredLeafgreen Sprite `json:"firered-leafgreen"`
	RubySapphire     Sprite `json:"ruby-sapphire"`
}

type GenerationIV struct {
	DiamondPearl        Sprite `json:"diamond-pearl"`
	HeartgoldSoulsilver Sprite `json:"heartgold-soulsilver"`
	Platinum            Sprite `json:"platinum"`
}

type GenerationV struct {
	BlackWhite Sprite `json:"black-white"`
}

type GenerationVI struct {
	OmegarubyAlphasapphire Sprite `json:"omegaruby-alphasapphire"`
	XY                     Sprite `json:"x-y"`
}

type GenerationVII struct {
	Icons             Sprite `json:"icons"`
	UltraSunUltraMoon Sprite `json:"ultra-sun-ultra-moon"`
}

type GenerationVIII struct {
	Icons Sprite `json:"icons"`
}

type Stat struct {
	BaseStat int  		 `json:"base_stat"`
	Effort   int  		 `json:"effort"`
	Details  StatDetails `json:"stat"`
}

type StatDetails struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}