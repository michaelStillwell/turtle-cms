package pkg

import "github.com/alexedwards/argon2id"

var hashParams = argon2id.DefaultParams

func HashPassword(p string) (string, error) {
	hash, err := argon2id.CreateHash(p, hashParams)
	return hash, err
}

func VerifyPassword(p, h string) bool {
	match, err := argon2id.ComparePasswordAndHash(p, h)
	if err != nil {
		return false
	}

	return match
}
