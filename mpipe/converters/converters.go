package converters

import (
	"fmt"

	"github.com/dialogbox/mpipego/mpipe/converters/influxlp"

	"github.com/dialogbox/mpipego/common"
	"github.com/dialogbox/mpipego/mpipe/converters/telegraf_json"
)

// New is factory method
func New(name string) (common.Converter, error) {
	switch name {
	case "telegraf_json":
		return telegraf_json.NewConverter(), nil
	case "influxlp":
		return influxlp.NewConverter(), nil
	default:
		return nil, fmt.Errorf("No such converter: %v", name)
	}
}
