package tools

import (
	"fmt"
)

type FieldError struct {
	Field string
	Msg   string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func CheckFieldExistance(fields map[string]any) (err error) {
	for field, value := range fields {
		if value == "" || value == 0 || value == 0.0 {
			err = &FieldError{
				Field: field,
				Msg:   "field is required",
			}
			return
		}
	}
	return
}
