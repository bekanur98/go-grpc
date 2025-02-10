package main

import (
	"errors"
	"fmt"
)

func main() {
	message, err := isGolangSuitableForMe([]string{"c#", "python", "java", "js"})
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
}

func isGolangSuitableForMe(langs []string) (bool, error) {
	if len(langs) == 0 {
		return false, fmt.Errorf("No languages provided")
	}
	for _, lang := range langs {
		if lang == "php" {
			return false, errors.New("Golang is not suitable because 'php' is in the list")
		}
	}
	return true, nil
}
