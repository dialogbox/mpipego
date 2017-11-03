package protobuf

import (
	"bytes"
	"encoding/json"
	fmt "fmt"
	"sync"
	"time"

	"github.com/dialogbox/mpipego/common"
	proto "github.com/golang/protobuf/proto"
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
	return "protobuf"
}

func (conv) Convert(d []byte) (common.Metric, error) {
	m := &Metric{}
	err := proto.Unmarshal(d, m)
	if err != nil {
		return nil, err
	}

	timestamp := time.Unix(0, m.GetTimestamp())

	name := m.GetName()

	tags, err := json.Marshal(m.GetTags())
	if err != nil {
		return nil, err
	}

	fields, err := encodeFields(m)
	if err != nil {
		return nil, err
	}

	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	common.FastMarshal(buf, timestamp, name, tags, fields)

	return &metric{
		ts:       timestamp,
		name:     name,
		jsonData: buf.String(),
	}, nil
}

func (conv) SetConfig(config map[string]interface{}) error {
	return nil
}

func encodeFields(m *Metric) (json.RawMessage, error) {
	f := make(map[string]interface{})

	for k, v := range m.GetFields() {
		switch v := v.GetValue().(type) {
		case *FieldValue_IntValue:
			f[k] = v.IntValue
		case *FieldValue_FloatValue:
			f[k] = v.FloatValue
		case *FieldValue_StringValue:
			f[k] = v.StringValue
		case *FieldValue_BoolValue:
			f[k] = v.BoolValue
		default:
			return nil, fmt.Errorf("Unsupported field data type: [%s, %v(%T)]", k, v, v)
		}
	}
	return json.Marshal(f)
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
