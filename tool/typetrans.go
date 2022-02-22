package tool

import "encoding/binary"

//字节转换成整形
func BytesToInt(data []byte) int {
	buff := make([]byte, 4)
	copy(buff, data)
	tmp := int32(binary.LittleEndian.Uint32(buff))
	return int(tmp)
}
