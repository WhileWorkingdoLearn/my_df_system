package transfer

type TransferHandler struct {
	Handle func(rw ResponseWriter, req *Request)
}

func (th TransferHandler) Transfer(rw ResponseWriter, req *Request) {
	
	th.Handle(rw, req)
}
