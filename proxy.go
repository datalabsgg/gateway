package gateway

type Proxy struct {
	test        string
	BeforeSpawn func(addr string)
	OnSpawn     func(string)
}

func New() *Proxy {
	a := "test"
	return &Proxy{test: a}
}
