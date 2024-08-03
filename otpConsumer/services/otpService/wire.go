//go:build wireinject
// +build wireinject

package otpservice

import (
	otprepository "otpServiceConsumer/otpRepository"
	dbservice "otpServiceConsumer/services/dbService"

	"github.com/google/wire"
)

func InitializeOTPService(phrase string)*OTPServiceImpl{
	wire.Build(NewOTPServiceProvider,dbservice.NewDBServiceClientProvider,otprepository.NewOTPRepoProvider)
	return &OTPServiceImpl{}
}