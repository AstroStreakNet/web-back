package auth

import "golang.org/x/crypto/bcrypt"

type BcryptEncoder struct {
	hashCost int
}

func NewBcryptEncoder(hashCost int) *BcryptEncoder {
	return &BcryptEncoder{
		hashCost,
	}
}

func (encoder *BcryptEncoder) EncodePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), encoder.hashCost)
	return string(bytes), err
}

func (encoder *BcryptEncoder) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
