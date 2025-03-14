package handler

import "log"

type HandleMethod struct {
	Path   string
	Method string
}

type DNS map[HandleMethod]MsgHandler

type SeverMux interface {
	Handle(path, method string, handle MsgHandler)
	ServeNMsgP(rw ResponseWriter, req *Request)
}

type Mux struct {
	domain string
	dns    DNS
}

func NewServerMux(domainName string) SeverMux {

	return &Mux{
		domain: domainName,
		dns:    make(DNS),
	}
}

func (mx *Mux) Handle(path, method string, handle MsgHandler) {
	hm := HandleMethod{Path: path, Method: method}
	if _, found := mx.dns[hm]; found {
		log.Fatal("handler for path already defined")
	}
	mx.dns[hm] = handle
}

func (mx *Mux) ServeNMsgP(rw ResponseWriter, req *Request) {
	hm := HandleMethod{Path: req.msgHeader.Endpoint, Method: req.Method()}
	if msgH, found := mx.dns[hm]; found {
		go msgH.ForwardMsg(rw, req)
	}
}
