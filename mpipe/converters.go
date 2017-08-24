package mpipe

import (
	"fmt"
	"time"
)

// Converter interface
type Converter interface {
	Convert([]byte) (*MetricData, error)
	SetConfig(map[string]interface{}) error
	Name() string
}

// MetricData payload struct
type MetricData struct {
	Timestamp time.Time
	Data      interface{}
}

func (m *MetricData) String() string {
	rawstring := fmt.Sprintf("%v", m.Data)

	data, ok := m.Data.(map[string]interface{})
	if !ok {
		return rawstring
	}

	name, ok := data["name"]
	if !ok {
		return rawstring
	}

	tags, ok := data["tags"]
	if !ok {
		return rawstring
	}

	temp, ok := data[name.(string)]
	if !ok {
		return rawstring
	}

	metrics, ok := temp.(map[string]interface{})
	if !ok {
		return rawstring
	}

	var keys []string
	for key := range metrics {
		keys = append(keys, key)
	}

	return fmt.Sprintf("[%v] name: %s, meta: %v, keys: %v", m.Timestamp, name, tags, keys)
}

func NewConverter(name string) (Converter, error) {
	switch name {
	case "telegraf_json":
		return &telegrafJSONConverter{}, nil
	default:
		return nil, fmt.Errorf("No such converter: %v", name)
	}
}
