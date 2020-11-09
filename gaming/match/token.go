package match

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

type Token [20]byte

func (tok Token) String() string {
	return base64.StdEncoding.EncodeToString(tok[:])
}

func simpleToken() (tok Token) {
	b := make([]byte, 20)
	_, err := rand.Read(b)
	if err != nil {
		log.Print(err)
		return
	}
	copy(tok[:], b)
	return
}
