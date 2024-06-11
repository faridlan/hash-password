package main

import (
	"fmt"

	"github.com/faridlan/hash-password/service"
)

func main() {

	password := "secret1234"

	hashPassword, err := service.HashPassword(password)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(hashPassword))

}
