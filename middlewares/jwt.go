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

func IsAuthorised(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       claims := &Claims{}
        if r.Header["Token"] != nil {
            token,err := jwt.ParseWithClaims(r.Header["Token"][0], claims,func(token *jwt.Token)(interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("there was an error with jwt token")
                }
                return mysignedKey, nil
            }) 
            if err != nil {
                log.Panic(err.Error())
            }
            if token.Valid {
                endpoint(w,r)
            }else {
                fmt.Fprintf(w, "Not Authorized")
            }

        }

   }) 
}

// func GenerateToken(username string) (token string, err error) {
//     token := jwt.New(jwt.SigningMethodES256)
    
//     claims := jwt.Claims.(jwt.MapClaims)

//     claims['authorized'] = true
//     claims['client'] = username
//     claims['exp'] = time.Now().Add(time.Minute * 30).Unix()

//     tokenString , err := token.SignedString(mysignedKey) 

//     if err != nil {
//         fmt.Errorf("something went wring with generating the token: %s",err.Error())
//     }

//     return tokenString, nil
// }


func CreateJwtToken(username string)(tokenString string, err error) {
    expirationTime :=  time.Now().Add(time.Minute * 30)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString , err = token.SignedString(mysignedKey)
    if err != nil {
        return "", err
    }
    return tokenString, err
}

//will used to Verify the JWT Token while using APIS
func VerifyToken(tokenString string )(username string, err error){
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error) {
        return mysignedKey, nil
    })

    if token != nil {
        return claims.Username, nil
    }
    return "", err
}
