package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func GenerateSecretToken() (string, error) {
	config := viper.New()

	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	if err != nil {
		return "", err
	}

	secretToken := config.GetString("SECRET_TOKEN")

	return secretToken, nil
}

func GenerateJWT(claims *jwt.MapClaims) (string, error) {
	secretToken, err := GenerateSecretToken()

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(secretToken))

	if err != nil {
		return "", err
	}

	return webToken, nil
}

func VerifyTokenJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, isValid := token.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}

		secretToken, err := GenerateSecretToken()
		if err != nil {
			return nil, err
		}

		return []byte(secretToken), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeTokenJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyTokenJWT(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}
	
	return nil, fmt.Errorf("INVALID TOKEN")
}