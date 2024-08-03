package otprepository

import "time"
 
type OTPSchema struct {
	    OtpID       int       `gorm:"column:otpId;primaryKey;autoIncrement"`
    UserID      int       `gorm:"column:userId;not null"`
    PhoneNumber string    `gorm:"column:phoneNumber;type:varchar(15);not null"`
    OtpNumber   int       `gorm:"column:otpNumber;not null"`
	OTPSessionId string   `gorm:"column:otp_session_id;char(36);not null"`
    IsVerified  bool      `gorm:"column:isVerified;not null;default:false"`
    CreatedAt   time.Time `gorm:"column:createdAt;not null;default:CURRENT_TIMESTAMP"`
}


type IOTPRepoSitory interface {
CreateOtp(UserId int, PhoneNumber string, OtpNumber int, OTPSessionId string)(error)
}

