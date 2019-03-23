package tycache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

type SmsCodeContext struct {
	Code  int32  // 短信验证码
	Token string // 用户登录token
}

type SmsCodeCache struct {
	c *cache.Cache
}

var (
	ins  *SmsCodeCache
	once sync.Once
)

func SmsCacheInsance() *SmsCodeCache {
	once.Do(func() {
		ins = &SmsCodeCache{
			c: cache.New(time.Minute*time.Duration(3), time.Minute*time.Duration(10)),
		}
	})

	return ins
}

func (this *SmsCodeCache) Add(token string, code int32) int {
	if len(token) == 0 {
		return -1
	}

	if code <= 0 {
		return -1
	}

	var context = SmsCodeContext{
		Token: token,
		Code:  code,
	}

	this.c.Set(token, context, cache.DefaultExpiration)

	return 0
}

func (this *SmsCodeCache) Remove(token string) {
	if len(token) == 0 {
		return
	}

	this.c.Delete(token)
}

func (this *SmsCodeCache) Find(token string) int32 {
	if len(token) == 0 {
		return -1
	}

	v, found := this.c.Get(token)
	if found {
		return v.(SmsCodeContext).Code
	}

	return 0
}
