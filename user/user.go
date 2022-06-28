package user

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	name     string
	email    Email
	password Password
}

func SetUser(name, email, password string) (User, bool) {
	if Email(email).isValid() && Password(password).isValid() {
		hashed_password := Password(password).hashedPassword()
		return User{
			name:     name,
			email:    Email(email),
			password: Password(hashed_password),
		}, true
	}
	return User{}, false

}

type Password string

func (p Password) isValid() bool {
	return len(p) > 7
}

func (p Password) hashedPassword() string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(p), 1)
	return string(hashed)
}

type Email string

func (e Email) isValid() bool {
	_, err := mail.ParseAddress(string(e))
	return err == nil
}
