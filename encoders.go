package idGen

import (
	"crypto/sha512"
	"encoding/base64"
	"hash"
	"sync"
)

//===========[INTERFACES]====================================================================================================

//Encoder defines interface for sha512encoder
type Encoder interface {
	Encode(string) string
}

//===========[STRUCTS]====================================================================================================

//sha512encoder defines safe type for encoding the string
type sha512encoder struct {
	c  hash.Hash
	mx sync.Mutex
}

//Encode encodes given string using sha512 algorithm
func (e *sha512encoder) Encode(s string) string {
	e.mx.Lock()
	e.c.Write([]byte(s))
	defer e.mx.Unlock()
	defer e.c.Reset()
	return base64.URLEncoding.EncodeToString(e.c.Sum(nil))
}

//===========[FUNCTIONALITY]============================================================================================

//newSha512Encoder returns newly initiated sha512encoder
func newSha512Encoder() Encoder {
	return &sha512encoder{
		c:  sha512.New(),
		mx: sync.Mutex{},
	}
}
