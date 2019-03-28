syntax = "proto3";

package routersvr; //包名

// Chat 微服务
service Router {
    rpc BidStream(stream ForwardRequest) returns (stream ForwardResponse){}
    rpc RegisterService(RequestRegisterService) returns (ResponseRegisterService){}
    rpc ForwardingData(ForwardRequest) returns (ForwardResponse){}
}

// ForwardRequest 请求数据格式
message ForwardRequest {
    int32 serviceId = 1;
    string uuid = 2;
    int32 mainId = 3;
    int32 subId = 4;
    string input = 6;
}

// ForwardResponse 响应数据格式
message ForwardResponse {
    int32 serviceId = 1;
    string uuid = 2;
    int32 mainId = 3;
    int32 subId = 4;
    string output = 6;
}

message RequestRegisterService {
    int32 serviceId = 1;
    string serviceName = 2;
    string serviceVersion = 3;
    string protocol = 4;
}

message ResponseRegisterService {
    int32 status = 1;
    string msg = 2;
}