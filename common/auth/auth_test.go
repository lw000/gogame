package ggauth

import (
	"log"
	"testing"
)

func TestB64Encode(t *testing.T) {
	s := B64Encode([]byte("1233123123123123"))
	log.Println(s)
}

func TestB64Decode(t *testing.T) {

}

func TestMD5(t *testing.T) {
	s, er := MD5([]byte("1111111111111111111111111"))
	if er != nil {
		log.Println(er)
	}
	log.Println(s)
}

func TestDesEcbEncrypt(t *testing.T) {

}

func TestDesEcbDecrypt(t *testing.T) {

}

func TestSha1(t *testing.T) {
	s := Sha1([]byte("1111111111111111111111111111111"))
	log.Println(s)
}

func TestSha224(t *testing.T) {
	s := Sha224([]byte("1111111111111111111111111111111"))
	log.Println(s)
}

func TestSha256(t *testing.T) {
	s := Sha256([]byte("1111111111111111111111111111111"))
	log.Println(s)
}

func TestSha512(t *testing.T) {
	s := Sha512([]byte("1111111111111111111111111111111"))
	log.Println(s)
}
