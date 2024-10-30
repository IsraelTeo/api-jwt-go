package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = email
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return jwtToken.SignedString([]byte(os.Getenv("API_SECRET")))
}

func ValidateToken(r *http.Request) error {
	jwtToken := GetToken(r)
	token, err := jwt.Parse(jwtToken, func(f *jwt.Token) (interface{}, error) {
		if _, ok := f.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", f.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func Pretty(data interface{}) {
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(pretty))
}

func GetToken(r *http.Request) string {
	params := r.URL.Query()
	token := params.Get("token")
	if token != "" {
		return token
	}
	tokenString := r.Header.Get("Authorization")
	if len(strings.Split(tokenString, " ")) == 2 {
		return strings.Split(tokenString, " ")[1]
	}
	return ""
}
