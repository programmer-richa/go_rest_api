package repository

import (
	"github.com/programmer-richa/go_rest_api/models"
	"github.com/programmer-richa/go_rest_api/settings"
)

// UserRepository contains functional operations on User struct
type UserRepository struct{}

// Signup adds a new user if the email is not registered with an existing account.
func (u UserRepository) Signup(user *models.User) bool {
	emailSpec := models.EmailSpecification{Email: user.Email}
	f := models.Filter{}
	found := false
	for _, v := range f.FilterUser(settings.Users, emailSpec) {
		_ = v
		found = true
		break
	}
	if !found {
		settings.Users = append(settings.Users, user)
	}

	return !found
}

// Login generates token for the user if valid credentials are provided.
func (u UserRepository) Login(user *models.User) *models.User {
	emailSpec := models.EmailSpecification{Email: user.Email}
	passwordSpec := models.PasswordSpecification{Password: user.Password}
	userSpec := models.UserCredentialsSpecification{
		EmailSpecification:    emailSpec,
		PasswordSpecification: passwordSpec,
	}
	f := models.Filter{}
	var validUser *models.User
	for _, v := range f.FilterUser(settings.Users, userSpec) {
		validUser = v
		break
	}
	return validUser
}
