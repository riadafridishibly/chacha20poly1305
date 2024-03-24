package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var (
		num         uint32 = 0x1a2b3c4d
		big, little [4]byte
	)

	binary.BigEndian.PutUint32(big[:], num)
	binary.LittleEndian.PutUint32(little[:], num)

	fmt.Printf("Human Readable: 0x%x\n", num)
	fmt.Printf("Big     Endian: 0x%x\n", big)
	fmt.Printf("Little  Endian: 0x%x\n", little)
}
