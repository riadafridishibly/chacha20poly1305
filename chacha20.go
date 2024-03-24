package chacha20poly1305

import (
	"github.com/riadafridishibly/chacha20poly1305/chacha20"
)

type Chacha20 struct {
	initialState chacha20.State
	counter      uint32
}

func NewChacha20(key, nonce []byte) *Chacha20 {
	return &Chacha20{
		initialState: chacha20.NewState(key, nonce),
	}
}

func (c *Chacha20) NextKeyStream() []byte {
	c.counter++
	return c.initialState.KeyStream(c.counter)
}
