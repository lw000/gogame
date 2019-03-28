cd ../../../../../
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=windows
cd src/demo/gogame/cmd/routerserv
go build -v -ldflags="-s -w"