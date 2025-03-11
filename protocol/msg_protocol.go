package transfer

type NodeMsg struct {
	Version     [1]byte
	MsgType     [1]byte
	Timestamp   [8]byte
	Timeout     [4]byte
	Auth        [16]byte
	PayloadType [1]byte
	ChunkSize   [4]byte
	ChunkIndex  [4]byte
	ChunkData   []byte
	ChunkHash   []byte
}

const (
	MsgTypePing  byte = 0x01
	MsgTypePong  byte = 0x02
	MsgTypeData  byte = 0x03
	MsgTypeAck   byte = 0x04
	MsgTypeError byte = 0x05
)

const (
	OffsetVersion     = iota                  // 0
	OffsetMsgType                             // 1
	OffsetTimestamp                           // 2  (Start)
	OffsetTimeout     = OffsetTimestamp + 8   // 10
	OffsetAuth        = OffsetTimeout + 4     // 14
	OffsetPayloadType = OffsetAuth + 16       // 30
	OffsetChunkSize   = OffsetPayloadType + 1 // 31
	OffsetChunkIndex  = OffsetChunkSize + 4   // 35
	OffsetChunkData   = OffsetChunkIndex + 4  // 75
)

/*
Protocol structure
  - Version     [1]byte  Protocol-Version (0-255)
  - MsgType     [1]byte  // Messagetype (Request, Response, Data-Chunk...)
  - Timestamp   [8]byte  // Unix-Timestamp (uint64, Nanoseconds-Precision)
  - Timeout     [4]byte  // relative timeout (uint32, 4 Bytes, milliseconds since receiving)
  - Auth        [16]byte // Authentification-Token or Signature (HMAC, JWT-Hash, etc.)
  - PayloadType [1]byte  // Typ e ofPayload (1 = JSON, 2 = Text, 3 = Binary, etc.)
  - ChunkSize   [4]byte  // Byte size of Data Chunk
  - NumChunks   [4]byte  //  Number of Chunks of File
    RootHash    [32]byte // Root Hash of Merkle Root for Data validation
    ChunkIndex  [4]byte  // Index of aktual Chunks
    ChunkData   []byte   // Data / Payload
    ChunkHash   [32]byte // CheckSum of Chunks
    Checksum    [4]byte  //  Checksum of Total message
    }
*/

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
*/

const (
	Unkown = iota
	JSON
	TEXT_utf8
	Binary
	Picture
	Video
	Audio
	Custom
)
