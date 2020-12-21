package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

var mysignedKey = []byte("thisisasuperscreetkeyforjwttoken")

func CreateJwtToken(username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Minute * 30)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(mysignedKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

//VerifyTOken will used to Verify the JWT Token while using APIS
func VerifyToken(tokenString string) (token *jwt.Token, err error) {
	claims := &Claims{}
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mysignedKey, nil
	})

	return token, err
}

func IsAuthorised(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// claims := &Claims{}
		// param := strings.Trim(r.Header.Get("Authorization"), "Bearer ")
		param := r.Header.Get("Token")
		token, err := VerifyToken(param)
		if err != nil {
			log.Panic(err.Error())
		}
		if token.Valid {
			endpoint(w, r)
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
