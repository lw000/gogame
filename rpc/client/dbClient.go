package rpcclient

import (
	log "github.com/alecthomas/log4go"
	"gogame/protos/db"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
)

type RpcDbClient struct {
	conn   *grpc.ClientConn
	client dbsvr.DBClient
}

type RpcDbStream struct {
	onMessage func(*dbsvr.Response)
	stream    dbsvr.DB_BidStreamClient
}

func (r *RpcDbClient) Start(address string) error {
	var err error
	r.conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Error("did not connect:%v", err)
		return err
	}

	r.client = dbsvr.NewDBClient(r.conn)

	return nil
}

func (r *RpcDbClient) Stop() error {
	err := r.conn.Close()
	if err != nil {

	}
	return err
}

func (r *RpcDbClient) CreateStream(onMessage func(response *dbsvr.Response)) (*RpcDbStream, error) {
	var err error
	var rpcStream RpcDbStream
	rpcStream.stream, err = r.client.BidStream(context.Background())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	rpcStream.onMessage = onMessage

	go rpcStream.run()

	return &rpcStream, nil
}

func (r *RpcDbStream) SendMessage(mainId int32, subId int32, requestId int32, input string) error {
	if err := r.stream.SendMsg(&dbsvr.Request{MainId: mainId, SubId: subId, RequestId: requestId, Input: input}); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *RpcDbStream) CloseSend() error {
	return r.stream.CloseSend()
}

func (r *RpcDbStream) run() {
	var err error
	var resp *dbsvr.Response
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
