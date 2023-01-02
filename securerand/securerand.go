package securerand

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func Bytes(length int) ([]byte, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func Hex(length int) (string, error) {
	b, err := Bytes(length)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func Base64(length int) (string, error) {
	b, err := Bytes(length)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
