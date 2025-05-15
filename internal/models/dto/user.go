package dto

type UserCreateDto struct {
	Name string `json:"name"`
}

type UserUpdateDto struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
