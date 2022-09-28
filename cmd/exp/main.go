package main

import (
	"errors"
	"fmt"
)

func main() {
	err := CreateUser()
	fmt.Println(err)

	err = CreateOrg()
	fmt.Println(err)
}
func Connect() error {
	return errors.New("CONNECTION FAILED")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("CREATE USER: %w", err)
	}
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("CREATE ORG: %w", err)
	}

	return nil
}
