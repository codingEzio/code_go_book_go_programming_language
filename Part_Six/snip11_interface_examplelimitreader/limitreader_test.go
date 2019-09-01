package limitreader

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	str := "heyguys"
	byt := &bytes.Buffer{}
	reader := LimitReader(strings.NewReader(str), 4)
	n, _ := byt.ReadFrom(reader)

	if n != 4 {
		t.Logf("n=%d", n)
		t.Fail()
	}
	if byt.String() != "heyg" {
		t.Logf(`"%s" != "%s"`, byt.String(), str)
	}
}
