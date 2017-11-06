package influxlp

// https://docs.influxdata.com/influxdb/v1.3/write_protocols/line_protocol_tutorial/

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dialogbox/mpipego/common"
)

func NewConverter() common.Converter {
	return &conv{}
}

type conv struct{}

func (conv) Name() string {
	return "influxlp"
}

func (conv) Convert(d []byte) (*common.Metric, error) {
	measureSlice, tagsSlice, fieldsSlice, timestampSlice := quickOverlook(d)
	if measureSlice == nil {
		return nil, fmt.Errorf("Influx LP parse error: invalid format [%s]", string(d))
	}

	// Convert measure
	name := string(measureSlice)

	// Parse tags
	tags := make(map[string]string)
	if len(tagsSlice) > 0 {
		for {
			advance, k, v, err := influxLPReadTag(tagsSlice)
			if err != nil {
				return nil, fmt.Errorf("Influx LP parse error: can not parse tags [%v]", err)
			}

			if advance > 0 {
				tags[string(k)] = string(v)
			}

			if advance == 0 || advance >= len(tagsSlice) {
				break
			}

			tagsSlice = tagsSlice[advance:]
		}
	}

	// Parse fields
	fields := make(map[string]interface{})
	for {
		advance, k, v, err := influxLPReadField(fieldsSlice)
		if err != nil {
			return nil, fmt.Errorf("Influx LP parse error: can not parse fields [%v]", err)
		}

		if advance > 0 {
			fields[string(k)] = v
		}

		if advance == 0 || advance >= len(fieldsSlice) {
			break
		}

		fieldsSlice = fieldsSlice[advance:]
	}

	// Convert timestamp
	var timestamp time.Time
	if len(timestampSlice) > 0 {
		ts, err := convertTimestamp(timestampSlice)
		if err != nil {
			return nil, fmt.Errorf("Influx LP parse error: can not parse timestamp [%v]", err)
		}
		timestamp = ts
	} else {
		timestamp = time.Now()
	}

	return &common.Metric{
		Name:      name,
		Timestamp: timestamp,
		Tags:      tags,
		Fields:    fields,
	}, nil
}

func (conv) SetConfig(config map[string]interface{}) error {
	return nil
}

// Overlook the input
func quickOverlook(l []byte) (measure []byte, tags []byte, fields []byte, timestamp []byte) {
	var i, p1, p2, p3 int
	length := len(l)

	for i = 0; i < length && (l[i] != ' '); i++ {
		if l[i] == '\\' {
			i++
		} else if p1 == 0 && l[i] == ',' {
			p1 = i // the first comma
		}
	}
	if i == length {
		return nil, nil, nil, nil
	}
	p2 = i // the first space

	for i++; i < length && (l[i] != ' '); i++ {
		if l[i] == '\\' {
			i++
		} else if l[i] == '"' {
			for i++; l[i] != '"'; i++ { // skip double quoted
				if l[i] == '\\' {
					i++
				}
			}
		}
	}
	if i < length {
		p3 = i // the second space
	}

	if p1 == 0 { // no tags
		measure = l[:p2]
	} else {
		measure = l[:p1]
		tags = l[p1+1 : p2]
	}
	if p3 == 0 { // no timestamp
		fields = l[p2+1:]
	} else {
		fields = l[p2+1 : p3]
		timestamp = l[p3+1:]
	}

	return
}

func _influxLPReadValue(data []byte) (advance int, value interface{}, err error) {
	if len(data) == 0 {
		return 0, nil, fmt.Errorf("Format error: empty value")
	}

	var token []byte
	if data[0] == '"' { // Double quoted string
		// Advances until the closing quote
		// 'advance' will point the next postion of closing quote
		quoteClosed := false
		for advance = 1; advance < len(data); advance++ {
			if data[advance] == '\\' {
				advance++
			} else if data[advance] == '"' {
				advance++
				quoteClosed = true
				break
			}
			token = append(token, data[advance])
		}

		// In case of the closing quote hasn't been founded
		if !quoteClosed {
			return 0, nil, fmt.Errorf("Format error: quote not matched: %s", string(data))
		}

		value = string(token)

		// In case of this was the final item
		if advance == len(data) {
			return
		}

		// In case of any unquoted chars are left after quote has closed
		if data[advance] != ',' && data[advance] != '=' {
			return 0, nil, fmt.Errorf("Format error: Unquoted data after quote has closed: input (%s), left (%s)", string(data), string(data[advance:]))
		}

		// Skip the delimiter
		advance++
	} else { // Unquoted string
		for advance = 0; advance < len(data); advance++ {
			if data[advance] == '\\' {
				advance++
			} else if data[advance] == ',' {
				advance++
				break
			}
			token = append(token, data[advance])
		}
		v, err := parseFieldValue(token)
		if err != nil {
			return 0, nil, err
		}
		value = v
	}

	return
}

func _influxLPReadKey(data []byte, delim byte) (advance int, token []byte, err error) {
	if len(data) == 0 {
		return
	}

	for advance = 0; advance < len(data); advance++ {
		switch data[advance] {
		case '\\':
			advance++
		case delim:
			advance++
			return
		}
		token = append(token, data[advance])
	}

	return
}

func influxLPReadTag(data []byte) (advance int, key []byte, value []byte, err error) {
	if len(data) == 0 {
		return
	}

	i, token, err := _influxLPReadKey(data, '=')
	if err != nil {
		return 0, nil, nil, err
	}
	data = data[i:]
	key = token

	j, token, err := _influxLPReadKey(data, ',')
	if err != nil {
		return 0, nil, nil, err
	}
	value = token

	advance = i + j

	return
}

func influxLPReadField(data []byte) (int, []byte, interface{}, error) {
	if len(data) == 0 {
		return 0, nil, nil, nil
	}

	i, key, err := _influxLPReadKey(data, '=')
	if err != nil {
		return 0, nil, nil, err
	}
	data = data[i:]

	j, value, err := _influxLPReadValue(data)
	if err != nil {
		return 0, nil, nil, err
	}

	return i + j, key, value, nil
}

func convertTimestamp(d []byte) (time.Time, error) {
	ts, err := strconv.ParseInt(string(d), 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("Format error: Invalid timestamp %s", string(d))
	}
	return time.Unix(0, ts), nil
}

func parseFieldValue(v []byte) (interface{}, error) {
	if v[len(v)-1] == 'i' {
		intv, err := strconv.ParseInt(string(v[:len(v)-1]), 10, 64)
		if err != nil {
			return nil, err
		}

		return intv, nil
	}

	switch v[0] {
	case 't':
		return true, nil
	case 'T':
		return true, nil
	case 'f':
		return false, nil
	case 'F':
		return false, nil
	}

	return strconv.ParseFloat(string(v), 64)
}