package nmsgp

/*
type HandleMethod struct {
	Path   string
	Method string
}

type DNS map[HandleMethod]MsgHandler

type SeverMux interface {
	Handle(path, method string, handle MsgHandler)
	HandleFunc(path, method string, handle func(res ResponseWriter, req *Request))
	Handler(r *Request) (h MsgHandler, pattern string)
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

func (mx *Mux) HandleFunc(path, method string, handle func(res ResponseWriter, req *Request)) {
}

func (mx *Mux) Handler(req *Request) (h MsgHandler, pattern string) {
	hm := HandleMethod{Path: req.msgHeader.Endpoint, Method: req.Method()}

	if msgH, found := mx.dns[hm]; found {
		return msgH, hm.Path
	}

	return nil, req.msgHeader.Endpoint
}
*/
