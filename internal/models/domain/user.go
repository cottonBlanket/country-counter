package domain

type User struct {
	Id   int64  `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
