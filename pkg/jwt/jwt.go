package jwt_helper

import (
	"BasketProjectGolang/pkg/config"
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type DecodedToken struct {
	Iat    int      `json:"iat"`
	Roles  []string `json:"roles"`
	UserId string   `json:"userId"`
	Email  string   `json:"email"`
	Iss    string   `json:"iss"`
}

func GenerateToken(claims *jwt.Token, cfg *config.Config) string {
	hmacSecretString := cfg.JWTConfig.SecretKey
	hmacSecret := []byte(hmacSecretString)
	token, _ := claims.SignedString(hmacSecret)

	return fmt.Sprintf("%s%s", cfg.JWTConfig.TokenPrefix, token)
}

func VerifyToken(token string, secret string) *DecodedToken {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil || !decoded.Valid {
		return nil
	}

	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)
	json.Unmarshal(jsonString, &decodedToken)

	return &decodedToken
}
