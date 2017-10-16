package telegraf_json

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	return "telegraf_json"
}

// ConvertTelegrafJSON Convert telegram generated JSON
func (conv) Convert(d []byte) (common.Metric, error) {
	orgData := struct {
		Timestamp int64           `json:"timestamp"`
		Name      string          `json:"name"`
		Tags      json.RawMessage `json:"tags"`
		Fields    json.RawMessage `json:"fields"`
	}{}

	if err := json.Unmarshal(d, &orgData); err != nil {
		return nil, err
	}

	timestamp := time.Unix(orgData.Timestamp, 0)

	return &metric{
		ts:       timestamp,
		name:     orgData.Name,
		jsonData: fastMarshal(timestamp, orgData.Name, orgData.Tags, orgData.Fields),
	}, nil
}

func (conv) SetConfig(config map[string]interface{}) error {
	return nil
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// json.Marshal consumes too much heap space and it makes GC busy.
// fastMarshal does ALMOST SAME JOB with much more effecient way (but could be less safe).
func fastMarshal(ts time.Time, name string, tags json.RawMessage, fields json.RawMessage) string {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()

	b.WriteString(fmt.Sprintf(`{"@timestamp":"%s","name":"%s"`,
		ts.Format(time.RFC3339),
		name,
	))
	b.WriteString(`,"t":`)
	b.Write(tags)
	b.WriteString(fmt.Sprintf(`,"m":{"%s":`, name))
	b.Write(fields)
	b.WriteString(`}}`)

	result := b.String()
	bufPool.Put(b)

	return result
}
