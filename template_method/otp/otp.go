package otp

import "fmt"

type IOtp interface {
	GenRandomOTP(int) string
	SaveOTPCache(string)
	GetMessage(string) string
	SendNotification(string) error
}

type Otp struct {
	IOtp IOtp
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
	otp := o.IOtp.GenRandomOTP(otpLength)
	o.IOtp.SaveOTPCache(otp)
	message := o.IOtp.GetMessage(otp)
	err := o.IOtp.SendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

func (o *Otp) GenRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("Common: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (o *Otp) SaveOTPCache(otp string) {
	fmt.Printf("Common: saving otp: %s to cache\n", otp)
}

func (o *Otp) GetMessage(otp string) string {
	return "Common OTP for login is " + otp
}

func (o *Otp) SendNotification(message string) error {
	fmt.Printf("Common: sending sms: %s\n", message)
	return nil
}
