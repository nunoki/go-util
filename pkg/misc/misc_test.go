package misc

import (
	"strconv"
	"testing"
)

func TestMust(t *testing.T) {
	a := "666"
	b := Must(strconv.Atoi(a))

	if b != int(666) {
		t.Logf("expected %d, received %d", 666, b)
	}

	a = "abc"
	b = Must(strconv.Atoi(a))

	if b != int(0) {
		t.Logf("expected %d, received %d", 0, b)
	}
}
