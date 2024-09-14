package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey string = "Thisisaverrrryyyyygoooooodddsecretkeyforgodsake"

func GenerateJWT(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func ValidateJWT(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		// checking signing type
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	// Check if token is signed with the private key
	if !(parsedToken.Valid) {
		return 0, errors.New("invalid token")
	}

	// checking if the claims is a MapClaims, as defined line 13 of this file
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token")
	}

	// email := claims["email"].(string)
	userID := claims["userId"].(int64)

	return userID, nil

}
