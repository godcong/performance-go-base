package libp2portcp

import (
	"context"
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
		buf := make([]byte, len(testData))
		s.Read(buf)
		s.Write(buf) // 回显数据
		s.Close()
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
		stream.Write(testData)
		buf := make([]byte, len(testData))
		stream.Read(buf)
		stream.Close()
	}
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
		s.Read(buf)
		s.Write(buf) // 回显数据
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

		stream.Write(testData)
		buf := make([]byte, len(testData))
		stream.Read(buf)
		stream.Close()
	}
}

func BenchmarkTCPTransferOneConnect(b *testing.B) {
	// 启动TCP服务器
	listener, err := net.Listen("tcp", "127.0.0.1:0")
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
			buf := make([]byte, len(testData))
			conn.Read(buf)
			conn.Write(buf) // 回显数据
			conn.Close()
		}
	}()
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		b.Fatal(err)
	}
	// 运行基准测试
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conn.Write(testData)
		buf := make([]byte, len(testData))
		conn.Read(buf)
	}
	conn.Close()
}

func BenchmarkTCPTransferMultiConnect(b *testing.B) {
	// 启动TCP服务器
	listener, err := net.Listen("tcp", "127.0.0.1:0")
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
			buf := make([]byte, len(testData))
			conn.Read(buf)
			conn.Write(buf) // 回显数据
			conn.Close()
		}
	}()

	// 运行基准测试
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", listener.Addr().String())
		if err != nil {
			b.Fatal(err)
		}
		conn.Write(testData)
		buf := make([]byte, len(testData))
		conn.Read(buf)
		conn.Close()
	}
}
