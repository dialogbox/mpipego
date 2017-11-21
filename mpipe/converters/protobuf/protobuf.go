package protobuf

import (
	fmt "fmt"
	"time"

	"github.com/dialogbox/mpipego/common"
	proto "github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func NewConverter() common.Converter {
	return &conv{}
}

type conv struct{}

func (conv) Name() string {
	return "protobuf"
}

func (conv) Convert(d []byte) (*common.Metric, error) {
	m := &Metric{}
	err := proto.Unmarshal(d, m)
	if err != nil {
		return nil, err
	}

	timestamp := time.Unix(0, m.GetTimestamp())

	name := m.GetName()

	tags := m.GetTags()

	fields, err := convertFields(m)
	if err != nil {
		return nil, err
	}

	return &common.Metric{
		Timestamp: timestamp,
		Name:      name,
		Tags:      tags,
		Fields:    fields,
	}, nil
}

func (conv) SetConfig(config map[string]interface{}) error {
	return nil
}

func (conv) Encode(om *common.Metric) ([]byte, error) {
	f, err := convertFieldsRev(om.Fields)
	if err != nil {
		return nil, errors.Wrap(err, "Can not convert fields to protobuf")
	}

	m := &Metric{
		Name:      om.Name,
		Timestamp: om.Timestamp.UnixNano(),
		Tags:      om.Tags,
		Fields:    f,
	}

	b, err := proto.Marshal(m)
	if err != nil {
		return nil, errors.Wrap(err, "Can not marshal metric to protobuf")
	}

	return b, nil
}

func convertFields(m *Metric) (map[string]interface{}, error) {
	f := make(map[string]interface{})

	for k, v := range m.GetFields() {
		switch v := v.GetValue().(type) {
		case *FieldValue_IntValue:
			f[k] = v.IntValue
		case *FieldValue_FloatValue:
			f[k] = v.FloatValue
		case *FieldValue_StringValue:
			f[k] = v.StringValue
		case *FieldValue_BoolValue:
			f[k] = v.BoolValue
		default:
			return nil, fmt.Errorf("Unsupported field data type: [%s, %v(%T)]", k, v, v)
		}
	}
	return f, nil
}

func convertFieldsRev(of map[string]interface{}) (map[string]*FieldValue, error) {
	f := make(map[string]*FieldValue)

	for k, v := range of {
		switch v := v.(type) {
		case int64:
			f[k] = &FieldValue{&FieldValue_IntValue{IntValue: v}}
		case int:
			f[k] = &FieldValue{&FieldValue_IntValue{IntValue: int64(v)}}
		case float64:
			f[k] = &FieldValue{&FieldValue_FloatValue{FloatValue: v}}
		case string:
			f[k] = &FieldValue{&FieldValue_StringValue{StringValue: v}}
		case bool:
			f[k] = &FieldValue{&FieldValue_BoolValue{BoolValue: v}}
		default:
			return nil, fmt.Errorf("Unsupported field data type: [%s, %v(%T)]", k, v, v)
		}
	}

	return f, nil
}
