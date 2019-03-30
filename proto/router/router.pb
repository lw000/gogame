syntax = "proto3";

package routersvr; //包名

// Router 微服务
service Router {
    rpc BindStream(stream RequestMessage) returns (stream ReponseMessage){}
}

// Message 数据格式
message RequestMessage {
    int32 serviceId = 1;
    string cuuid = 2;
    string uuid = 3;
    int32 msgType = 4;
    bytes msg = 5;
}

message ReponseMessage {
    int32 serviceId = 1;
    string cuuid = 2;
    string uuid = 3;
    int32 msgType = 4;
    bytes msg = 5;
}