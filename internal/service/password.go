package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher defines an interface for hashing passwords
type PasswordHasher interface {
	GenerateHash(password string) (string, error)
	CompareHashAndPassword(hash, password []byte) error
}

// BcryptPasswordHasher is an implementation of PasswordHasher that uses bcrypt
type BcryptPasswordHasher struct{}

// GenerateHash generates a hash for the given password using bcrypt
func (b *BcryptPasswordHasher) GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("can't generate password hash")
	}
	return string(hashedPassword), nil
}

func (b *BcryptPasswordHasher) CompareHashAndPassword(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
