package common

import (
	"bytes"
	"reflect"

	"fmt"
	"time"

	"github.com/json-iterator/go"
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
	nameJSON, err := jsoniter.Marshal(m.Name)
	if err != nil {
		return nil, err
	}
	tagsJSON, err := jsoniter.Marshal(m.Tags)
	if err != nil {
		return nil, err
	}
	fieldsJSON, err := jsoniter.Marshal(m.Fields)
	if err != nil {
		return nil, err
	}
	timeStr := m.Timestamp.Format(time.RFC3339)

	j := fmt.Sprintf(`{"@timestamp":"%s","name":%s,"t":%s,"m":{%s:%s}}`,
		timeStr, nameJSON, tagsJSON, nameJSON, fieldsJSON)

	return []byte(j), nil
}

// Identical deeply compares two metric
// It used by test.
func (m *Metric) Identical(m2 *Metric) bool {
	if m2 == nil {
		return false
	}

	if m.Name != m2.Name || m.Timestamp != m2.Timestamp {
		return false
	}

	if !reflect.DeepEqual(m.Tags, m2.Tags) {
		return false
	}

	if len(m.Fields) != len(m2.Fields) {
		return false
	}

	for k, v := range m.Fields {
		if fmt.Sprint(v) != fmt.Sprint(m2.Fields[k]) {
			return false
		}
	}

	return true
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
