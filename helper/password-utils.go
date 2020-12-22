package helper

import (
	"github.com/raja/argon2pw"
)

func HashPassword(value string) (error, string) {
	hashed, err := argon2pw.GenerateSaltedHash(value)
	if err != nil {
		return err, hashed
	}

	return nil, hashed
}

func VerifyPassword(hash string, password string) bool {
	valid, _ := argon2pw.CompareHashWithPassword(hash, password)

	return valid
}
