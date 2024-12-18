package libp2portcp

import (
	"context"
	"fmt"
	"io"
	randv2 "math/rand/v2"
	"net"
	"testing"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

func BenchmarkLibp2pTransferOneStream(b *testing.B) {
	// 创建两个libp2p节点
	h1, err := libp2p.New()
	if err != nil {
		b.Fatal(err)
	}
	defer h1.Close()

	h2, err := libp2p.New()
	if err != nil {
		b.Fatal(err)
	}
	defer h2.Close()

	// 准备测试数据
	testData := make([]byte, 1024) // 1KB的测试数据
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// 设置协议处理器
	protocol := protocol.ID("/benchmark/1.0.0")
	h2.SetStreamHandler(protocol, func(s network.Stream) {
		for i := 0; i < b.N; i++ {
			buf := make([]byte, len(testData))
			read, err := s.Read(buf)
			if err != nil && err != io.EOF {
				b.Error(err)
				return
			}
			write, err := s.Write(buf)
			if err != nil && err != io.EOF {
				b.Error(err)
				return
			} // 回显数据
			if read != write {
				fmt.Println("read!= write")
			}
		}
		_ = s.Close()
	})

	// 获取节点地址信息
	h2Info := peer.AddrInfo{
		ID:    h2.ID(),
		Addrs: h2.Addrs(),
	}
	// 连接节点
	if err := h1.Connect(context.Background(), h2Info); err != nil {
		b.Fatal(err)
	}
	stream, err := h1.NewStream(context.Background(), h2.ID(), protocol)
	if err != nil {
		b.Fatal(err)
	}
	// 运行基准测试
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		write, err := stream.Write(testData)
		if err != nil && err != io.EOF {
			b.Error(err)
			return
		}
		buf := make([]byte, len(testData))
		read, err := stream.Read(buf)
		if err != nil && err != io.EOF {
			b.Error(err)
			return
		}
		if read != write {
			fmt.Println("read!= write")
		}
	}
	stream.Close()
}

func BenchmarkLibp2pTransferMultiConnectStream(b *testing.B) {
	// 创建两个libp2p节点
	h1, err := libp2p.New()
	if err != nil {
		b.Fatal(err)
	}
	defer h1.Close()

	h2, err := libp2p.New()
	if err != nil {
		b.Fatal(err)
	}
	defer h2.Close()

	// 准备测试数据
	testData := make([]byte, 1024) // 1KB的测试数据
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// 设置协议处理器
	protocol := protocol.ID("/benchmark/1.0.0")
	h2.SetStreamHandler(protocol, func(s network.Stream) {
		buf := make([]byte, len(testData))
		read, err := s.Read(buf)
		if err != nil && err != io.EOF {
			b.Error(err)
			return
		}
		write, err := s.Write(buf)
		if err != nil && err != io.EOF {
			b.Error(err)
			return
		} // 回显数据
		if read != write {
			fmt.Println("read != write")
		}
		s.Close()
	})

	// 运行基准测试
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 获取节点地址信息
		h2Info := peer.AddrInfo{
			ID:    h2.ID(),
			Addrs: h2.Addrs(),
		}
		// 连接节点
		if err := h1.Connect(context.Background(), h2Info); err != nil {
			b.Fatal(err)
		}
		stream, err := h1.NewStream(context.Background(), h2.ID(), protocol)
		if err != nil {
			b.Fatal(err)
		}

		write, err := stream.Write(testData)
		if err != nil && err != io.EOF {
			b.Error(err)
			return
		}
		buf := make([]byte, len(testData))
		read, err := stream.Read(buf)
		if err != nil && err != io.EOF {
			b.Error(err)
			return
		}
		if read != write {
			fmt.Println("read!= write")
		}
		stream.Close()
	}
}

func BenchmarkTCPTransferOneConnect(b *testing.B) {

	// Generate a random port between 10000 and 20000
	port := 10000 + randv2.Int32N(20000)
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	// 启动TCP服务器
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		b.Fatal(err)
	}
	defer listener.Close()

	// 准备测试数据
	testData := make([]byte, 1024) // 1KB的测试数据
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}

			for i := 0; i < b.N; i++ {
				buf := make([]byte, len(testData))
				read, err := conn.Read(buf)
				if err != nil && err != io.EOF {
					b.Error(err)
					return
				}
				write, err := conn.Write(buf)
				if err != nil && err != io.EOF {
					b.Error(err)
					return
				} // 回显数据
				if read != write {
					fmt.Println("read!= write")
				}
			}
			_ = conn.Close()

		}
	}()
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		b.Fatal(err)
	}
	// 运行基准测试
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		write, err := conn.Write(testData)
		if err != nil && err != io.EOF {
			b.Error(err)
			return
		}
		buf := make([]byte, len(testData))
		read, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			b.Error(err)
			return
		}
		if read != write {
			fmt.Println("read!= write")
		}
	}
	_ = conn.Close()
}

func BenchmarkTCPTransferMultiConnect(b *testing.B) {
	// Generate a random port between 10000 and 20000
	port := 10000 + randv2.Int32N(20000)
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	// 启动TCP服务器
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		b.Fatal(err)
	}
	defer listener.Close()

	// 准备测试数据
	testData := make([]byte, 1024) // 1KB的测试数据
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer c.Close()
				buf := make([]byte, len(testData))
				read, err := c.Read(buf)
				if err != nil && err != io.EOF {
					b.Error(err)
					return
				}
				write, err := c.Write(buf)
				if err != nil && err != io.EOF {
					b.Error(err)
					return
				} // 回显数据
				if read != write {
					fmt.Println("read!= write")
				}
			}(conn)
		}
	}()

	// 运行基准测试
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", listener.Addr().String())
		if err != nil {
			b.Fatal(err)
		}
		write, err := conn.Write(testData)
		if err != nil && err != io.EOF {
			b.Error(err)
			return
		}
		buf := make([]byte, len(testData))
		read, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			b.Error(err)
			return
		}
		if read != write {
			fmt.Println("read!= write")
		}
		_ = conn.Close()
	}
}
