package transfer

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

type ByteNodeMsgHeader struct {
	Version     [1]byte
	MsgType     [1]byte
	Method      [1]byte
	Timestamp   [8]byte
	Timeout     [4]byte
	Domain      [32]byte
	Endpoint    [32]byte
	Auth        [16]byte
	PayloadType [1]byte
}

type ByteNodeMsg struct {
	ByteNodeMsgHeader
	ChunkSize  [4]byte
	ChunkIndex [4]byte
	ChunkData  []byte
	ChunkHash  []byte
}

const (
	OffsetByteVersion   = iota                     // 0
	OffsetByteMsgType                              // 1
	OffsetByteMethod                               // 2
	OffsetByteTimestamp = OffsetByteMethod + 1     // 8
	OffsetByteTimeout   = OffsetByteTimestamp + 4  // 12
	OffsetByteDomain    = OffsetByteTimeout + 4    // 16
	OffsetByteEndpoint  = OffsetByteTimestamp + 32 // 48
	OffsetByteAuth      = OffsetByteTimeout + 32   // 80
	OffsetByteDataType  = OffsetByteAuth + 16      // 96
	MsgHeaderEnd        = OffsetByteDataType + 1   //97

	OffsetByteChunkSize  = MsgHeaderEnd + 1 //
	OffsetByteChunkIndex = OffsetByteChunkSize + 4
)

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
	data_unkown = iota
	data_json
	data_text_utf8
	data_binary
	data_picture
	data_video
	data_audio
	data_custom
)
