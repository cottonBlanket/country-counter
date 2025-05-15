package dto

type CountryCreateDto struct {
	Name string `json:"name"`
}

type CountryUpdateDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
