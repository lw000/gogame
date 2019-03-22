// gogame project main.go
package main

import (
	"bufio"
	"bytes"
	"demo/gogame/packet"
	"demo/gogame/platform"
	"fmt"
	// "github.com/golang/protobuf/proto"
	// "github.com/golang/protobuf/protoc-gen-go"
)

func main() {
	p := platform.NewPlatform(1, "levi")
	if p != nil {
		p.CreateRoom()
	}

	b := packet.NewNetBuffer()
	b.Add([]byte("1234567890"))
	b.Add([]byte("1234567890"))
	b.Add([]byte("1234567890"))
	b.Add([]byte("1234567890"))
	b.Add([]byte("1234567890"))

	d := b.Read(5)
	fmt.Println(string(d))

	d1 := b.Read(5)
	fmt.Println(string(d1))

	d2 := b.Read(5)
	fmt.Println(string(d2))

	d3 := b.Read(5)
	fmt.Println(string(d3))

	var data []byte
	data = []byte("1234567890123456789012345678901234567890")
	{
		ssss := bytes.NewReader(data)
		r := bufio.NewReader(ssss)
		vv, er := r.Peek(5)
		if er != nil {

		}
		bbb := make([]byte, 5)
		n, er := r.Read(bbb)
		fmt.Println(bbb[0:n])
		fmt.Println(vv)
	}
	//{
	//	bw := bytes.new(data)
	//	w := bufio.NewWriter(bw)
	//	vv, er := w.Write([]byte("22222222222222222"))
	//	if er != nil {
	//
	//	}
	//}
}
