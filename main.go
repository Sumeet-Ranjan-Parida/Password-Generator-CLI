package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var length int

	fmt.Printf("Enter the Password length: ")
	fmt.Scanf("%d", &length)
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*-"

	rand.Seed(time.Now().UnixNano())
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = characters[rand.Intn(len(characters))]
	}
	str := string(buf)
	fmt.Printf("Your Password: %s", str)
}
