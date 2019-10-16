package jstr

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func NewUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func NewUUID36() string {
	return NewUUID()
}

func NewUUID32() string {
	return strings.ReplaceAll(NewUUID(), "-", EmptyString)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
