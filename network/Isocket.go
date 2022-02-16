package network

import "net"

type Socket struct {
	m_Conn				net.Conn
	m_Port				int
	m_IP				string
}

func (this *Socket) SetTcpConn(conn net.Conn){
	this.m_Conn = conn
}

func (this *Socket) Close() {
	if this.m_Conn != nil {
		this.m_Conn.Close()
	}
}


