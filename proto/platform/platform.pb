syntax = "proto3";

package platformsvr; //包名

// Chat 微服务
service Platform {
    rpc BidStream(stream Request) returns (stream Response){}
    rpc RegisterService(RequestRegisterService) returns (ResponseRegisterService){}
}

// Request 请求数据格式
message Request {
    int32 mainId = 1;
    int32 subId = 2;
    int32 requestId = 3;
    string input = 4;
}

// Response 响应数据格式
message Response {
    int32 mainId = 1;
    int32 subId = 2;
    int32 requestId = 3;
    string output = 4;
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