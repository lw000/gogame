syntax = "proto3";

package routersvr; //包名

// Router 微服务
service Router {
    rpc RegisterService(RequestRegisterService) returns (ResponseRegisterService){}
    rpc ForwardingData(ForwardMessage) returns (ForwardMessage){}
    rpc ForwardingDataStream(stream ForwardMessage) returns (stream ForwardMessage){}
}

// ForwardMessage 转发数据格式
message ForwardMessage {
    int32 serviceId = 1;
    string uuid = 2;
    bytes msg = 3;
}

message RequestRegisterService {
    int32 serviceId = 1;
    string serviceName = 2;
    string serviceVersion = 3;
    string srviceProtocol = 4;
    repeated RouterProtocol protocols = 5;
}

message ResponseRegisterService {
    int32 status = 1;
    string uuid = 2;
    string msg = 3;
}

message RouterProtocol {
    int32 mainId = 2;
    int32 subId = 3;
}