package pokeapi

type LocationArea struct {
	Count	 int 				   `json:"count"`
	Next	 *string			   `json:"next"`
	Previous *string			   `json:"previous"`
	Results	 []LocationAreaDetails `json:"results"`
}

type LocationAreaDetails struct {
	Name string `json:"name"`
	URL	 string `json:"url"`
}