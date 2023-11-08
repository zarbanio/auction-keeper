package store

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/zarbanio/auction-keeper/domain"
)

type sysLogModel struct {
	id      int64  `json:"id"`
	level   string `json:"level"`
	message string `json:"message"`
	feilds  string `json:"fields"`
}

func (e sysLogModel) toDomain() *domain.SysLog {

	return &domain.SysLog{
		Id:      e.id,
		Level:   e.level,
		Message: e.message,
		Fields:  e.feilds,
	}
}

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

	fieldsJSON, err := json.Marshal(entry)
	if err != nil {
		return 0, err
	}
	stmt := `
		INSERT INTO logs (level, message, fields) 
		VALUES ($1, $2, $3)
	`

	_, err = p.conn.Exec(ctx, stmt, level, message, string(fieldsJSON))
	if err != nil {
		return 0, err
	}

	return len(b), nil
}
