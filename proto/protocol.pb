syntax = "proto3";

package protocol; //包名

message RequestLogin {
    string username = 1;
    string userpsd = 2;
}

message ResponseLogin {
    int32 status = 1;
    string msg = 2;
}

message RequestChat {
    string uid = 1;
    string msg = 2;
}

message ResponseChat {
    string msg = 2;
}