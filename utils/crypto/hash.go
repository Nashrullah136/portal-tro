package crypto

import "golang.org/x/crypto/bcrypt"

type Hash interface {
	Hash(pwd string) (string, error)
	Compare(plainPassword, hashedPassword string) error
}

type bcryptHash struct {
}

func NewBcryptHash() Hash {
	return bcryptHash{}
}

func (b bcryptHash) Hash(pwd string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword[:]), nil
}

func (b bcryptHash) Compare(plainPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
