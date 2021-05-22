package validators

import (
	"api-grpc/pb"
	"errors"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

var (
	ErrInvalidUserId     = errors.New("Invalid userId")
	ErrEmptyName         = errors.New("name can't be empty")
	ErrEmptyEmail        = errors.New("email can't be empty")
	ErrEmptyPassword     = errors.New("password can't be empty")
	ErrEmailAlredyExists = errors.New("email already exists")
)

func ValidateSignUp(user *pb.User) error {

	if !bson.IsObjectIdHex(user.Id) {
		return ErrInvalidUserId
	}

	if user.Name == "" {
		return ErrEmptyName
	}

	if user.Email == "" {
		return ErrEmptyEmail
	}

	if user.Password == "" {
		return ErrEmptyPassword
	}

	return nil
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
