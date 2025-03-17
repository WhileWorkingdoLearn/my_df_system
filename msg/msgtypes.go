package msg

const (
	IdxPING = iota
	IdxPONG
	IdxHEADER
	IdxDATA
	IdxERROR
	IdxEND
	IdxCUSTOM
)

type MsgType string

const (
	PING    MsgType = "PING"
	PONG    MsgType = "PONG"
	HEADER  MsgType = "HEADER"
	DATA    MsgType = "DATA"
	ERROR   MsgType = "ERROR"
	END     MsgType = "END"
	CUSTOM  MsgType = "CUSTOM"
	UNKNOWN MsgType = "UNKNOWN"
)

func (m *MsgType) Ping() string    { return "PING" }
func (m *MsgType) Pong() string    { return "PONG" }
func (m *MsgType) Header() string  { return "HEADER" }
func (m *MsgType) Data() string    { return "DATA" }
func (m *MsgType) Error() string   { return "ERROR" }
func (m *MsgType) End() string     { return "END" }
func (m *MsgType) Custom() string  { return "CUSTOM" }
func (m *MsgType) Unkwown() string { return "UNKNOWN" }
func (m MsgType) String() string   { return string(m) }

func ConvertMsgType(msgtype int) MsgType {

	switch msgtype {
	case IdxPING:
		return PING
	case IdxPONG:
		return PONG
	case IdxHEADER:
		return HEADER
	case IdxDATA:
		return DATA
	case IdxERROR:
		return ERROR
	case IdxEND:
		return END
	case IdxCUSTOM:
		return CUSTOM
	default:
		return UNKNOWN
	}
}

const (
	IdxFETCH = iota
	IdxCREATE
	IdxUPDATE
	IdxCOPY
	idxDELETE
)

type MethodType string

const (
	FETCH   MethodType = "FETCH"
	CREATE  MethodType = "CREATE"
	UPADATE MethodType = "UPDATE"
	COPY    MethodType = "COPY"
	DELETE  MethodType = "DELETE"
)

func (mt *MethodType) Fetch() string  { return "FETCH" }
func (mt *MethodType) Create() string { return "CREATE" }
func (mt *MethodType) Update() string { return "UPDATE" }
func (mt *MethodType) Copy() string   { return "COPY" }
func (mt *MethodType) Delete() string { return "DELETE" }
func (mt MethodType) String() string  { return string(mt) }

func ConvertMethod(methodtype int) MethodType {
	switch methodtype {
	case IdxFETCH:
		return FETCH
	case IdxCREATE:
		return CREATE
	case IdxUPDATE:
		return UPADATE
	case IdxCOPY:
		return COPY
	case idxDELETE:
		return DELETE
	default:
		return FETCH
	}
}

const (
	IdxUnkown = iota
	IdxJSON
	IdxTEXT
	IdxBINARY
	IdxPICTURE
	IdxVIDEO
	IdxAUDIO
	IdxCUSTOM_MEDIA
)

type MediaType string

const (
	JSON         MediaType = "application/json"
	TEXT         MediaType = "application/txt"
	BINARY       MediaType = "application/binary"
	PICTURE      MediaType = "media/imgage"
	VIDEO        MediaType = "media/video"
	AUDIO        MediaType = "media/audio"
	CUSTOM_MEDIA MediaType = "custom"
	UNKOWN_MEDIA MediaType = "UNKNOWN"
)

func (mt *MediaType) Json() string    { return string(JSON) }
func (mt *MediaType) Text() string    { return string(TEXT) }
func (mt *MediaType) Binary() string  { return string(BINARY) }
func (mt *MediaType) Picture() string { return string(PICTURE) }
func (mt *MediaType) Video() string   { return string(VIDEO) }
func (mt *MediaType) Audio() string   { return string(AUDIO) }
func (mt *MediaType) Custom() string  { return string(CUSTOM_MEDIA) }
func (mt MediaType) String() string   { return string(mt) }

func ConvertMediaType(mediatype int) MediaType {
	switch mediatype {
	case IdxJSON:
		return JSON
	case IdxTEXT:
		return TEXT
	case IdxBINARY:
		return BINARY
	case IdxPICTURE:
		return PICTURE
	case IdxVIDEO:
		return VIDEO
	case IdxAUDIO:
		return AUDIO
	case IdxCUSTOM_MEDIA:
		return CUSTOM_MEDIA
	default:
		return UNKOWN_MEDIA
	}
}

const (
	None = iota * 100
	ErrInfo
	ErrRedirection
	ErrClient
	ErrServer
)

const (
	ErrInfo2 = iota + ErrInfo
)

const (
	ErrRedirection1 = iota + ErrRedirection
)

const (
	ErrVersion = iota + ErrClient
	ErrMsgType
)

const (
	ErrDomain = iota + ErrServer
	ErrEndpoint
)
