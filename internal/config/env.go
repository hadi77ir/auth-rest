package config

import "os"

func IsDevMode() bool {
	if envDev := os.Getenv("APP_ENV"); envDev == "development" {
		return true
	}
	return false
}

func IsProductionMode() bool {
	return !IsDevMode() && !IsStagingMode()
}

func IsStagingMode() bool {
	if envDev := os.Getenv("APP_ENV"); envDev == "staging" {
		return true
	}
	return false
}
