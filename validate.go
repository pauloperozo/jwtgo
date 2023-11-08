package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InBhdWxvcGVyb3pvQGdtYWlsLmNvbSIsImV4cCI6MTY5OTUwOTgyOSwiaWF0IjoxNjk5NDIzNDI5LCJ1c2VyIjoiUGF1bG8ifQ.VnjM84LoVLxHfhMos8Brh5L4brahZrhXL1FTuosFR0k"
	secret      = []byte("this is secret")
)

func validateMethodAndGetSecret(token *jwt.Token) (any, error) {

	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("Metodo no es de la familia")
	}
	return secret, nil
}

func Validate() {
	token, err := jwt.Parse(tokenString, validateMethodAndGetSecret)
	if err != nil {
		fmt.Println("Token no valido")
	}

	userData, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error Payload")
		return
	}

	fmt.Println("User", userData["user"])
	fmt.Println("User", userData["email"])

}
