package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashMyPassword( password string) (string, error)  {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}


func CheckHashPassword(password string , hashedpassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
    return err == nil
}
