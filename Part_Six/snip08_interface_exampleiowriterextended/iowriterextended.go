// Wraps a writer to count written words
package wrapwritercountwords

import (
	"io"
)

type byteCounter struct {
	w       io.Writer
	written int64
}

func (c *byteCounter) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	c.written += int64(n)
	return
}

func CounterWriter(w io.Writer) (io.Writer, *int64) {
	c := &byteCounter{w, 0}
	return c, &c.written
}
