package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main3() {

	user := "Paulo"
	email := "pauloperozo@gmail.com"
	now := time.Now()
	iat := now.Unix()
	eat := now.AddDate(0, 0, 1).Unix()
	secret := "this is secret"

	payload := jwt.MapClaims{
		"user":  user,
		"email": email,
		"iat":   iat,
		"exp":   eat,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tokenString)

	//Validate()
}
