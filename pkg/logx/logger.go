package logx

import (
	"encoding/json"
	"fmt"
	"time"
)

type Entry struct {
	Level   string         `json:"level"`
	Message string         `json:"message"`
	Time    string         `json:"time"`
	Fields  map[string]any `json:"fields,omitempty"`
}

func Print(level, message string, fields map[string]any) error {
	entry := Entry{
		Level:   level,
		Message: message,
		Time:    time.Now().Format(time.RFC3339),
		Fields:  fields,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}
