package main

import (
	"fmt"
	"template_method/otp"
)

func main() {
	smsOTP := &otp.Sms{}
	o := otp.Otp{
		IOtp: smsOTP,
	}
	o.GenAndSendOTP(4)

	fmt.Println("")
	emailOTP := &otp.Email{}
	o = otp.Otp{
		IOtp: emailOTP,
	}
	o.GenAndSendOTP(4)
}
