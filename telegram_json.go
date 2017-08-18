package mpipego

import (
	"encoding/json"
	"time"
)

// ConvertTelegramJSON Convert telegram generated JSON
func ConvertTelegramJSON(d []byte) *MetricData {
	var f interface{}
	if err := json.Unmarshal(d, &f); err != nil {
		panic(err)
	}

	m := f.(map[string]interface{})

	m[m["name"].(string)] = m["fields"]
	timestamp := time.Unix(int64(m["timestamp"].(float64)), 0)
	delete(m, "fields")
	delete(m, "timestamp")

	m["@timestamp"] = timestamp.Format(time.RFC3339)

	return &MetricData{timestamp, m }
}
