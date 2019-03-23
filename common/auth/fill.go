package ggauth

import (
	"bytes"
)

func Pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	paddinig := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(paddinig)}, paddinig)
	return append(ciphertext, padtext...)
}

func Pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}
