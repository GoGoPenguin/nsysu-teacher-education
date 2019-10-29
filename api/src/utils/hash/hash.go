package hash

import (
	"math/rand"

	"github.com/nsysu/teacher-education/src/utils/typecast"
	"golang.org/x/crypto/bcrypt"
)

//New takes a string and returns a hashed one
func New(data string) string {
	salt, _ := bcrypt.GenerateFromPassword([]byte(typecast.ToString(rand.Int())), bcrypt.MinCost)
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.MinCost)

	if err != nil {
		panic(err)
	}
	return string(salt) + string(hashed)
}

// Verify decides whether hashed is generated from raw
func Verify(raw string, hashed string) bool {
	password := hashed[60:]
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(raw))

	if err != nil {
		return false
	}
	return true
}
