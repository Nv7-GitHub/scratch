package types

import (
	"fmt"
	"math/rand"
	"time"
)

type empty struct{}

const length = 32

var used = make(map[string]empty)

func GetRandomString() string {
	exists := true
	var txt string
	for exists {
		rand.Seed(time.Now().UnixNano())
		b := make([]byte, length)
		rand.Read(b)
		txt = fmt.Sprintf("%x", b)[:length]
		_, exists = used[txt]
	}
	used[txt] = empty{}
	return txt
}
