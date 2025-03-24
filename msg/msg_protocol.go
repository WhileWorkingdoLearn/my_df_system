package msg

import (
	"hash/crc32"
	"time"
)

/*
Protocol structure
  	- Version     	[1]byte  	Protocol-Version (0-255)
  	- MsgType     	[1]byte  	Messagetype (Request, Response, Data-Chunk...)
	- Method   		[1]byte
  	- Timestamp   	[8]bytes	Unix-Timestamp (uint64, Nanoseconds-Precision)
  	- Timeout     	[4]bytes   relative timeout (uint32, 4 Bytes, milliseconds since receiving)
	- Domain      	[32]bytes
	- Endpoint   	[32]bytes
	- HasAuth     	[1]byte  // FLag for has Auth. If set to false payload type and size get ignored
  	- Auth        	[16]byte // Authentification-Token or Signature (HMAC, JWT-Hash, etc.)
	- HasPayload	[1]byte  	// FLag for has payload. If set to false payload type and size get ignored
  	- PayloadType 	[1]byte  // Typ e ofPayload (1 = JSON, 2 = Text, 3 = Binary, etc.)
	- PayloadSize 	[8]bytes
	-Checksum       [4]bytes
    }
*/

type MsgHeader struct {
	Version     int
	MsgType     int
	Error       int
	Method      int
	Timestamp   time.Time
	Timeout     time.Duration
	Domain      string
	Endpoint    string
	HasAuth     bool
	Auth        string
	HasPayload  bool
	PayloadType int
	PayloadSize int
	Checksum    int
}

/*
	type Msg struct {
		Header     MsgHeader
		ChunkSize  int
		ChunkIndex int
		ChunkData  []byte
		ChunkHash  []byte
	}
*/
func computeCRC32Checksum(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

func ValidateCRC32Checksum(data []byte, expectedChecksum uint32) bool {
	return computeCRC32Checksum(data) == expectedChecksum
}
