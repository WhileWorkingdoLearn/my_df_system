package transfer

import (
	"crypto/sha256"
	"encoding/binary"
)

// FileChunk zur Hash-Verarbeitung
type FileChunk struct {
	Index uint32
	Data  []byte
	Hash  []byte
}

func hashChunk(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func hashChunkMiddleware(chunks *[]FileChunk, process func(FileChunk)) func(FileChunk) {
	return func(chunk FileChunk) {
		*chunks = append(*chunks, chunk)
		process(chunk)
	}
}

// **Konvertiert FileChunk in BinaryMsg & sendet es**
func sendChunkMessage(send func(NodeMsg)) func(FileChunk) {
	return func(chunk FileChunk) {
		var msg NodeMsg
		msg.Version[0] = 1
		msg.MsgType[0] = 1 // Chunk-Transfer
		binary.LittleEndian.PutUint32(msg.ChunkSize[:], uint32(len(chunk.Data)))
		binary.LittleEndian.PutUint32(msg.ChunkIndex[:], chunk.Index)
		copy(msg.ChunkHash[:], chunk.Hash)
		msg.ChunkData = chunk.Data

		// Nachricht senden
		send(msg)
	}
}

// **Encodiert BinaryMsg in ein Binärformat**
func encodeBinaryMsg(msg NodeMsg) ([]byte, error) {
	totalSize := 10
	/*MsgByteSizes.Version +
	MsgByteSizes.MsgType +
	MsgByteSizes.Timestamp +
	MsgByteSizes.Timeout +
	MsgByteSizes.Auth +
	MsgByteSizes.PayloadType +
	+len(msg.ChunkData) + 32*/

	buffer := make([]byte, totalSize)

	offset := 0
	buffer[offset] = msg.Version[0]
	offset++
	buffer[offset] = msg.MsgType[0]
	offset++
	binary.LittleEndian.PutUint32(buffer[offset:], binary.LittleEndian.Uint32(msg.ChunkSize[:]))
	offset += 4
	binary.LittleEndian.PutUint32(buffer[offset:], binary.LittleEndian.Uint32(msg.ChunkIndex[:]))
	offset += 4
	copy(buffer[offset:], msg.ChunkData)
	offset += len(msg.ChunkData)
	copy(buffer[offset:], msg.ChunkHash[:])

	return buffer, nil
}
