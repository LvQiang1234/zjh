package network

import (
	"net"
	"zjh/pb"
)

const (
	TCP_HEAD_SIZE        = 4     //包头的长度
	ReceiveBufferSize    = 1024  //一次接受数据缓存的最大值
	MaxReceiveBufferSize = 10240 //整个缓冲区的最大值
)

type NetMsg struct {
	playerId uint
	data     []byte
}

type Socket struct {
	m_Conn net.Conn
	m_Port int
	m_IP   string

	m_ReceiveBufferSize    int //一次接受数据缓存的最大值
	m_MaxReceiveBuffer     []byte
	m_MaxReceiveBufferSize int //整个缓冲区的最大值

	sendChan chan *pb.MsgPacket //发送消息的管道
}

func (this *Socket) Init(ip string, port int) bool {
	this.m_IP = ip
	this.m_Port = port

	this.sendChan = make(chan *pb.MsgPacket, 0)

	return true
}

func (this *Socket) SetTcpConn(conn net.Conn) {
	this.m_Conn = conn
}

func (this *Socket) Close() {
	if this.m_Conn != nil {
		this.m_Conn.Close()
	}
}
