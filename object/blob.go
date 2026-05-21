package object

import (
	"bytes"
	"strconv"
)

type Blob struct {
	Data []byte
}

func (b Blob) Body() []byte {

	var buf bytes.Buffer
	buf.WriteString(b.Type())
	buf.WriteString(" ")
	buf.WriteString(strconv.Itoa(len(b.Data)))
	buf.WriteByte(0)
	buf.Write(b.Data)

	return buf.Bytes()
}

func (b Blob) Type() string {
	return "blob"
}
