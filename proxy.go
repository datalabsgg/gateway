package gateway

type Proxy struct {
	test       string
	BeforeJoin func(addr string)
	OnJoin     func(string)
}

func New() *Proxy {
	a := "test"
	return &Proxy{test: a}
}
