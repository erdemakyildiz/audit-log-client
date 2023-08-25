package handler

const Secret = "erdem"

type ErrorResponse struct {
	Error string `json:"error"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	ID             int
	Email          string
	Password       string
	FavoritePhrase string
}

type GetFiltersResponse struct {
	Data []string `json:"data"`
}
