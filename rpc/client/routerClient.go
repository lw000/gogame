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
	ServiceId int32
	UUID      string
	conn      *grpc.ClientConn
	client    routersvr.RouterClient
	status    int
	m         sync.RWMutex
	streams   []routersvr.Router_ForwardingDataStreamClient
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
	serviceId int32
	uuid      string
	onMessage func(response *routersvr.ForwardMessage)
	stream    routersvr.Router_ForwardingDataStreamClient
}

func (r *RpcRouterClient) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Error("did not connect: %v", er)
		return er
	}
	r.client = routersvr.NewRouterClient(r.conn)

	return nil
}

func (r *RpcRouterClient) RegisterService(protocols []*routersvr.RouterProtocol) error {
	ctx := context.Background()
	reply, er := r.client.RegisterService(ctx, &routersvr.RequestRegisterService{ServiceId: r.ServiceId, ServiceName: "platform", ServiceVersion: "1.0.1", Protocols: protocols})
	if er != nil {
		log.Error("did not connect:%v", er)
		return er
	}

	if reply.Status != 1 {
		log.Error(reply)
	}

	return nil
}

func (r *RpcRouterClient) ForwardingData(uuid string, msg []byte) (*routersvr.ForwardMessage, error) {
	ctx := context.Background()
	return r.client.ForwardingData(ctx, &routersvr.ForwardMessage{ServiceId: r.ServiceId, Uuid: uuid, Msg: msg})
}

func (r *RpcRouterClient) Stop() error {
	er := r.conn.Close()
	if er != nil {

	}
	return er
}

func (r *RpcRouterClient) CreateStream(onMessage func(response *routersvr.ForwardMessage)) (*RpcRouterStream, error) {
	var er error
	rpcStream := &RpcRouterStream{serviceId: r.ServiceId, uuid: ggutilty.UUID(), onMessage: onMessage}
	rpcStream.stream, er = r.client.ForwardingDataStream(context.Background())
	if er != nil {
		log.Error(er)
		return nil, er
	}
	r.streams = append(r.streams, rpcStream.stream)

	go rpcStream.run()

	return rpcStream, nil
}

func (r *RpcRouterStream) SendMessage(uuid string, msg []byte) error {
	if er := r.stream.Send(&routersvr.ForwardMessage{ServiceId: r.serviceId, Uuid: uuid, Msg: msg}); er != nil {
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
	var resp *routersvr.ForwardMessage
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
