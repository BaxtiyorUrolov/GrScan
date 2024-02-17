package check

import (
	"fmt"
	"grscan/storage"
	"unicode"
)

func PhoneNumber(phone string) bool {
	for _, r := range phone {
		if r == '+' {
			continue
		} else if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func ValidatePassword(password string) bool {
	if len(password) < 6 {
		return false
	}

	var (
		hasUpperCase bool
		hasLowerCase bool
	)

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpperCase = true
		}
		if unicode.IsLower(char) {
			hasLowerCase = true
		}
	}

	return hasUpperCase && hasLowerCase
}

func IsLoginExist(login string, userStorage storage.IUserStorage) (bool, error) {
	exists, err := userStorage.IsLoginExist(login)
	if err != nil {
		return false, fmt.Errorf("error while checking login existence: %w", err)
	}
	return exists, nil
}
