package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func GenerarToken(correo string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = correo
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwtToken.SignedString([]byte(os.Getenv("API_SECRET")))

}

func ValidarToken(r *http.Request) error {
	jwtToken := ExtraerToken(r)
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(
				"Método indesperado: %s",
				token.Header["alg"])
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
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}

func ExtraerToken(r *http.Request) string {
	parametros := r.URL.Query()
	token := parametros.Get("token")
	if token != "" {
		return token
	}
	tokenString := r.Header.Get("Authorization")
	if len(strings.Split(tokenString, "")) == 2 {

		return strings.Split(tokenString, "")[1]
	}
	return ""

}
