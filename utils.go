package gokalshi

import (
	"errors"
	"fmt"
	"strconv"
)

func Sprintf(t any) string {
	s, _ := SprintfWithError(t, false)
	return s
}

func SprintfWithError(t any, requireNonNil bool) (string, error) {
	if t == nil && requireNonNil {
		return "", errors.New("value cannot be nil")
	}
	switch v := t.(type) {
	case string:
		return v, nil
	case int:
		return strconv.Itoa(v), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(v), nil
	default:
		return fmt.Sprintf("%v", v), nil
	}
}
