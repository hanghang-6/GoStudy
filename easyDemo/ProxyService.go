package main

import (
	"bufio"
	"log"
	"net"
)

func process(conn net.Conn) {
	// 记得关闭连接
	defer conn.Close()
	// 带缓冲的 只读流  带缓冲的流可以减少底层系统调用的次数
	reader := bufio.NewReader(conn)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}
func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")

	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept err:%v\n", err)
			continue
		}
		//轻松高并发
		go process(client)
	}
}
