package services

import (
	"math/rand"
	"time"
)

func GenratePassword(length int, charTypes string) string {
	var charset string
	if charTypes == "letters" || charTypes == "letters,numbers,symbols" {
		charset += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if charTypes == "numbers" || charTypes == "letters,numbers,symbols" {
		charset += "0123456789"
	}
	if charTypes == "symbols" || charTypes == "letters,numbers,symbols" {
		charset += "!@#$%^&*()_+-=[]{}|;:',.<>?/"
	}

	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}