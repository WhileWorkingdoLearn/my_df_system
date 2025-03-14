package msg

import "time"

/*
Protocol structure
  	- Version     	[1]byte  	Protocol-Version (0-255)
  	- MsgType     	[1]byte  	Messagetype (Request, Response, Data-Chunk...)
	- Method   		[1]byte
  	- Timestamp   	[8]byte	Unix-Timestamp (uint64, Nanoseconds-Precision)
  	- Timeout     	[4]byte   relative timeout (uint32, 4 Bytes, milliseconds since receiving)
	- Domain      	[32]byte
	- Endpoint   	[32]byte
	- HasAuth     	1 bit	// FLag for has Auth. If set to false payload type and size get ignored
  	- Auth        	[16]byte // Authentification-Token or Signature (HMAC, JWT-Hash, etc.)
	- HasPayload	1 bit	// FLag for has payload. If set to false payload type and size get ignored
  	- PayloadType 	[1]byte  // Typ e ofPayload (1 = JSON, 2 = Text, 3 = Binary, etc.)
	- PayloadSize 	[8]
    }
*/

type MsgHeader struct {
	Version     int
	MsgType     int
	Method      int
	Timestamp   int
	Timeout     time.Duration
	Domain      string
	Endpoint    string
	HasAuth     bool
	Auth        string
	HasPayload  bool
	PayloadType int
	PayloadSize int
}

type Msg struct {
	Header     MsgHeader
	ChunkSize  int
	ChunkIndex int
	ChunkData  []byte
	ChunkHash  []byte
}
