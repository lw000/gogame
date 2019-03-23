package tyIdWorker

import (
	"log"
	"testing"
)

func TestIdworkerServ(t *testing.T) {
	idw := IdworkerServ()
	idw.Start(1)
	for i := 0; i < 100; i++ {
		log.Println("|", i, "|", idw.NewId())
	}
}
