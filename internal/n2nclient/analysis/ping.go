package analysis

type Pinger struct {
	OnSend func()
	OnRecv func()
	OnIdle func()
}

func NewPinger() *Pinger {
	p := &Pinger{}
	p.OnSend = func() {}
	p.OnRecv = func() {}
	p.OnIdle = func() {}
	return p
}

func (p *Pinger) Ping() error {
	return nil
}
