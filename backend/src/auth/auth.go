package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

/*
Generates a bcryp hash from the provided password
*/
func HashPassword(password string) []byte {
	// Generate bcrypt hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}
	// Return the hash
	return hashedPassword
}

/*
Validates the provided password against the stored hashed password.
Returns a JWT token
*/
func CheckPassword(storedHashedPassword []byte, password, name, surname, email string) (string, error) {

	// Validate the provided password using bcrypt
	err := bcrypt.CompareHashAndPassword(storedHashedPassword, []byte(password))
	handleErrors(err)

	// If the password is valid, create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set JWT claims with user information
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix() //Set expiration time to 1 hour
	claims["name"] = name
	claims["surname"] = surname
	claims["email"] = email

	//Get the secret key

	secretKey, err := GetSecretKey()
	handleErrors(err)

	// Add the secret key to sign the token
	tokenString, err := token.SignedString([]byte(secretKey))
	handleErrors(err)

	// Return the generated JWT token string
	return tokenString, nil
}

func handleErrors(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
