package rpcclient

import (
	"demo/gogame/common/utilty"
	"demo/gogame/proto/router"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"syncthing/lib/sync"
)

type RpcRouterClient struct {
	conn    *grpc.ClientConn
	c       routersvr.RouterClient
	status  int
	m       sync.RWMutex
	streams []routersvr.Router_BidStreamClient
}

func (r *RpcRouterClient) Status() int {
	r.m.RLock()
	status := r.status
	r.m.RUnlock()

	return status
}

func (r *RpcRouterClient) SetStatus(status int) {
	r.m.Lock()
	defer r.m.Unlock()
	r.status = status
}

type RpcRouterStream struct {
	uuid      string
	onMessage func(response *routersvr.ForwardResponse)
	stream    routersvr.Router_BidStreamClient
}

func (r *RpcRouterClient) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Error("did not connect: %v", er)
		return er
	}
	r.c = routersvr.NewRouterClient(r.conn)
	return nil
}

func (r *RpcRouterClient) ForwardingData(mainId int32, subId int32, uuid string, input string) (*routersvr.ForwardResponse, error) {
	ctx := context.Background()
	return r.c.ForwardingData(ctx, &routersvr.ForwardRequest{ServiceId: 100, MainId: mainId, SubId: subId, Uuid: uuid, Input: input})
}

func (r *RpcRouterClient) Stop() error {
	er := r.conn.Close()
	if er != nil {

	}
	return er
}

func (r *RpcRouterClient) CreateStream(onMessage func(response *routersvr.ForwardResponse)) (*RpcRouterStream, error) {
	var er error
	rpcStream := &RpcRouterStream{uuid: ggutilty.UUID(), onMessage: onMessage}
	rpcStream.stream, er = r.c.BidStream(context.Background())
	if er != nil {
		log.Error(er)
		return nil, er
	}
	r.streams = append(r.streams, rpcStream.stream)

	go rpcStream.run()

	return rpcStream, nil
}

func (r *RpcRouterStream) SendMessage(mainId int32, subId int32, uuid string, input string) error {
	if er := r.stream.Send(&routersvr.ForwardRequest{MainId: mainId, SubId: subId, Uuid: uuid, Input: input}); er != nil {
		log.Error(er)
		return er
	}
	return nil
}

func (r *RpcRouterStream) CloseSend() error {
	return r.stream.CloseSend()
}

func (r *RpcRouterStream) run() {
	var er error
	var resp *routersvr.ForwardResponse
	for {
		resp, er = r.stream.Recv()
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
