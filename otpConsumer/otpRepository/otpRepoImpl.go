package otprepository

import (
	"errors"
	"fmt"
	"log"
	dbservice "otpServiceConsumer/services/dbService"
	"time"

	"gorm.io/gorm"
)

type OTPRepoImpl struct{
	db *dbservice.MYSQLDBService
}

func (o *OTPRepoImpl) CreateOtp(UserId int, PhoneNumber string, OtpNumber int, OTPSessionId string) error {
    // Check if an OTP with the same session ID already exists
    var existingOtp OTPSchema
	db,err:=o.db.GetDb();
    if err!=nil {
		log.Fatalf("Error occured while getting db instance:%v",err)
	}
    result := db.Table("OTP").Where("otp_session_id = ?", OTPSessionId).First(&existingOtp)
    if result.Error == nil {
        // OTP with the same session ID already exists
        log.Printf("OTP with session ID %s already exists", OTPSessionId)
        return fmt.Errorf("OTP with session ID %s already exists", OTPSessionId)
    } else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
        // An error other than 'record not found' occurred
        log.Fatalf("Failed to check existing OTP: %v", result.Error)
        return result.Error
    }

    // No existing OTP with the same session ID, create a new one
    otp := OTPSchema{
        UserID:       UserId,
        PhoneNumber:  PhoneNumber,
        OtpNumber:    OtpNumber,
        OTPSessionId: OTPSessionId,
        IsVerified:   false,
        CreatedAt:    time.Now(),
    }
    result = db.Table("OTP").Create(&otp)

    if result.Error != nil {
        log.Fatalf("Failed to create OTP row: %v", result.Error)
        return result.Error
    } else {
        log.Printf("OTP row created successfully with ID: %d", otp.OtpID)
    }
    return nil
}


func NewOTPRepoProvider(db *dbservice.MYSQLDBService)*OTPRepoImpl{
	return &OTPRepoImpl{db: db}
}