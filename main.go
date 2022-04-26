package main

import (
	"fmt"
	"github.com/google/uuid"
	"go-jwt-rsa/tokens"
)

func main() {
	id := uuid.New()

	token, err := tokens.New(id.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(tokens.ToString(token))
}


