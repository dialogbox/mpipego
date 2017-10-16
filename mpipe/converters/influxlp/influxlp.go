package influxlp

// https://docs.influxdata.com/influxdb/v1.3/write_protocols/line_protocol_tutorial/

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/dialogbox/mpipego/common"
)

type metric struct {
	ts       time.Time
	name     string
	jsonData string
}

func (m *metric) Timestamp() time.Time {
	return m.ts
}

func (m *metric) Name() string {
	return m.name
}

func (m *metric) JSON() string {
	return m.jsonData
}

func NewConverter() common.Converter {
	return &conv{}
}

type conv struct{}

func (conv) Name() string {
	return "influxlp"
}

func (conv) Convert(d []byte) (common.Metric, error) {
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()

	tags := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(tags)
	tags.Reset()

	fields := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(fields)
	fields.Reset()

	cur := d

	// Read measure
	cur = readNextToken(buf, cur)
	name := buf.String()

	isFirst := true
	tags.WriteByte('{')
	for cur[0] != ' ' {
		if isFirst {
			isFirst = false
		} else {
			tags.WriteByte(',')
		}
		cur = readNextToken(buf, cur[1:])
		byteBufferWriteAll(tags, '"', buf.Bytes(), `":"`)
		cur = readNextToken(buf, cur[1:])
		byteBufferWriteAll(tags, buf.Bytes(), '"')

	}
	tags.WriteByte('}')

	// Read fields
	isFirst = true
	fields.WriteByte('{')
	for cur[0] != ' ' || isFirst {
		if isFirst {
			isFirst = false
		} else {
			fields.WriteByte(',')
		}

		// Read field name
		cur = readNextToken(buf, cur[1:])
		byteBufferWriteAll(fields, '"', buf.Bytes(), `":`)

		// Read field value
		cur = readNextFieldValue(buf, cur[1:])
		fieldValue, err := parseFieldValue(buf.Bytes())
		if err != nil {
			return nil, fmt.Errorf("Format error: Can not parse field value %v", err)
		}

		switch fieldValue := fieldValue.(type) {
		case string:
			fields.WriteString(fmt.Sprintf(`"%v"`, fieldValue))
		default:
			fields.WriteString(fmt.Sprintf(`%v`, fieldValue))
		}
	}
	fields.WriteString("}")

	// Read timestamp
	readNextToken(buf, cur[1:])
	ts, err := strconv.ParseInt(buf.String(), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Format error: Invalid timestamp %s", buf.String())
	}
	timestamp := time.Unix(0, ts)

	common.FastMarshal(buf, timestamp, name, tags.Bytes(), fields.Bytes())

	return &metric{
		ts:       timestamp,
		name:     name,
		jsonData: buf.String(),
	}, nil
}

func (conv) SetConfig(config map[string]interface{}) error {
	return nil
}

func readNextToken(buffer *bytes.Buffer, lp []byte) []byte {
	buffer.Reset()
	i := 0
	for ; i < len(lp); i++ {
		switch lp[i] {
		case ',':
			return lp[i:]
		case ' ':
			return lp[i:]
		case '=':
			return lp[i:]
		case '\\':
			i++
			buffer.WriteByte(lp[i])
		default:
			buffer.WriteByte(lp[i])
		}
	}

	return lp[i:]
}

func readNextFieldValue(buffer *bytes.Buffer, lp []byte) []byte {
	buffer.Reset()
	var i int

	if lp[0] != '"' { // Numeric value
		for i = 0; lp[i] != ',' && lp[i] != ' '; i++ {
			buffer.WriteByte(lp[i])
		}
	} else { // String value
		buffer.WriteByte('"')
		for i = 1; lp[i] != '"'; i++ {
			switch lp[i] {
			case '\\':
				i++
				buffer.WriteByte(lp[i])
			default:
				buffer.WriteByte(lp[i])
			}
		}
		buffer.WriteByte('"')
		i++
	}

	return lp[i:]
}

func parseFieldValue(v []byte) (interface{}, error) {
	if v[0] == '"' {
		return string(v[1 : len(v)-1]), nil
	}

	if v[len(v)-1] == 'i' {
		intv, err := strconv.ParseInt(string(v[:len(v)-1]), 10, 64)
		if err != nil {
			return nil, err
		}

		return intv, nil
	}

	switch v[0] {
	case 't':
		return true, nil
	case 'T':
		return true, nil
	case 'f':
		return false, nil
	case 'F':
		return false, nil
	}

	return strconv.ParseFloat(string(v), 64)
}

func byteBufferWriteAll(buf *bytes.Buffer, data ...interface{}) {
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

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
