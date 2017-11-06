package msgpack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

//go:generate msgp

// Metric is structure to define MessagePack message format
// will be used by msgp code generator
type Metric struct {
	N      string                 `msg:"name"`
	Time   time.Time              `msg:"time"`
	Tags   map[string]string      `msg:"tags"`
	Fields map[string]interface{} `msg:"fields"`
}

func (m *Metric) Timestamp() time.Time {
	return m.Time
}

func (m *Metric) Name() string {
	return m.N
}

func (m *Metric) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}

	nameb, err := json.Marshal(m.N)
	if err != nil {
		return nil, err
	}
	buf.WriteString(fmt.Sprintf(`{"name":%s,"@timestamp":"%s"`, nameb, m.Time.Format(time.RFC3339)))

	b, err := json.Marshal(m.Tags)
	buf.WriteString(fmt.Sprintf(`,t:%s`, b))

	b, err = json.Marshal(m.Fields)
	buf.WriteString(fmt.Sprintf(`,m:{%s:%s}}`, nameb, b))

	return buf.Bytes(), nil
}

func (m *Metric) JSON() string {
	b, err := m.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(b)
}
