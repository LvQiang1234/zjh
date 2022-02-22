package network

import (
	"fmt"
	"io"
	"net"
	"zjh/log"
	"zjh/tool"
)

type ClientSocket struct {
	Socket
}

func (this *ClientSocket) Init(ip string, port int) bool {
	if this.m_IP == ip || this.m_Port == port {
		return false
	}
	this.m_IP = ip
	this.m_Port = port
	return true
}

//客户端主动断开连接
func (this *ClientSocket) Disconnect() {
	this.Close()
}

func (this *ClientSocket) Connect() bool {
	address := fmt.Sprintf("%s:%d", this.m_IP, this.m_Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		log.Error("parsing ip err: %v", err)
	}
	conn, err1 := net.DialTCP("tcp4", nil, tcpAddr)
	if err1 != nil {
		return false
	}
	this.SetTcpConn(conn)
	return true
}

func handleError(err error) {
	if err == nil {
		return
	}
	log.Error("错误：%s\n", err.Error())
}

func (this *ClientSocket) Run() bool {
	var buff = make([]byte, this.m_ReceiveBufferSize)
	loop := func() bool {
		defer func() {
			if err := recover(); err != nil {
				tool.TraceCode(err)
			}
		}()
		n, err := this.m_Conn.Read(buff)
		if err == io.EOF {
			fmt.Printf("远程链接：%s已经关闭！\n", this.m_Conn.RemoteAddr().String())
			this.Disconnect()
			return false
		}
		if err != nil {
			handleError(err)
			this.Disconnect()
			return false
		}
		if n > 0 {
			this.ReceivePacket(this.m_ClientId, buff[:n])
		}
		return true
	}
	for {
		if !loop() {
			break
		}
	}

	this.Close()
	return true
}

func (this *ClientSocket) Start() bool {
	if this.m_IP == "" {
		this.m_IP = "127.0.0.1"
	}
	if this.Connect() {
		this.m_Conn.(*net.TCPConn).SetNoDelay(true)
		go this.Run()
	}
	return true
}
