package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/riadafridishibly/chacha20poly1305"
	"github.com/riadafridishibly/chacha20poly1305/chacha20"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// key, err := hex.DecodeString(strings.ReplaceAll("00:01:02:03:04:05:06:07:08:09:0a:0b:0c:0d:0e:0f:10:11:12:13:14:15:16:17:18:19:1a:1b:1c:1d:1e:1f", ":", ""))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// nonce, err := hex.DecodeString(strings.ReplaceAll("00:00:00:09:00:00:00:4a:00:00:00:00", ":", ""))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// c := chacha20poly1305.NewChacha20(key, nonce)
	//
	// fmt.Println(hex.Dump(c.NextKeyStream()))
	// fmt.Println(hex.Dump(c.NextKeyStream()))
	// encryptionTest()
	stateExperiment()
}

const original = "22:4f:51:f3:40:1b:d9:e1:2f:de:27:6f:b8:63:1d:ed:8c:13:1f:82:3d:2c:06:e2:7e:4f:ca:ec:9e:f3:cf:78:8a:3b:0a:a3:72:60:0a:92:b5:79:74:cd:ed:2b:93:34:79:4c:ba:40:c6:3e:34:cd:ea:21:2c:4c:f0:7d:41:b7:69:a6:74:9f:3f:63:0f:41:22:ca:fe:28:ec:4d:c4:7e:26:d4:34:6d:70:b9:8c:73:f3:e9:c5:3a:c4:0c:59:45:39:8b:6e:da:1a:83:2c:89:c1:67:ea:cd:90:1d:7e:2b:f3:63"

func encryptionTest() {
	const key = "00:01:02:03:04:05:06:07:08:09:0a:0b:0c:0d:0e:0f:10:11:12:13:14:15:16:17:18:19:1a:1b:1c:1d:1e:1f"
	const nonce = "00:00:00:00:00:00:00:4a:00:00:00:00"
	keyBytes, err := hex.DecodeString(strings.ReplaceAll(key, ":", ""))
	if err != nil {
		log.Fatal(err)
	}
	nonceBytes, err := hex.DecodeString(strings.ReplaceAll(nonce, ":", ""))
	if err != nil {
		log.Fatal(err)
	}

	c := chacha20poly1305.NewChacha20(keyBytes, nonceBytes)
	var buf []byte
	buf = append(buf, c.NextKeyStream()...)
	buf = append(buf, c.NextKeyStream()...)
	buf = append(buf, c.NextKeyStream()...)
	buf = append(buf, c.NextKeyStream()...)
	buf = append(buf, c.NextKeyStream()...)

	var b strings.Builder
	b.Grow(len(buf) * 3)

	for _, v := range buf {
		b.WriteString(fmt.Sprintf("%02x:", v))
	}

	ourString := b.String()[:len(original)]
	fmt.Println("OUR:", ourString)
	fmt.Println("ORG:", original)
	fmt.Println(ourString == original)
}

func stateExperiment() {
	// State (mutate)
	s := chacha20.NewState(bytes.Repeat([]byte{1}, 32), bytes.Repeat([]byte{0}, 12))
	fmt.Printf("=> %x\n", s.KeyStream(1))
	fmt.Printf("=> %x\n", s.KeyStream(2))
	fmt.Printf("=> %x\n", s.KeyStream(3))

	// Chacha20
	c := chacha20poly1305.NewChacha20(bytes.Repeat([]byte{1}, 32), bytes.Repeat([]byte{0}, 12))
	fmt.Printf("chacha20 => %x\n", c.NextKeyStream())
	fmt.Printf("chacha20 => %x\n", c.NextKeyStream())
	fmt.Printf("chacha20 => %x\n", c.NextKeyStream())
}
