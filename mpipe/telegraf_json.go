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

type oldMetric struct {
	Timestamp int64            `json:"timestamp"`
	Name      string           `json:"name"`
	Tags      *json.RawMessage `json:"tags"`
	Fields    *json.RawMessage `json:"fields"`
}

type metricDataJSONPayload struct {
	Timestamp string                      `json:"@timestamp"`
	Tags      *json.RawMessage            `json:"t"`
	Fields    map[string]*json.RawMessage `json:"m"`
}

// ConvertTelegrafJSON Convert telegram generated JSON
func (c *telegrafJSONConverter) Convert(d []byte) (*MetricData, error) {
	var olddata oldMetric
	if err := json.Unmarshal(d, &olddata); err != nil {
		return nil, err
	}

	timestamp := time.Unix(olddata.Timestamp, 0)
	newdata := &metricDataJSONPayload{
		timestamp.Format(time.RFC3339), // @timestamp
		olddata.Tags,                   // t
		map[string]*json.RawMessage{olddata.Name: olddata.Fields}, //m
	}

	result := &MetricData{
		timestamp,
		newdata,
	}

	return result, nil
}

func (c *telegrafJSONConverter) SetConfig(config map[string]interface{}) error {
	return nil
}
