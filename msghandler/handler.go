package handler

type ResponseWriter interface{}

type MsgHandler interface {
	ForwardMsg(rw ResponseWriter, req *Request)
}

type Handler struct {
	dns    map[string]interface{}
	Handle func(rw ResponseWriter, req *Request)
}

func (th Handler) ForwardMsg(rw ResponseWriter, req *Request) {

	th.Handle(rw, req)
}
