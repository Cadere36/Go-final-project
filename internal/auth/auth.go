package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:])
}

func signingKey(password string) []byte {
	return []byte(password)
}

func GenerateToken(password string) (string, error) {

	claims := jwt.MapClaims{"hash": hashPassword(password)}
	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signingKey(password))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(token, password string) bool {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(password), nil
	})
	if err != nil {
		return false
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return false
	}
	hash, ok := claims["hash"].(string)
	if !ok {
		return false
	}
	if hash != hashPassword(password) {
		return false
	}
	return true

}
