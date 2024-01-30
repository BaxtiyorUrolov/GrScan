package check

import (
	"errors"
	"fmt"
	"grscan/storage"
	"math/rand"
	"time"
	"unicode"

	"github.com/sfreiberg/gotwilio"
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

func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password length should be more than 6")
	}
	return nil
}

func IsLoginExist(login string, userStorage storage.IUserStorage) (bool, error) {
	exists, err := userStorage.IsLoginExist(login)
	if err != nil {
		return false, fmt.Errorf("error while checking login existence: %w", err)
	}
	return exists, nil
}

func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	return fmt.Sprintf("%06d", rand.Intn(max-min+1)+min)
}

func Send(toNumber, fromNumber, code string) error {
	accountSid := "YOUR_TWILIO_ACCOUNT_SID"
	authToken := "YOUR_TWILIO_AUTH_TOKEN"

	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	message := fmt.Sprintf("Your verification code is: %s", code)

	_, _, err := twilio.SendSMS(fromNumber, toNumber, message, "", "")
	if err != nil {
		return err
	}

	return nil
}
