syntax = "proto3";

package room; //包名

message RequestJoinRoom {
    int32 roomId = 1;
    string roomName = 2;
}

message ReponseJoinRoom {
    int32 roomId = 1;
    string roomName = 2;
}

message RequestLeaveRoom {
    int32 roomId = 1;
}

message ReponseLeaveRoom {
    int32 code = 1;
}