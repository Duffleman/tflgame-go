package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pwd := getPwd()

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(hash))
}

func getPwd() []byte {
	fmt.Println("Enter a password")

	var pwd string

	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}

	return []byte(pwd)
}
