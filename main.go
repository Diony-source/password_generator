package passwordgenerator

import (
	"fmt"
	"password-generator/handlers"
)

func main() {
	fmt.Println("Welcom to the Random Password Generator!")
	handlers.Start()
}