package telegraf_json

import (
	"time"

	"github.com/json-iterator/go"

	"github.com/dialogbox/mpipego/common"
)

func NewConverter() common.Converter {
	return &conv{}
}

type conv struct{}

func (conv) Name() string {
	return "telegraf_json"
}

// ConvertTelegrafJSON Convert telegram generated JSON
func (conv) Convert(d []byte) (*common.Metric, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	m := &struct {
		Name      string
		Timestamp int64
		Tags      map[string]string
		Fields    map[string]interface{}
	}{}
	err := json.Unmarshal(d, m)
	if err != nil {
		return nil, err
	}

	return &common.Metric{
		Name:      m.Name,
		Timestamp: time.Unix(m.Timestamp, 0),
		Tags:      m.Tags,
		Fields:    m.Fields,
	}, err
}

func (conv) SetConfig(config map[string]interface{}) error {
	return nil
}
