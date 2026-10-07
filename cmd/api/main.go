package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	fmt.Println("env vars read and ready to go!")
}
