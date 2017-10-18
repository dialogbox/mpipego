package common

import (
	"bytes"
	"fmt"
	"time"
)

// Converter interface
type Converter interface {
	Convert([]byte) (Metric, error)
	Name() string
}

type Metric interface {
	Timestamp() time.Time
	Name() string
	JSON() string
}

// FastMarshal does ALMOST SAME JOB as json.Marshal using much more effecient way (but could be less safe).
// json.Marshal consumes too much heap space and it makes GC busy.
func FastMarshal(buf *bytes.Buffer, ts time.Time, name string, tags []byte, fields []byte) {
	buf.Reset()

	buf.WriteString(fmt.Sprintf(`{"@timestamp":"%s","name":"%s"`,
		ts.Format(time.RFC3339),
		name,
	))
	buf.WriteString(`,"t":`)
	buf.Write(tags)
	buf.WriteString(fmt.Sprintf(`,"m":{"%s":`, name))
	buf.Write(fields)
	buf.WriteString(`}}`)
}

// ByteBufferWriteAll write all parameters to buf
func ByteBufferWriteAll(buf *bytes.Buffer, data ...interface{}) {
	for i := range data {
		switch d := data[i].(type) {
		case rune:
			buf.WriteByte(byte(d))
		case byte:
			buf.WriteByte(d)
		case []byte:
			buf.Write(d)
		case string:
			buf.WriteString(d)
		default:
			fmt.Printf("unexpected type %T : %v\n", d, d)
		}
	}
}
