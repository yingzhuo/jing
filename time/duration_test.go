package time

import (
	"fmt"
	"testing"
)

func TestParseDurationWildly(t *testing.T) {

	d, err := ParseDurationWildly("2019-11-05 15:26:00 +0800")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}

}
