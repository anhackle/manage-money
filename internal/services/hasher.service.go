package service

import "golang.org/x/crypto/bcrypt"

type IPasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) bool
}

type bcryptPasswordHasher struct{}

func (ph *bcryptPasswordHasher) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (ph *bcryptPasswordHasher) Compare(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewPasswordHahser() IPasswordHasher {
	return &bcryptPasswordHasher{}
}
