package middlewares

import (
    "golang.org/x/crypto/bcrypt"
)

// Const
const hashCost int = 8 // Amount of iterations of password hashing, min 2 and max 31

// Hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return string(bytes), err
}

// Compares a password with its hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
