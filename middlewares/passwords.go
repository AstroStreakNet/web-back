package middlewares

import "golang.org/x/crypto/bcrypt"

// Const
const hashCost int = 8 // Amount of iterations of password hashing, min 2 and max 31

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
