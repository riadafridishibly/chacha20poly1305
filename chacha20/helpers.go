package chacha20

import (
	"bytes"
	"encoding/binary"
)

func Bytes2UintsKey(key []byte) *[8]uint32 {
	if len(key) != 32 {
		panic("key must be 32 bytes")
	}
	var data [8]uint32
	err := binary.Read(bytes.NewReader(key), binary.LittleEndian, &data)
	if err != nil {
		panic(err)
	}
	return &data
}

func Bytes2UintsNonce(nonce []byte) *[3]uint32 {
	if len(nonce) != 12 {
		panic("nonce must be 12 bytes")
	}
	var data [3]uint32
	err := binary.Read(bytes.NewReader(nonce), binary.LittleEndian, &data)
	if err != nil {
		panic(err)
	}
	return &data
}
