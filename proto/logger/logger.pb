syntax = "proto3";

package loggersvr; //包名

// Logger 微服务
service Logger {
    rpc BidStream(stream Request) returns (stream Response){}
    rpc WriteLogger(Request) returns (Response){}
    rpc RegisterService(RequestRegisterService) returns (ResponseRegisterService){}
}

// Request 请求数据格式
message Request {
    int32 serverId = 1;
    string serverTag = 2;
    string msg = 3;
}

// Response 响应数据格式
message Response {
    int32 status = 1;
}

message RequestRegisterService {
    int32 serviceId = 1;
    string serviceName = 2;
    string serviceVersion = 3;
}

message ResponseRegisterService {
    int32 status = 1;
    string msg = 2;
}
