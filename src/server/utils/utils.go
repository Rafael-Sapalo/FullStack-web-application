package utils

import "net/http"

type UserData struct {
	Username string `json:"username"`; // `json:"username"` is a struct tag
	Email string `json:"email"`; // `json:"email"` is a struct tag
	Password string `json:"password"`; // `json:"password"` is a struct tag
}

type ErrorMessage struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

func (er *ErrorMessage) Error() string {
	return er.Message
}

var (
	ErrInternalServerError = &ErrorMessage{Message: "Internal Server Error", Code: http.StatusInternalServerError}
	ErrorEmailAlrdExists = &ErrorMessage{Message: "Email already exists or Username already exist", Code: http.StatusConflict}
	ErrorHashingPassword = &ErrorMessage{Message: "Error hashing password", Code: http.StatusBadRequest}
	ErrorInsertingUserData = &ErrorMessage{Message: "Error couldn't insert user data", Code: http.StatusInternalServerError}
	ErrorCommit = &ErrorMessage{Message: "Error couldn't commit the transaction", Code: http.StatusInternalServerError}
	ErrorUnauthorized = &ErrorMessage{Message: "Error Unauthorized", Code: http.StatusUnauthorized}
)

var (
	SuccessfullyRegistered = &ErrorMessage{Message: "Successfully registered", Code: http.StatusOK}
	SuccessfullyLoggedIn = &ErrorMessage{Message: "Successfully logged  in", Code: http.StatusOK}
	IsGoodPassword = &ErrorMessage{Message: "The password provided correspond", Code: http.StatusOK}
)