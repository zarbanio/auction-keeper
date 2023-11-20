package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Log struct {
	Id      int64  `json:"id"`
	Level   string `json:"string"`
	Service string `json:"service"`
	Message string `json:"message"`
	Fields  string `json:"fields"`
}

// Scan - Implement the database/sql Scanner interface
func (s *Log) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	data, ok := value.([]byte)
	if !ok {
		return errors.New("Scan received invalid type: not a byte slice")
	}

	err := json.Unmarshal(data, s)
	if err != nil {
		return err
	}

	return nil
}

func (l Log) Value() (driver.Value, error) {
	if l == (Log{}) {
		return nil, nil
	}

	data, err := json.Marshal(l)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Value failed to marshal JSON: %v", err))
	}

	return data, nil
}
