package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const socks5Ver = 0x05

// commend = 1 表示 connection命令  即 代理服务器与目标服务建立新的TCP连接
const cmdBind = 0x01
const atypeIPV4 = 0x01
const atypeHOST = 0x03
const atypeIPV6 = 0x04

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD

	// 读版本号
	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	// 比较
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	method := make([]byte, methodSize)
	//将读到的填充到创建的切片中
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}
	log.Println("ver", ver, "method", method)

	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// +----+--------+

	// 认证响应   写入比特流是TCP中常见的通信方式
	_, err = conn.Write([]byte{socks5Ver, 0x00}) // 0x00代表不需要验证
	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}
	return nil
}

func connect(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER 版本号，socks5的值为0x05
	// CMD 0x01表示CONNECT请求
	// RSV 保留字段，值为0x00
	// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
	//   0x01表示IPv4地址，DST.ADDR为4个字节
	//   0x03表示域名，DST.ADDR是一个可变长度的域名
	// DST.ADDR 一个可变长度的值
	// DST.PORT 目标端口，固定2个字节

	// 大小为4byte的buffer
	buf := make([]byte, 4)
	// ReadFull函数 读满 buf 才罢手
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return fmt.Errorf("read header failed:%w", err)
	}
	// 对应协议的字段 其中 RSV保留字不予理会
	ver, cmd, atyp := buf[0], buf[1], buf[3]
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	if cmd != cmdBind {
		return fmt.Errorf("not supported cmd:%v", cmd)
	}
	addr := ""
	switch atyp {
	case atypeIPV4:
		// ipv4 地址长度为32bit=4byte  正好用buf接收
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read atyp failed:%w", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case atypeHOST:
		// 协议构成： 如果atypehost 那么 下一个字段的第一个byte存放 hostSize字段用于存放host长度  这是能直接读出来的
		hostSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("read hostSize failed:%w", err)
		}
		host := make([]byte, hostSize)
		_, err = io.ReadFull(reader, host)
		if err != nil {
			return fmt.Errorf("read host failed:%w", err)
		}
		addr = string(host)
	case atypeIPV6:
		return errors.New("IPv6: no supported yet")
	default:
		return errors.New("invalid atyp")
	}
	// 读两字节  port
	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed:%w", err)
	}
	port := binary.BigEndian.Uint16(buf[:2]) // 大端 . 无符号整数
	// 建立网络连接
	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil {
		// %w返回原始错误
		return fmt.Errorf("dial dst failed:%w", err)
	}
	defer dest.Close()
	// 打印一下从包中解析出来的 addr，port
	log.Println("dial", addr, port)

	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER socks版本，这里为0x05
	// REP Relay field,内容取值如下 X’00’ succeeded
	// RSV 保留字段
	// ATYPE 地址类型
	// BND.ADDR 服务绑定的地址
	// BND.PORT 服务绑定的端口DST.PORT
	// 发送初始数据包
	_, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	//创建上下文和取消函数
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 确保函数退出时清理上下文

	go func() {
		// 实现一个单向数据转发
		_, _ = io.Copy(dest, reader) // client->proxy->dest
		cancel()
	}()
	go func() {
		_, _ = io.Copy(conn, dest) // dest->proxy->client
		cancel()
	}()

	// 阻塞等待上下文完成   只要有一个方向的发送完成了就可以了，一次请求和产生的响应，可以看作两个单向数据流
	<-ctx.Done()
	return nil
}

func process(conn net.Conn) {
	// 记得关闭连接
	defer conn.Close()
	// 带缓冲的 只读流  带缓冲的流可以减少底层系统调用的次数
	// 可以从conn（转发服务器和client的TCP连接） 中读取字节流
	reader := bufio.NewReader(conn)
	err := auth(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed: %v", conn.RemoteAddr(), err)
		return
	}
	err = connect(reader, conn)
	if err != nil {
		log.Printf("client %v connect failed: %v", conn.RemoteAddr(), err)
		return
	}
}
func main() {
	// 在1080上创建tcp服务器
	server, err := net.Listen("tcp", "127.0.0.1:1080")

	if err != nil {
		panic(err)
	}
	for {
		// 接受传入的连接请求
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept err:%v\n", err)
			continue
		}
		//轻松高并发
		// 处理某一个客户端的连接请求  （同一时刻可能有多个客户端连接请求  故  进行并发处理）
		go process(client)
	}
}
