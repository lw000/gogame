syntax = "proto3";

package routersvr; //包名

// Router 微服务
service Router {
    rpc BindStream(stream RequestMessage) returns (stream ReponseMessage){}
}

// Message 数据格式
message RequestMessage {
    int32 serviceId = 1;
    string serviceVersion = 2;
    string cuuid = 3;
    string uuid = 4;
    int32 msgType = 5;
    bytes msg = 6;
}

message ReponseMessage {
    int32 serviceId = 1;
    string cuuid = 2;
    string uuid = 3;
    int32 msgType = 4;
    bytes msg = 5;
}