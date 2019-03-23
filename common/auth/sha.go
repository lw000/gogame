package tyauth

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func Sha1(str []byte) string {
	if len(str) == 0 {
		return ""
	}

	h := sha1.New()
	h.Write(str)
	bs := h.Sum(nil)
	return /*fmt.Sprintf("%x", bs)*/ hex.EncodeToString(bs)
}

func Sha224(str []byte) string {
	if len(str) == 0 {
		return ""
	}

	h := sha256.New224()
	h.Write(str)
	bs := h.Sum(nil)
	return /*fmt.Sprintf("%x", bs)*/ hex.EncodeToString(bs)
}

func Sha256(str []byte) string {
	if len(str) == 0 {
		return ""
	}

	h := sha256.New()
	h.Write(str)
	bs := h.Sum(nil)
	return /*fmt.Sprintf("%x", bs)*/ hex.EncodeToString(bs)
}

func Sha512(str []byte) string {
	if len(str) == 0 {
		return ""
	}

	h := sha512.New()
	h.Write(str)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
