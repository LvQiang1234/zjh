package network

import (
	"fmt"
	"net"
	"zjh/log"
)

type ClientSocket struct{
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

func (this *ClientSocket) Run() bool {
	loop := func() bool {
		
	}
}

func (this *ClientSocket) Start() bool {
	if this.m_IP == "" {
		this.m_IP = "127.0.0.1"
	}
	if this.Connect() {
		this.m_Conn.(*net.TCPConn).SetNoDelay(true)
		go this.Run()
	}
}
