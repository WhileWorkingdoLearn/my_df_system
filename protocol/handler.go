package transfer

type ResponseWriter interface{}

type MsgHandler struct {
	dns    map[string]interface{}
	Handle func(rw ResponseWriter, req *Request)
}

func (th MsgHandler) Transfer(rw ResponseWriter, req *Request) {

	th.Handle(rw, req)
}
