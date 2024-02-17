package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"grscan/api/models"
	"grscan/config"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	return fmt.Sprintf("%06d", rand.Intn(max-min+1)+min)
}

func Send(toNumber, code string) error {
	fromNumber := "4546"
	// Eskiz API endpoint for sending SMS
	apiURL := "https://notify.eskiz.uz/api/message/sms/send"

	cfg := config.Load()

	// Create SMS data
	smsData := models.SMS{
		MobilePhone: toNumber,        
		Message:     fmt.Sprintf("Sizning tasdiqlash kodingiz: %s", code), 
		From:        fromNumber,           
	}

	// Convert SMS data to JSON
	jsonData, err := json.Marshal(smsData)
	if err != nil {
		return err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// Add authorization header
	req.Header.Set("Authorization", "Bearer "+cfg.Token) // Token o'zgaruvchisi config obyektidan olinadi
	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse JSON response
	var smsResponse models.SMSResponse
	err = json.Unmarshal(body, &smsResponse)
	if err != nil {
		return err
	}

	// Check response status
	fmt.Println("Message ID:", smsResponse.MessageID)
	fmt.Println("Status:", smsResponse.Status)

	return nil
}
