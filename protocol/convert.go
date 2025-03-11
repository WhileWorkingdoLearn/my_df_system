package transfer

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

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
	sendFunc := sendChunkMessage(func(msg NodeMsg) {
		fmt.Printf("Sende Chunk %d mit %d Bytes\n", binary.LittleEndian.Uint32(msg.ChunkIndex[:]), len(msg.ChunkData))
	})

	// Konvertierung & Verarbeitung mit bestehender Pipeline
	hashTracker := hashChunkMiddleware(&[]FileChunk{}, sendFunc)
	hashTracker(chunk)

	// In BinaryMsg encodieren & zurückgeben
	return encodeBinaryMsg(NodeMsg{
		Version:    [1]byte{1},
		MsgType:    [1]byte{1},
		ChunkSize:  [4]byte{},
		ChunkIndex: [4]byte{},
		ChunkData:  chunkData,
		ChunkHash:  chunkHash,
	})
}

func MessageTypeToString(msgType byte) string {
	switch msgType {
	case MsgTypePing:
		return "PING"
	case MsgTypePong:
		return "PONG"
	case MsgTypeData:
		return "DATA"
	case MsgTypeAck:
		return "ACK"
	case MsgTypeError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
