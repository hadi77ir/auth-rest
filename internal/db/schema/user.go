package schema

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Phone          string `gorm:"unique"`
	Verified       bool
	Name           string
	Role           string
	OTPCode        string
	OTPCodeExpires time.Time
}
