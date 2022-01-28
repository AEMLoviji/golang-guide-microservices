package oauth

import "github.com/aemloviji/golang-guide-microservices/src/api/utils/errors"

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"elvin": {Id: 123, Username: "elvin"},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	if user, ok := users[username]; ok {
		return user, nil
	}

	return nil, errors.NewNotFoundError("no user found with given parameters")
}
