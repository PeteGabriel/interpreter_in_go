package repl

import (
	"bytes"
	is2 "github.com/matryer/is"
	"testing"
)

func TestStart(t *testing.T) {
	is := is2.New(t)
	in := bytes.NewReader([]byte("10 + 10"))
	var out bytes.Buffer
	Start(in, &out)

	vals := out.String()
	exp := "{Type:INT Literal:10}\n{Type:+ Literal:+}\n{Type:INT Literal:10}\n"

	is.True(vals == exp)
}