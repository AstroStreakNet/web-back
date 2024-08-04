package auth

type Manager interface {
	GenerateJWT(role string, minutes int) (string, error)
	CheckJWT(jwt string) (bool, error)
	GetClaim(jwt, claim string) (bool, error)
}
