package object

import (
	"bytes"
	"strconv"
)

type Blob struct {
	Data []byte
}

func (b Blob) Body() string {

	var buf bytes.Buffer
	buf.WriteString(b.Type())
	buf.WriteString(" ")
	buf.WriteString(strconv.Itoa(len(b.Data)))
	buf.WriteByte(0)
	buf.Write(b.Data)

	return buf.String()
}

func (b Blob) Type() string {
	return "blob"
}
