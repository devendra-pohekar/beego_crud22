package controllers

import (
	"crudDemo/helpers"
	requestStruct "crudDemo/requstStruct"
	"encoding/json"
	"log"
	"strings"

	"github.com/astaxie/beego"
	"github.com/sfreiberg/gotwilio"
	// "github.com/sfreiberg/gotwilio"
)

type ThirdPartyOtpSend struct {
	beego.Controller
}

func (c *ThirdPartyOtpSend) SendOTP() {
	var mobileNumber requestStruct.MobileNumber
	json.Unmarshal(c.Ctx.Input.RequestBody, &mobileNumber)
	log.Print(mobileNumber)

	if !strings.HasPrefix(mobileNumber.MobileNumber, "+") {
		mobileNumber.MobileNumber = "+91" + mobileNumber.MobileNumber // Assuming +1 for the United States
	}

	accountSid := "AC7df5ab60abcbbde9fcbda3c32b5c40d5"
	authToken := "3c4648b746d6c8377a4df2ff03c5f9e6"

	// Your Twilio phone number
	from := "+91 6264736064"

	// Recipient's phone number (to whom you want to send OTP)
	to := mobileNumber.MobileNumber
	log.Print(to)

	// Generate a random 6-digit OTP
	otp := helpers.GenerateOTP()
	log.Print(otp)

	// Your Twilio client
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)
	// Compose the message
	message := "Your OTP is: " + otp

	// Send the message
	_, exception, err := twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		// Handle the error
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else if exception != nil {
		// Handle the exception
		log.Println("Twilio Exception:", exception.Message)
		c.Data["json"] = map[string]interface{}{"exception": exception.Message}
	} else {
		// Message sent successfully
		c.Data["json"] = map[string]interface{}{"success": true}
	}

	c.ServeJSON()
}

func (c *ThirdPartyOtpSend) VerifyOTP() {
	// Retrieve the OTP entered by the user
	userOTP := c.GetString("otp")
	// Retrieve the stored OTP from session or database
	storedOTP, ok := c.GetSession("otp").(string)
	if !ok {
		// Handle the case where OTP is not found
		c.Data["json"] = map[string]interface{}{"error": "OTP not found"}
		c.ServeJSON()
		return
	}
	// Compare the entered OTP with the stored OTP
	if userOTP == storedOTP {
		// OTP is correct, you can proceed with further actions
		c.Data["json"] = map[string]interface{}{"success": true, "message": "OTP verification successful"}
	} else {
		// OTP is incorrect
		c.Data["json"] = map[string]interface{}{"error": "Incorrect OTP"}
	}

	c.ServeJSON()
}
