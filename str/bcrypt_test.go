package str

import (
	"fmt"
	"testing"
)

func TestCheckBCrypt(t *testing.T) {

	raw := "hello"
	hashed := BCrypt(raw)

	fmt.Println(raw)
	fmt.Println(hashed)

	fmt.Println(CheckBCrypt(raw, hashed))
}
