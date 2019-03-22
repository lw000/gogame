syntax = "proto3";

package platform; //包名

// Chat 微服务
service Game {
    rpc BidStream(stream Request) returns (stream Response){}
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