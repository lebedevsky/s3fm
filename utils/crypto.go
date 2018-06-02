package utils

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

func GetRandomString(len int) (string, error){
	randBytes := make([]byte, len)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", err
	}

	str := base64.StdEncoding.EncodeToString(randBytes)[:len]
	// Base 64 can be longer than len
	return str, nil
}

func GetSalt() (string, error) {
	return GetRandomString(32)
}

func GetRandomPassword() (string, error) {
	return GetRandomString(16)
}

func GetPasswordHash(pass, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(salt + pass + salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(originHash, originSalt, newPass string) (bool, error) {
	newHash, err := GetPasswordHash(newPass, originSalt)
	if err != nil {
		return false, err
	}
	if originHash == newHash {
		return true, nil
	}
	return false, nil
}