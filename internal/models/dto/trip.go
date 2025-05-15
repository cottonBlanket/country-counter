package dto

type TripCreateDto struct {
	UserId    int64  `json:"user_id"`
	CountryId int64  `json:"country_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type TripUpdateDto struct {
	Id        int64  `json:"id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
