package otp

import "fmt"

type Email struct {
	Otp
}

func (s *Email) SendNotification(message string) error {
	fmt.Printf("Email: sending sms: %s\n", message)
	return nil
}
