package utils

type UserData struct {
	Username string `json:"username"`; // `json:"username"` is a struct tag
	Email string `json:"email"`; // `json:"email"` is a struct tag
	Password string `json:"password"`; // `json:"password"` is a struct tag
}
