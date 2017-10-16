package telegraf_json

import (
	"bytes"
	"sync"
	"time"

	"github.com/buger/jsonparser"

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
	ts, err := jsonparser.GetInt(d, "timestamp")
	if err != nil {
		return nil, err
	}
	timestamp := time.Unix(ts, 0)

	name, err := jsonparser.GetString(d, "name")
	if err != nil {
		return nil, err
	}

	tags, elemType, _, err := jsonparser.Get(d, "tags")
	if err != nil || elemType == jsonparser.NotExist {
		return nil, err
	}
	fields, elemType, _, err := jsonparser.Get(d, "fields")
	if err != nil || elemType == jsonparser.NotExist {
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

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
