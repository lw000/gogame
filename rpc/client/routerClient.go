package rpcclient

import (
	"demo/gogame/protos/router"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"sync"
)

type RpcRouterClient struct {
	ServiceId      int32
	ServiceName    string
	ServiceVersion string
	UUID           string
	conn           *grpc.ClientConn
	client         routersvr.RouterClient
	status         int
	m              sync.RWMutex
	streams        []*RpcRouterStreamClient
}

type RpcRouterStreamClient struct {
	clientUuid string
	onMessage  func(*routersvr.ReponseMessage)
	client     *RpcRouterClient
	stream     routersvr.Router_BindStreamClient
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

func (r *RpcRouterClient) Start(address string) error {
	var err error
	// auth := AuthItem{
	// 	Username:"11111",
	// 	Password:"22222",
	// }
	// r.conn, err = grpc.Dial(address, grpc.WithPerRPCCredentials(&auth))
	r.conn, err = grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Error("did not connect: %v", err)
		return err
	}
	r.client = routersvr.NewRouterClient(r.conn)

	return nil
}

func (r *RpcRouterClient) Stop() error {
	var err error
	for _, s := range r.streams {
		err = s.CloseSend()
		if err != nil {
			log.Error(err)
		}
	}
	err = r.conn.Close()
	if err != nil {
		log.Error(err)
	}
	return err
}

func (r *RpcRouterClient) CreateStream(onMessage func(resp *routersvr.ReponseMessage)) (*RpcRouterStreamClient, error) {
	var err error
	rpcStream := &RpcRouterStreamClient{client: r, onMessage: onMessage}
	rpcStream.stream, err = r.client.BindStream(context.Background())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	r.streams = append(r.streams, rpcStream)

	go rpcStream.run()

	return rpcStream, nil
}

func (r *RpcRouterStreamClient) ClientUuid() string {
	return r.clientUuid
}

func (r *RpcRouterStreamClient) SetClientUuid(clientUuid string) {
	r.clientUuid = clientUuid
}

func (r *RpcRouterStreamClient) RegisterService(msg []byte) error {
	if err := r.stream.Send(&routersvr.RequestMessage{ServiceId: r.client.ServiceId, Cuuid: r.clientUuid, Uuid: "", MsgType: 0, Msg: msg}); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *RpcRouterStreamClient) SendMessage(uuid string, msg []byte) error {
	if err := r.stream.Send(&routersvr.RequestMessage{ServiceId: r.client.ServiceId, Cuuid: r.clientUuid, Uuid: uuid, MsgType: 1, Msg: msg}); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *RpcRouterStreamClient) CloseSend() error {
	return r.stream.CloseSend()
}

func (r *RpcRouterStreamClient) run() {
	var err error
	var resp *routersvr.ReponseMessage
	for {
		resp, err = r.stream.Recv()
		if err == io.EOF {
			log.Error("接收到服务端的结束信号 %v", err)
			break
		}

		if err != nil {
			log.Error("接收数据错误 %v", err)
			break
		}

		if r.onMessage != nil {
			r.onMessage(resp)
		}
	}
}
