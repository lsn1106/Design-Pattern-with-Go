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

	//Otp구조체의 필드를 피우고 다음과같이 사용할수도 있다.
	//smsOTP := &otp.Sms{}
	//smsOTP.GenAndSendOTP(4)
	//
	//fmt.Println("")
	//emailOTP := &otp.Email{}
	//emailOTP.GenAndSendOTP(4)
}
