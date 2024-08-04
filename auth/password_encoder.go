package auth

type PasswordEncoder interface {
	EncodePassword(password string) (string, error)
	CheckPassword(password, hash string) bool
}
