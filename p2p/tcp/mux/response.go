package nmsgp

type MsgHeader struct {
	HasAuth     bool
	Auth        string
	PayloadType int
}

type ResponseWriter interface {
	Header() MsgHeader
	WritePayload([]byte) (int, error)
}

type Response struct {
	header MsgHeader
	data   []byte
}

func NewResponse() *Response {
	return &Response{
		header: MsgHeader{},
		data:   make([]byte, 0),
	}
}

func (r *Response) Header() MsgHeader                { return r.header }
func (r *Response) WritePayload([]byte) (int, error) { return 0, nil }
func (r *Response) Data() []byte                     { return r.data }
