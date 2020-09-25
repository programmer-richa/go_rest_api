package models

import "strings"

// UserSpecification defines filters applicable on the User struct.
type UserSpecification interface {
	IsSatisfied(u *User) bool
}

type EmailSpecification struct {
	Email string
}

func (spec EmailSpecification) IsSatisfied(u *User) bool {
	return strings.ToLower(u.Email) == strings.ToLower(spec.Email)
}

type PasswordSpecification struct {
	Password string
}

func (spec PasswordSpecification) IsSatisfied(u *User) bool {
	return u.Password == spec.Password
}

type UserCredentialsSpecification struct {
	EmailSpecification
	PasswordSpecification
}

func (spec UserCredentialsSpecification) IsSatisfied(u *User) bool {

	return spec.EmailSpecification.IsSatisfied(u) &&
		spec.PasswordSpecification.IsSatisfied(u)

}

type Filter struct{}

func (f *Filter) FilterUser(users []*User, spec UserSpecification) []*User {
	result := make([]*User, 0)
	for i, v := range users {
		if spec.IsSatisfied(v) {
			result = append(result, users[i])
		}
	}
	return result
}
