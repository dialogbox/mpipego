package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Converter interface
type Converter interface {
	Convert([]byte) (*Metric, error)
	Name() string
}

type Metric struct {
	Name      string
	Timestamp time.Time
	Tags      map[string]string
	Fields    map[string]interface{}
}

func (m *Metric) MarshalJSON() ([]byte, error) {
	nameJSON, err := json.Marshal(m.Name)
	if err != nil {
		return nil, err
	}
	tagsJSON, err := json.Marshal(m.Tags)
	if err != nil {
		return nil, err
	}
	fieldsJSON, err := json.Marshal(m.Fields)
	if err != nil {
		return nil, err
	}
	timeStr := m.Timestamp.Format(time.RFC3339)

	j := fmt.Sprintf(`{"@timestamp":"%s","name":%s,"t":%s,"m":{%s:%s}}`,
		timeStr, nameJSON, tagsJSON, nameJSON, fieldsJSON)

	return []byte(j), nil
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
