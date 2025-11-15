package security

import "golang.org/x/crypto/bcrypt"

// Receives a string and generates a hash form this password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compare password with a hash and returns if they are equals
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(password))
}

