package utils

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func GetEnv[T any](key string, defaultValue T) T {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	var result T
	switch any(result).(type) {
	case string:
		return any(valueStr).(T)
	case int:
		if val, err := strconv.Atoi(valueStr); err == nil {
			return any(val).(T)
		}
	case int64:
		if val, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
			return any(val).(T)
		}
	case bool:
		if val, err := strconv.ParseBool(valueStr); err == nil {
			return any(val).(T)
		}
	case []int64:
		parts := strings.Split(valueStr, ",")
		ids := make([]int64, 0, len(parts))
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if id, err := strconv.ParseInt(part, 10, 64); err == nil {
				ids = append(ids, id)
			}
		}
		return any(ids).(T)
	case []string:
		parts := strings.Split(valueStr, ",")
		result := make([]string, len(parts))
		for i, part := range parts {
			result[i] = strings.TrimSpace(part)
		}
		return any(result).(T)
	case time.Duration:
		if val, err := time.ParseDuration(valueStr); err == nil {
			return any(val).(T)
		}
	default:
	}

	return defaultValue
}

func GetEnvRequired(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Required environment variable " + key + " is not set")
	}
	return value
}
