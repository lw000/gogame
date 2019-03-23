package ggwhitelist

import (
	"errors"
	"net"
	"net/http"
	"sync"
)

var (
	_whiteListServer *WhiteList
	_whiteListOnce   sync.Once
)

func WhiteListSrv() *WhiteList {
	_whiteListOnce.Do(func() {
		_whiteListServer = &WhiteList{ips: make(map[string]bool), errMsg: "illegal access"}
	})
	return _whiteListServer
}

type WhiteList struct {
	ips    map[string]bool
	errMsg string
	m      sync.RWMutex
}

func (wls *WhiteList) ErrMsg() string {
	return wls.errMsg
}

func (wls *WhiteList) SetErrMsg(errMsg string) {
	wls.errMsg = errMsg
}

func (wls *WhiteList) SetIp(ips []string) {
	wls.m.Lock()
	defer wls.m.Unlock()

	for _, v := range ips {
		wls.ips[v] = true
	}
}

func (wls *WhiteList) GetIp() []string {
	wls.m.Lock()
	defer wls.m.Unlock()

	var ips []string
	for k, ok := range wls.ips {
		if ok {
			ips = append(ips, k)
		}
	}
	return ips
}

func (wls *WhiteList) CheckWhiteList(w http.ResponseWriter, r *http.Request) error {
	if len(wls.ips) == 0 {
		return nil
	}

	clientIp, _, _ := net.SplitHostPort(r.RemoteAddr)
	if len(clientIp) == 0 {
		clientIp = r.Header.Get("Remote_addr")
	}

	v, ok := wls.ips[clientIp]
	if ok {
		return nil
	}

	if v {
		return nil
	}

	return errors.New(wls.ErrMsg())
}
