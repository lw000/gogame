package ggtest

import (
	"log"
	"testing"
)

func TestTExec(t *testing.T) {
	v, er := TExec("", func() (i interface{}, e error) {
		return "ok", nil
	})
	if er != nil {
		log.Panic(er)
	}
	log.Println(v)
}
