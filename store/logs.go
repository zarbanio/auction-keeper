package store

import (
	"context"
	"encoding/json"
	"errors"
)

func (p postgres) CreateSysLog(ctx context.Context, b []byte) (int, error) {

	var entry map[string]interface{}
	err := json.Unmarshal(b, &entry)
	if err != nil {
		return 0, err
	}

	level := ConvertToString(entry["level"])
	if level == "" {
		return 0, errors.New("log level didn't found")
	}

	if level == "debug" {
		return len(b), nil
	}

	message := ConvertToString(entry["message"])
	if message == "" {
		return 0, errors.New("log message didn't found")
	}

	service := ConvertToString(entry["service"])
	if service == "" {
		return 0, errors.New("log message didn't found")
	}

	fieldsJSON, err := json.Marshal(entry)
	if err != nil {
		return 0, err
	}

	stmt := `
		INSERT INTO logs (level, service, message, fields) 
		VALUES ($1, $2, $3, $4)
	`

	_, err = p.conn.Exec(ctx, stmt, level, service, message, string(fieldsJSON))
	if err != nil {
		return 0, err
	}

	return len(b), nil
}
