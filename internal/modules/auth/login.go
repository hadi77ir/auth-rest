package auth

import (
	"auth-rest/internal/dal"
	"auth-rest/internal/db/schema"
	"auth-rest/internal/modules/sms"
	"auth-rest/internal/modules/users"
	"crypto/rand"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
	RoleSuper = "super"
)

func InitiateUserLogin(usersRepo dal.UsersRepository, smsProvider sms.SMSProvider, phoneNumber string) (string, error) {
	user, err := usersRepo.ByPhoneNumber(phoneNumber)
	newUser := false
	if err != nil {
		if !errors.Is(err, dal.ErrRecordNotFound) {
			return "", err
		}
		user = &schema.User{Phone: phoneNumber, Role: RoleUser}
		newUser = true
	}
	if time.Now().After(user.OTPCodeExpires) {
		user.OTPCode = generateNewOtp()
		user.OTPCodeExpires = time.Now().Add(2 * time.Minute)
		if newUser {
			err = usersRepo.Create(user)
		} else {
			err = usersRepo.Update(user)
		}
		if err != nil {
			return "", err
		}
	}
	err = smsProvider.SendSMS(phoneNumber, fmt.Sprintf("Code: %s", user.OTPCode))
	if err != nil {
		return "", err
	}
	return user.OTPCode, nil
}

var ErrInvalidOTP = errors.New("invalid OTP code")

func VerifyOTP(usersRepo dal.UsersRepository, phone, otpCode string) (users.UserModel, error) {
	var empty users.UserModel
	if otpCode == "" {
		return empty, ErrInvalidOTP
	}
	user, err := usersRepo.ByPhoneNumber(phone)
	if err != nil {
		return empty, err
	}
	if user.OTPCode != otpCode || user.OTPCodeExpires.Before(time.Now()) {
		return empty, ErrInvalidOTP
	}
	user.OTPCodeExpires = time.Now().Add(-1 * time.Minute)
	user.OTPCode = ""
	user.Verified = true
	err = usersRepo.Update(user)
	if err != nil {
		return empty, err
	}
	return users.ToModel(user), nil
}

func generateNewOtp() string {
	buf := &strings.Builder{}
	rands := [6]byte{}
	_, _ = rand.Read(rands[:])
	for i := 0; i < 6; i++ {
		buf.WriteString(strconv.Itoa(int(rands[i] % 10)))
	}
	return buf.String()
}
