package mpipego

import (
	"fmt"
	"time"
)

type Converter func ([]byte) *MetricData

type MetricData struct {
	Timestamp time.Time
	Data      interface{}
}

func ConverterForType(typeName string) Converter {
	switch typeName {
	case "telegraf_json":
		return ConvertTelegramJSON
	default:
		panic(fmt.Sprintf("No converter for type: %v", typeName))
	}
}