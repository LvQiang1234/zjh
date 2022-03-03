package tool

import "encoding/binary"

//字节转换成整形
func BytesToInt(data []byte) int {
	buff := make([]byte, 4)
	copy(buff, data)
	tmp := int32(binary.LittleEndian.Uint32(buff))
	return int(tmp)
}

//整形转换成字节
func UintToBytes(val uint) []byte {
	tmp := uint32(val)
	buff := make([]byte, 4)
	binary.LittleEndian.PutUint32(buff, tmp)
	return buff
}

//字节型转换成int16
func BytesToUint16(data []byte) uint16 {
	buff := make([]byte, 2)
	copy(buff, data)
	tmp := binary.LittleEndian.Uint16(buff)
	return uint16(tmp)
}
