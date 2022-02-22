package network

import (
	"fmt"
	"net"
	"zjh/tool"
)

const (
	TCP_HEAD_SIZE = 4
)

type Socket struct {
	m_Conn net.Conn
	m_Port int
	m_IP   string

	m_ReceiveBufferSize    int //一次接受数据缓存的最大值
	m_MaxReceiveBuffer     []byte
	m_MaxReceiveBufferSize int //整个缓冲区的最大值

	m_ClientId uint64
}

func (this *Socket) SetTcpConn(conn net.Conn) {
	this.m_Conn = conn
}

func (this *Socket) Close() {
	if this.m_Conn != nil {
		this.m_Conn.Close()
	}
}

func (this *Socket) HandlePacket(Id uint64, dat []byte) {

}

func (this *Socket) ReceivePacket(Id uint64, dat []byte) bool {
	//找包结束
	seekToTcpEnd := func(buff []byte) (bool, int) {
		nLen := len(buff)
		if nLen < TCP_HEAD_SIZE {
			return false, 0
		}

		nSize := tool.BytesToInt(buff[0:4])
		if nSize+TCP_HEAD_SIZE <= nLen {
			return true, nSize + TCP_HEAD_SIZE //返回一个整包的长度
		}
		return false, 0
	}

	buff := append(this.m_MaxReceiveBuffer, dat...)
	this.m_MaxReceiveBuffer = []byte{}
	nCurSize := 0
	//fmt.Println(this.m_MaxReceiveBuffer)
ParsePacekt:
	nPacketSize := 0
	nBufferSize := len(buff[nCurSize:])
	bFindFlag := false
	bFindFlag, nPacketSize = seekToTcpEnd(buff[nCurSize:])
	//fmt.Println(bFindFlag, nPacketSize, nBufferSize)
	if bFindFlag {
		if nBufferSize == nPacketSize { //完整包
			this.HandlePacket(Id, buff[nCurSize+TCP_HEAD_SIZE:nCurSize+nPacketSize])
			nCurSize += nPacketSize
		} else if nBufferSize > nPacketSize { //大于一个完整包
			this.HandlePacket(Id, buff[nCurSize+TCP_HEAD_SIZE:nCurSize+nPacketSize])
			nCurSize += nPacketSize
			goto ParsePacekt
		}
	} else if nBufferSize < this.m_MaxReceiveBufferSize {
		this.m_MaxReceiveBuffer = buff[nCurSize:]
	} else {
		fmt.Println("超出最大包限制，丢弃该包")
		return false
	}
	return true
}
