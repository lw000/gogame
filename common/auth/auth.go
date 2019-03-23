package ggauth

import (
	"crypto/md5"
	"crypto/rc4"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

func RC4(key []byte, str []byte) string {
	if len(key) == 0 {
		return ""
	}

	if len(str) == 0 {
		return ""
	}

	r, err := rc4.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	dst := make([]byte, len(str))
	r.XORKeyStream(dst, str)

	return base64.StdEncoding.EncodeToString(dst)
}

func MD5(str []byte) (string, error) {
	if len(str) == 0 {
		return "", errors.New("str is empty")
	}

	h := md5.New()
	_, err := h.Write(str)
	if err != nil {
		return "", err
	}

	md5str := hex.EncodeToString(h.Sum(nil))
	return md5str, nil
}

func B64Encode(src []byte) string {
	if len(src) == 0 {
		return ""
	}

	base64Str := base64.StdEncoding.EncodeToString(src)
	return base64Str
}

func B64Decode(str string) []byte {
	if len(str) == 0 {
		return nil
	}

	base64Str, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil
	}
	return base64Str
}
