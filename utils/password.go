package utils

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckM5Password(password, md5Password string) bool {
	return Md5Password(password) == md5Password
}

func Md5Password(password string) string {
	data := []byte(password)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
