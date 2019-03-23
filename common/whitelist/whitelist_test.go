package tywhitelist

import (
	"log"
	"testing"
)

func TestWhiteListSrv(t *testing.T) {
	white := WhiteListSrv()
	white.SetIp([]string{"127.0.0.1", "192.168.1.73"})
	white.SetErrMsg("whitelist error context")
	log.Println(white.GetIp())
	log.Println(white.ErrMsg())
}
