package handler



type MsgHandler interface {
	ForwardMsg(rw ResponseWriter, req *Request)
}

type Handler struct {
	Handle func(rw ResponseWriter, req *Request)
}

func (th Handler) ForwardMsg(rw ResponseWriter, req *Request) {

	th.Handle(rw, req)
}


