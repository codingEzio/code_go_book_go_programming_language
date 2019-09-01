// Wraps a writer to count written words
package wrapwritercountwords

import (
	"bytes"
	"testing"
)

func TestCounterWriter(t *testing.T) {
	b := &bytes.Buffer{}
	c, n := CounterWriter(b)

	data := []byte("Hi there")
	c.Write(data)

	if *n != int64(len(data)) {
		t.Logf("%d != %d", n, len(data))
		t.Fail()
	}
}
