cd ../../../../../
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=linux
cd src/demo/gogame/cmd/gamesrv
go build -v -ldflags="-s -w"