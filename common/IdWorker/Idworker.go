package tyIdWorker

import (
	"sync"

	// "github.com/gitstliu/go-id-worker"
	// "github.com/bwmarrin/snowflake"
	"github.com/zheng-ji/goSnowFlake"
)

// sonyflake 是 Sony 公司的一个开源项目，基本思路和 snowflake 差不多，不过位分配上稍有不同：

// +-----------------------------------------------------------------------------+
// | 1 Bit Unused | 39 Bit Timestamp |  8 Bit Sequence ID  |   16 Bit Machine ID |
// +-----------------------------------------------------------------------------+
// 这里的时间只用了 39 个 bit，但时间的单位变成了 10ms，所以理论上比 41 位表示的时间还要久(174 years)。

type IdWorker struct {
	worker *goSnowFlake.IdWorker
}

type IdworkerServer interface {
	InitIdWorker(workerid int64)
	NextId() int64
}

var (
	_idworker *IdWorker
	_m        sync.Mutex
)

func init() {

}

func IdworkerServ() *IdWorker {
	if _idworker == nil {
		_m.Lock()
		defer _m.Unlock()

		if _idworker == nil {
			_idworker = &IdWorker{}
		}
	}
	return _idworker
}

func (iw *IdWorker) Start(workerid int64) bool {
	// iw.worker.InitIdWorker(iw.workerid, 1)
	var (
		err error
	)
	iw.worker, err = goSnowFlake.NewIdWorker(workerid)
	if err != nil {
		return false
	}

	return true
}

func (iw *IdWorker) NewId() int64 {
	newId, err := iw.worker.NextId()
	if err != nil {
		return -1
	}
	return newId
}

// type IdWorker struct {
// 	worker *idworker.IdWorker
// }

// type IdworkerServer interface {
// 	InitIdWorker(workerid int64)
// 	ProductId() int64
// }

// var (
// 	ins *IdWorker
// 	m   sync.Mutex
// )

// func SharedIdworkerInstance() *IdWorker {
// 	if ins == nil {
// 		m.Lock()
// 		defer m.Unlock()

// 		if ins == nil {
// 			ins = &IdWorker{
// 				worker: &idworker.IdWorker{},
// 			}
// 		}
// 	}
// 	return ins
// }

// func (iw *IdWorker) InitIdWorker(workerid int64) {
// 	iw.worker.InitIdWorker(iw.workerid, 1)
// }

// func (iw *IdWorker) ProductId() int64 {
// 	newId, err := iw.worker.NextId()
// 	if err != nil {
// 		log.Printf(err.Error())
// 		return -1
// 	}
// 	return newId
// }
