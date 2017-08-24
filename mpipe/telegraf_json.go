package mpipe

import (
	"encoding/json"
	"time"
)

type telegrafJSONConverter struct {
}

func (c *telegrafJSONConverter) Name() string {
	return "telegraf_json"
}

// ConvertTelegrafJSON Convert telegram generated JSON
func (c *telegrafJSONConverter) Convert(d []byte) (*MetricData, error) {
	var f interface{}
	if err := json.Unmarshal(d, &f); err != nil {
		return nil, err
	}

	m := f.(map[string]interface{})

	m[m["name"].(string)] = m["fields"]
	timestamp := time.Unix(int64(m["timestamp"].(float64)), 0)
	delete(m, "fields")
	delete(m, "timestamp")

	m["@timestamp"] = timestamp.Format(time.RFC3339)

	return &MetricData{timestamp, m}, nil
}

func (c *telegrafJSONConverter) SetConfig(config map[string]interface{}) error {
	return nil
}
