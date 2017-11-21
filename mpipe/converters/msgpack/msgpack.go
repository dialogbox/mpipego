package msgpack

import (
	"fmt"

	"github.com/dialogbox/mpipego/common"
	"github.com/pkg/errors"
)

func NewConverter() common.Converter {
	return &conv{}
}

type conv struct{}

func (conv) Name() string {
	return "msgpack"
}

func (conv) Convert(d []byte) (*common.Metric, error) {
	m := &Metric{}
	left, err := m.UnmarshalMsg(d)

	if err != nil {
		return nil, err
	}

	if len(left) != 0 {
		return nil, fmt.Errorf("Trailing garbage data has been received")
	}

	return &common.Metric{
		Name:      m.Name,
		Timestamp: m.Time,
		Tags:      m.Tags,
		Fields:    m.Fields,
	}, nil
}

func (conv) SetConfig(config map[string]interface{}) error {
	return nil
}

func (conv) Encode(om *common.Metric) ([]byte, error) {
	m := &Metric{
		Name:   om.Name,
		Time:   om.Timestamp,
		Tags:   om.Tags,
		Fields: om.Fields,
	}

	b, err := m.MarshalMsg(nil)
	if err != nil {
		return nil, errors.Wrap(err, "Can not marshal metric")
	}

	return b, nil
}
