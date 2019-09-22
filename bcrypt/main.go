package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("bcrypted:", string(bs))

	if err := bcrypt.CompareHashAndPassword(bs, []byte(password)); err != nil {
		fmt.Println("invalid password")
		os.Exit(1)
	}
	fmt.Println("matched", string(bs), "and", password)

	cost, err := bcrypt.Cost(bs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("cost", cost)
}
