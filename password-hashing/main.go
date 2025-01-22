package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(bytes), err
}

func CheckPasswordHash(pw, hashedPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(pw))
	return err == nil
}

func main() {
	pw := "secret"
	hashPw, _ := HashPassword(pw) // ignore error for the sake of simplicity

	fmt.Println("Password:", pw)
	fmt.Println("Hash:    ", hashPw)

	match := CheckPasswordHash(hashPw, pw)
	fmt.Println("Match:   ", match)
}

// $ go run passwords.go
// Password: secret
// Hash:     $2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK
// Match:    true
