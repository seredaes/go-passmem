package EncryptData

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var bytes = []byte("dvDG6CutiUj83DKN")

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encrypt(data string, encrypt_key string) (string, error) {
	block, err := aes.NewCipher([]byte(encrypt_key))
	if err != nil {
		return "", err
	}

	plainText := []byte(data)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return encode(cipherText), nil
}

func Decrypt(cryptedText string, encrypt_key string) (string, error) {
	block, err := aes.NewCipher([]byte(encrypt_key))
	if err != nil {
		return "", err
	}
	cipherText := decode(cryptedText)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
