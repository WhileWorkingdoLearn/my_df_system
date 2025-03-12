package transfer

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

/*
	Version     [1]byte
	MsgType     [1]byte
	Method      [6]byte
	Timestamp   [8]byte
	Timeout     [4]byte
	Domain      [32]byte
	Endpoint    [32]byte
	Auth        [16]byte
	PayloadType [1]byte
*/

func EncodeNodeMsgHeader(msgType, method IntByteType, timestamp, lifespan, domain, endpoint, auth StringByteType) ([]byte, error) {
	msg := make([]byte, 22)
	buff := [8]byte{}
	binary.LittleEndian.PutUint64(buff[:], uint64(1))
	copy(msg[OffsetByteVersion:], buff[:])

	//clear(buff[:])
	binary.LittleEndian.PutUint64(buff[:], uint64(msgType))
	copy(msg[OffsetByteMsgType:], buff[:])

	binary.LittleEndian.PutUint64(buff[:], uint64(method))
	copy(msg[OffsetByteMethod:], buff[:])

	//binary.LittleEndian.PutUint64(msg[OffsetByteMsgType:], uint64(method))

	/*
		tmstmp := timestamp.To8Bytes()
		msg = append(msg, tmstmp[:]...)

		duration := lifespan.To8Bytes()
		msg = append(msg, duration[:]...)

		dns := domain.To32Byte()
		msg = append(msg, dns[:]...)

		path := endpoint.To32Byte()
		msg = append(msg, path[:]...)

		authkey := auth.To16Byte()
		msg = append(msg, authkey[:]...)
	*/
	return msg, nil
}

func encodeTimestamp(t time.Time) [8]byte {
	timestamp := uint64(t.UnixNano())
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], timestamp)
	return buf
}

func decodeTimestamp(data [8]byte) time.Time {
	timestamp := binary.LittleEndian.Uint64(data[:])
	return time.Unix(0, int64(timestamp))
}

func ReadToBinaryMsg(file *bufio.Reader, chunkSize int) ([]byte, error) {
	buffer := make([]byte, chunkSize)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}
	if n == 0 {
		return nil, io.EOF // Datei komplett gelesen
	}

	// Chunk verarbeiten
	chunkData := buffer[:n]
	chunkHash := hashChunk(chunkData)

	chunk := FileChunk{
		Index: 0, // Initial, wird später gesetzt
		Data:  chunkData,
		Hash:  chunkHash,
	}

	// Senden des Chunks in BinaryMsg-Format
	sendFunc := sendChunkMessage(func(msg ByteNodeMsg) {
		fmt.Printf("Sende Chunk %d mit %d Bytes\n", binary.LittleEndian.Uint32(msg.ChunkIndex[:]), len(msg.ChunkData))
	})

	// Konvertierung & Verarbeitung mit bestehender Pipeline
	hashTracker := hashChunkMiddleware(&[]FileChunk{}, sendFunc)
	hashTracker(chunk)

	// In BinaryMsg encodieren & zurückgeben
	return encodeBinaryMsg(ByteNodeMsg{

		ChunkSize:  [4]byte{},
		ChunkIndex: [4]byte{},
		ChunkData:  chunkData,
		ChunkHash:  chunkHash,
	})
}
