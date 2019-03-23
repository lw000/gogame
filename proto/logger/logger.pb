syntax = "proto3";

package Loggersvr; //包名

// Logger 微服务
service Logger {
    rpc WriteLogger(Request) returns (Response){}
}

// Request 请求数据格式
message Request {
    int32 serverId = 1;
    string msg = 4;
}

// Response 响应数据格式
message Response {
    int32 status = 1;
}


