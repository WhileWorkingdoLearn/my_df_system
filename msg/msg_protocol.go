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
  	- Auth        	[16]byte // Authentification-Token or Signature (HMAC, JWT-Hash, etc.)
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

// ByteNodeMsg Struktur
type Msg struct {
	Header     MsgHeader
	ChunkSize  int
	ChunkIndex int
	ChunkData  []byte
	ChunkHash  []byte
}

/*
Value	Typ	Examples
0	Unbekannt	-
1	JSON	application/json
2	Text (UTF-8)	text/plain
3	Binary	application/octet-stream
4	Picture	image/png
5	Video	video/mp4
6	Audio	audio/mpeg
7	Custom (defined through sender)


const (
	data_unkown = iota
	data_json
	data_text_utf8
	data_binary
	data_picture
	data_video
	data_audio
	data_custom
)
*/
