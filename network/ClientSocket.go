package network

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"strconv"
	"strings"
	"zjh/log"
	"zjh/msgcenter"
	"zjh/pb"
	"zjh/tool"
)

type ClientSocket struct {
	Socket
}

func (this *ClientSocket) Init(ip string, port int) bool {
	arr := strings.Split(ip, ":")
	this.m_IP = arr[0]
	this.m_Port, _ = strconv.Atoi(arr[1])
	this.Socket.Init(ip, port)
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
			this.ReceivePacket(buff[:n])
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

func (this *ClientSocket) DoSend() {
	for {
		select {
		case response := <-this.sendChan:
			var buff []byte = make([]byte, 0)
			msg, _ := proto.Marshal(response)
			buff = append(buff, tool.UintToBytes(uint(len(msg)))...)
			buff = append(buff, msg...)
			this.m_Conn.Write(buff)
		}
	}
}

func (this *ClientSocket) Send(dat []byte, msgId pb.MsgId) {
	response := pb.MsgPacket{MsgId: &msgId, Data: dat}
	this.sendChan <- &response
}

func (this *ClientSocket) HandlePacket(data []byte) {
	msgPacket := pb.MsgPacket{}
	proto.Unmarshal(data, &msgPacket)
	handler := msgcenter.GetHandler(*msgPacket.MsgId)
	dat, msgType := handler(*msgPacket.PlayerId, msgPacket.Data)
	this.Send(dat, msgType)
}

func (this *ClientSocket) ReceivePacket(dat []byte) bool {
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
			this.HandlePacket(buff[nCurSize+TCP_HEAD_SIZE : nCurSize+nPacketSize])
			nCurSize += nPacketSize
		} else if nBufferSize > nPacketSize { //大于一个完整包
			this.HandlePacket(buff[nCurSize+TCP_HEAD_SIZE : nCurSize+nPacketSize])
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

func (this *ClientSocket) ServerClientStart() bool {
	this.m_Conn.(*net.TCPConn).SetNoDelay(true)
	go this.Run()
	go this.DoSend()
	return true
}

func (this *ClientSocket) ClientServerStart() bool {
	if this.m_IP == "" {
		this.m_IP = "127.0.0.1"
	}
	if this.Connect() {
		this.m_Conn.(*net.TCPConn).SetNoDelay(true)
		go this.Run()
		go this.DoSend()
	}
	return true
}
