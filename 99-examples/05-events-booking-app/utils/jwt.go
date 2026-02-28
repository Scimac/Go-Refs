package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// should be secure key.
const secretKey = "GangadharIsShaktimaan"

func GenerateJwtToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"uid":   userId,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	// byte slice needed even though SignedSteing says any {}interface
	return token.SignedString([]byte(secretKey))
}

func VerifyJwtToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		// HMAC is the parent of HS256 method
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Siging Method not correct.")
		}

		// since we converted the key into byte slice in GenerateJwtToken:21
		return []byte(secretKey), nil
	})
	if err != nil {
		// userId returned as 0.
		return 0, errors.New("Could not parse the token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid Token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Could not parse claims")
	}

	// JWT Claims store it as float64, need to convert
	exp, ok := claims["exp"].(float64)
	if !time.Now().Add(-time.Hour).Before(time.Unix(int64(exp), 0)) {
		// token should not be beofre one hr from now
		return 0, errors.New("Token is expired")
	}

	// email = claims["email"].(string)
	// JWT Claims store it as float64, need to convert
	userId := int64(claims["uid"].(float64))

	return userId, nil
}
