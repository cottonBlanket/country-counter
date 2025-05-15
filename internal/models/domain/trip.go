package domain

type Trip struct {
	Id int64 `json:"id"`

	UserId int64 `json:"user_id"`
	User   User  `json:"user"`

	CountryId int64   `json:"country_id"`
	Country   Country `json:"country"`

	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
