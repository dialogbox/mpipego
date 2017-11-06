package msgpack

import (
	"fmt"

	"github.com/dialogbox/mpipego/common"
)

func NewConverter() common.Converter {
	return &conv{}
}

type conv struct{}

func (conv) Name() string {
	return "msgpack"
}

func (conv) Convert(d []byte) (common.Metric, error) {
	m := &Metric{}
	left, err := m.UnmarshalMsg(d)

	if err != nil {
		return nil, err
	}

	if len(left) != 0 {
		return nil, fmt.Errorf("Trailing garbage data has been received")
	}

	return m, nil
}

func (conv) SetConfig(config map[string]interface{}) error {
	return nil
}
