package service

import "golang.org/x/crypto/bcrypt"

type IPasswordHasherService interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
}

type bcryptPasswordHasher struct{}

func (ph *bcryptPasswordHasher) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (ph *bcryptPasswordHasher) Compare(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func NewPasswordHasher() IPasswordHasherService {
	return &bcryptPasswordHasher{}
}
