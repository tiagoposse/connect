package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type StringArray []string

func (o *StringArray) Scan(src any) error {
	*o = strings.Split(src.(string), ",")

	return nil
}

func (o StringArray) Value() (driver.Value, error) {
	if len(o) == 0 {
		return nil, nil
	}

	return strings.Join(o, ","), nil
}

type Duration time.Duration 

func (d Duration) MarshalJSON() ([]byte, error) {
    return json.Marshal(time.Duration(d).String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
    var v interface{}
    if err := json.Unmarshal(b, &v); err != nil {
        return err
    }
    switch value := v.(type) {
    case float64:
        *d = Duration(time.Duration(value))
        return nil
    case string:
        tmp, err := time.ParseDuration(value)
        if err != nil {
            return err
        }
        *d = Duration(tmp)
        return nil
    default:
        return errors.New("invalid duration")
    }
}
