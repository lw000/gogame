package rpcclient

import (
	"demo/gogame/proto/db"
	log "github.com/alecthomas/log4go"
	"google.golang.org/grpc"
	"io"
)

type RpcClientInterface interface {
	Stream()
}

type RpcClient struct {
	conn      *grpc.ClientConn
	onMessage func(response *dbsvr.Response)
}

func (r *RpcClient) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Error("did not connect:%v", er)
		return er
	}

	return nil
}

func (r *RpcClient) NewClient(f func(conn *grpc.ClientConn)) {
	f(r.conn)
}

func (r *RpcClient) Stop() {

}

func (r *RpcClient) run() {
	var er error
	var resp *dbsvr.Response
	for {
		//resp, er = r.stream.Recv()
		if er == io.EOF {
			log.Error("接收到服务端的结束信号 %v", er)
			break
		}

		if er != nil {
			log.Error("接收数据错误 %v", er)
			break
		}

		if r.onMessage != nil {
			r.onMessage(resp)
		}
	}
}
