cd ../../../../../
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=windows
cd src/demo/gogame/cmd/dbserv
go build -v -ldflags="-s -w"