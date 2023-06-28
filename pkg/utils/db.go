package utils

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GetDBNameFromDriverSource(str string) string {
	if str == "" {
		return ""
	}
	if strings.Contains(str, "dbname=") {
		array1 := strings.Split(str, "dbname=")
		if len(array1) == 0 {
			return ""
		}
		array2 := strings.Split(array1[1], " ")
		if len(array2) == 0 {
			return ""
		}
		return array2[0]
	}
	if strings.Contains(str, "database=") {
		array1 := strings.Split(str, "database=")
		if len(array1) == 0 {
			return ""
		}
		array2 := strings.Split(array1[1], ";")
		if len(array2) == 0 {
			return ""
		}
		return array2[0]
	}
	return ""
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
