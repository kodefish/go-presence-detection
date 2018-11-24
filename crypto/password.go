package crypto

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt returns a hashed and salted version of the password
func HashAndSalt(password string) string {
	return string(hashAndSalt([]byte(password)))
}

func hashAndSalt(pwd []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return hash
}

// ComparePasswords returns true if the password matches the hash
func ComparePasswords(hashedPwd, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}
