package otpservice

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	otprepository "otpServiceConsumer/otpRepository"
	"strconv"
)

type TwoFactorSendOTPResponse struct {
	Status string
	Details string
	OTP string
}
type IOTPService interface {
	/**
	The primary function of it is to send otp and create one corresponsing otp entry in the db.
	*/
SendOtpWith2Factor(UserId int,PhoneNumber string) error
}

type OTPServiceImpl struct {
repo otprepository.IOTPRepoSitory	
}

func (o *OTPServiceImpl) SendOtpWith2Factor(url string,UserId int,PhoneNumber string) error {
	var completeURL string=fmt.Sprintf("%s%s/AUTOGEN2/testTemplate",url,PhoneNumber)
	println(UserId)
	 	resp,err:=http.Get(completeURL)
		 if err != nil {
        log.Fatalf("Failed to send GET request: %v", err)
    }
    defer resp.Body.Close()

	 body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read response body: %v", err)
    }

    // Print the response status and body
    fmt.Printf("Response Status: %s\n", resp.Status)
    fmt.Printf("Response Body: %s\n", string(body))

	var apiResponse TwoFactorSendOTPResponse;
	// pointer to apiResponse because it is gonna mutate the original predeclared apiResponse of type TwoFactorSendOTPResponse
	err=json.Unmarshal(body,&apiResponse)
	
	if err != nil {
        log.Fatalf("Failed to parse JSON response: %v", err)
    }
otp,err:=strconv.Atoi(apiResponse.OTP)
if err!=nil{
    log.Fatalf("Error parsing string to int: %v", err)
}
	o.repo.CreateOtp(UserId,PhoneNumber,otp,apiResponse.Details)

    // Print the parsed response
    fmt.Printf("Parsed Response: %+v\n", apiResponse)
	
return nil	
}

func NewOTPServiceProvider(repo *otprepository.OTPRepoImpl) *OTPServiceImpl{
	return &OTPServiceImpl{repo: repo}
}