package telegraf_json

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dialogbox/mpipego/common"
)

type metric struct {
	ts         time.Time
	name       string
	TimeString string                      `json:"@timestamp"`
	Tags       *json.RawMessage            `json:"t"`
	Fields     map[string]*json.RawMessage `json:"m"`
}

func (m *metric) Timestamp() time.Time {
	return m.ts
}

func (m *metric) Data() interface{} {
	return m
}

func (m *metric) JSON() []byte {
	jsondata, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return jsondata
}

func (m *metric) String() string {
	return fmt.Sprintf("[%v] name: %s, tags: %v, fields: %v", m.ts, m.name, m.Tags, m.Fields)
}

func NewConverter() *conv {
	return &conv{}
}

type conv struct{}

func (conv) Name() string {
	return "telegraf_json"
}

// ConvertTelegrafJSON Convert telegram generated JSON
func (conv) Convert(d []byte) (common.Metric, error) {
	orgData := struct {
		Timestamp int64            `json:"timestamp"`
		Name      string           `json:"name"`
		Tags      *json.RawMessage `json:"tags"`
		Fields    *json.RawMessage `json:"fields"`
	}{}

	if err := json.Unmarshal(d, &orgData); err != nil {
		return nil, err
	}

	timestamp := time.Unix(orgData.Timestamp, 0)
	result := &metric{
		timestamp,
		orgData.Name,
		timestamp.Format(time.RFC3339), // @timestamp
		orgData.Tags,                   // t
		map[string]*json.RawMessage{orgData.Name: orgData.Fields}, //m
	}

	return result, nil
}

func (conv) SetConfig(config map[string]interface{}) error {
	return nil
}
