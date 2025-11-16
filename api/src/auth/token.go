package auth

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Define token with claims
func GenerateToken(userID uint64) (string, error) {
	/*Defining a Token*/
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userId"] = userID

	/*Generate Token*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	/*Sign token
	Input a secret as parameter, should be securely generated, also as a good practice remove from source code and put it into an .env file
	*/
	return token.SignedString([]byte(config.SecretKey))
}