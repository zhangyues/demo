package main

type Pub struct {
	pc   chan handler
	name string
	h    handler
}
type handler func(name string, req []byte)

func NewRegisterSubscribe(name string, h handler) {
	pub := new(Pub)
	pc := make(chan handler)
	pc <- h
	pub.pc = pc
}
