syntax = "proto3";

package common; //包名

message RequestRegisterService {
    int32 serviceId = 1;
    string serviceName = 2;
    string serviceVersion = 3;
}

message ResponseRegisterService {
    int32 status = 1;
    string msg = 2;
}