cd ../../../../../
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=linux
cd src/demo/gogame/cmd/test
go build -v -ldflags="-s -w"