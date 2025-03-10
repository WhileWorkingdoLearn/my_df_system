package p2p

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
)

type BinaryMsg struct {
	MsgType [1]byte
	Length  [2]byte
	Payload []byte
	Cheksum [4]byte
}

/*calculateCRC32*/
func calcChecksum(data []byte) [4]byte {
	checksum := crc32.ChecksumIEEE(data)
	var checksumBytes [4]byte
	binary.BigEndian.PutUint32(checksumBytes[:], checksum)
	return checksumBytes
}

func ConvertToBinary(msgType byte, payload []byte) []byte {
	lenght := uint16(len(payload))
	var conentLength [2]byte
	binary.BigEndian.PutUint16(conentLength[:], lenght)

	buff := bytes.Buffer{}
	buff.WriteByte(msgType)
	buff.Write(conentLength[:])
	buff.Write(payload)
	checksum := calcChecksum(buff.Bytes())
	buff.Write(checksum[:])

	return buff.Bytes()
}

func ConvertFromBinary(data []byte) (*BinaryMsg, error) {
	if len(data) < 7 {
		return nil, fmt.Errorf("data to short, must be at least 7 bytes")
	}

	msg := &BinaryMsg{}
	copy(msg.MsgType[:], data[:1])
	copy(msg.Length[:], data[1:3])

	dataLength := binary.BigEndian.Uint16(msg.Length[:])
	if len(data) < int(7 + +dataLength) {
		return nil, fmt.Errorf("invalid data, payload is to short")
	}

	msg.Payload = make([]byte, dataLength)
	copy(msg.Payload, data[3:3+dataLength])

	copy(msg.Cheksum[:], data[3+dataLength:])

	check := calcChecksum(data[:3+dataLength])
	if !bytes.Equal(check[:], msg.Cheksum[:]) {
		return nil, fmt.Errorf("invalid checksum! exp: %v got %v", check, msg.Cheksum)
	}

	return msg, nil

}
