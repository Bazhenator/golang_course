package models

import (
	"regexp"
	"strings"
)

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

func IsValidPhoneNumber(phone string) bool {
	pattern := `^\+\d{11}$`
	match, _ := regexp.MatchString(pattern, phone)
	return match
}

func IsValidPassword(password string) bool {
	return len(password) >= 8 && strings.ContainsAny(password, "0123456789") && strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") &&
		strings.ContainsAny(password, "~!@#$^&*()_+")
}

func IsEmailUnique(email string) bool {
	var count int
	GetDB().Table("accounts").Where("email = ?", email).Count(&count)
	return count == 0
}

func IsValidName(name string) bool {
	return ((strings.ContainsAny(name, "0123456789") || strings.ContainsAny(name, "!@#$%^&*()~")) &&
		strings.ContainsAny(name, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")) && name != "" || strings.ContainsAny(name, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
