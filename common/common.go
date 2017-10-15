package common

import "time"

// Converter interface
type Converter interface {
	Convert([]byte) (Metric, error)
	Name() string
}

type Metric interface {
	Timestamp() time.Time
	Data() interface{}
	JSON() []byte
}
