package handlers

import (
	"bufio"
	"fmt"
	"os"
	"password-generator/entities"
	"password-generator/services"
)

func Start() {
	var length int
	var charTypes string

	fmt.Println("Enter the desired lenght of the password:")
	fmt.Scanln(&length)

	fmt.Println("Enter the character types (letters, numbers, symbols, or combinations like letters,numbers):")
	fmt.Scanln(&charTypes)

	request := entities.PasswordRequest{
		Length: length,
		CharacterTypes: charTypes,
	}

	password := services.GenratePassword(request.Length, request.CharacterTypes)

	fmt.Printf("Generated password: %s\n", password)

	saveToFile(password)
}

func saveToFile(password string) {
	file, err := os.Create("generated_password.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(password)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	writer.Flush()

	fmt.Println("Password saved to generated_password.txt")
}