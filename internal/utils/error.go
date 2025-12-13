package utils

import (
	"fmt"
	"log"
	"regexp"
	"runtime"
	"strings"
)

func WrapError(err error) error {
	return fmt.Errorf("%s %w", getCallerName(2), err)
}

func WrapErrorf(err error, text string, args ...interface{}) error {
	return fmt.Errorf("%s %s: %w", getCallerName(2), fmt.Sprintf(text, args...), err)
}

func WrapFunctionError(fn func() error) error {
	if err := fn(); err != nil {
		log.Printf("i'm here, but error exists: %t", err != nil)
		return fmt.Errorf("%s %w", getCallerName(2), err)
	}
	return nil
}

func WrapFunction[T any](fn func() (T, error)) (T, error) {
	value, err := fn()
	if err != nil {
		var zero T
		return zero, fmt.Errorf("%s %w", getCallerName(2), err)
	}
	return value, nil
}

func getCallerName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}

	fullName := fn.Name()

	// Извлекаем только имя метода
	parts := strings.Split(fullName, ".")
	if len(parts) > 0 {
		return fmt.Sprintf("[%s]", toSnakeCaseRegex(parts[len(parts)-1]))
	}

	return fmt.Sprintf("[%s]", toSnakeCaseRegex(fullName))
}

func toSnakeCaseRegex(s string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
