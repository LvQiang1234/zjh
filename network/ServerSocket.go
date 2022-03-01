package network

import (
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
)

type ServerSocket struct {
	Socket
	m_nClientCount int
	m_nIdSeed      uint64
	m_ClientList   map[uint64]*ClientSocket
	m_ClientLocker *sync.RWMutex
	m_Listen       *net.TCPListener
	m_Lock         sync.Mutex
}

func (this *ServerSocket) Init(ip string, port int) bool {
	this.Socket.Init(ip, port)
	this.m_ClientList = make(map[uint64]*ClientSocket)
	this.m_ClientLocker = &sync.RWMutex{}
	return true
}

func (this *ServerSocket) AssignClientId() uint64 {
	return atomic.AddUint64(&this.m_nIdSeed, 1)
}

func (this *ServerSocket) AddClinet(tcpConn *net.TCPConn, addr string) *ClientSocket {
	client := &ClientSocket{}
	if client != nil {
		client.Init("", 0)
		client.m_ReceiveBufferSize = this.m_ReceiveBufferSize
		client.m_MaxReceiveBufferSize = this.m_MaxReceiveBufferSize
		client.m_ClientId = this.AssignClientId()
		client.m_IP = addr
		client.SetTcpConn(tcpConn)
		this.m_ClientLocker.Lock()
		this.m_ClientList[client.m_ClientId] = client
		this.m_ClientLocker.Unlock()
		client.Start()
		this.m_nClientCount++
		return client
	} else {
		log.Printf("%s", "无法创建客户端连接对象")
	}
	return nil
}

func (this *ServerSocket) handleConn(tcpConn *net.TCPConn, addr string) bool {
	if tcpConn == nil {
		return false
	}

	pClient := this.AddClinet(tcpConn, addr)
	if pClient == nil {
		return false
	}

	return true
}

func (this *ServerSocket) Run() bool {
	for {
		tcpConn, err := this.m_Listen.AcceptTCP()
		handleError(err)
		if err != nil {
			return false
		}

		fmt.Printf("客户端：%s已连接！\n", tcpConn.RemoteAddr().String())
		//延迟，关闭链接
		//defer tcpConn.Close()
		this.handleConn(tcpConn, tcpConn.RemoteAddr().String())
	}
}

func (this *ServerSocket) Start() bool {
	if this.m_IP == "" {
		this.m_IP = "127.0.0.1"
	}

	var strRemote = fmt.Sprintf("%s:%d", this.m_IP, this.m_Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", strRemote)
	if err != nil {
		log.Fatalf("%v", err)
	}
	ln, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		log.Fatalf("%v", err)
		return false
	}

	fmt.Printf("启动监听，等待链接！\n")

	this.m_Listen = ln
	//延迟，监听关闭
	//defer ln.Close()
	go this.Run()
	return true
}
