package dto

type UserCreateDto struct {
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	MiddleName string `json:"middleName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

type UserFilter struct {
	Search string
}
