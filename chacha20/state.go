package chacha20

import (
	"bytes"
	"encoding/binary"
	"math/bits"
)

type State [16]uint32

func NewState(key, nonce []byte) State {
	keyData := Bytes2UintsKey(key)
	nonceData := Bytes2UintsNonce(nonce)
	var state State

	// NOTE: Constants are defined in RFC
	// https://datatracker.ietf.org/doc/html/rfc7539#page-4:~:text=The%20first%20four%20words%20(0%2D3)%20are%20constants%3A%200x61707865%2C%200x3320646e%2C%0A%20%20%20%20%20%200x79622d32%2C%200x6b206574.
	const (
		c0 = 0x61707865
		c1 = 0x3320646e
		c2 = 0x79622d32
		c3 = 0x6b206574
	)

	state[0] = c0
	state[1] = c1
	state[2] = c2
	state[3] = c3
	copy(state[4:], keyData[:])

	state[12] = 0 // counter
	copy(state[13:], nonceData[:])
	return state
}

func (s *State) SetCounter(v uint32) {
	s[12] = v
}

// https://datatracker.ietf.org/doc/html/rfc7539#section-2.1
func quarterRound(a, b, c, d uint32) (uint32, uint32, uint32, uint32) {
	a += b
	d ^= a
	d = bits.RotateLeft32(d, 16)
	c += d
	b ^= c
	b = bits.RotateLeft32(b, 12)
	a += b
	d ^= a
	d = bits.RotateLeft32(d, 8)
	c += d
	b ^= c
	b = bits.RotateLeft32(b, 7)
	return a, b, c, d
}

func (s *State) quarterRound(ia, ib, ic, id int) {
	s[ia], s[ib], s[ic], s[id] = quarterRound(s[ia], s[ib], s[ic], s[id])
}

// https://datatracker.ietf.org/doc/html/rfc7539#section-2.3.1
func (s *State) innerBlock() {
	s.quarterRound(0, 4, 8, 12)
	s.quarterRound(1, 5, 9, 13)
	s.quarterRound(2, 6, 10, 14)
	s.quarterRound(3, 7, 11, 15)
	s.quarterRound(0, 5, 10, 15)
	s.quarterRound(1, 6, 11, 12)
	s.quarterRound(2, 7, 8, 13)
	s.quarterRound(3, 4, 9, 14)
}

// KeyStream function calculates the state for a counter and returns the serialized data
// Notice we're not receiving a pointer here with `s State`, because we're working with
// the copy of the `State` to preserve the initial state as is.
//   - https://datatracker.ietf.org/doc/html/rfc7539#section-2.3.1:~:text=4%2C%209%2C14)%0A%20%20%20%20%20%20%20%20%20end-,chacha20_block,-(key%2C%20counter%2C%20nonce
func (s State) KeyStream(counter uint32) []byte {
	s.SetCounter(counter)
	var workingState = s // copy
	for i := 0; i < 10; i++ {
		s.innerBlock()
	}
	// add initial state to the current
	for i := range s {
		s[i] += workingState[i]
	}
	// serialize
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, s[:])
	return buf.Bytes()
}
