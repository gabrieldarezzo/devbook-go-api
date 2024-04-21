package security

import "golang.org/x/crypto/bcrypt"

// GenerateHash transform a string into a hash string
func GenerateHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPassword compare password and hashString and return if they are equals
func CheckPassword(passwordHashed string, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(passwordString))
}
