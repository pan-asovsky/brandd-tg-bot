package utils

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

func Error(err error) error {
	return fmt.Errorf("%s %w", GetCallerName(), err)
}

func GetCallerName() string {
	pc, _, _, ok := runtime.Caller(1)
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
		return fmt.Sprintf("[%s]", ToSnakeCaseRegex(parts[len(parts)-1]))
	}

	return fmt.Sprintf("[%s]", ToSnakeCaseRegex(fullName))
}

func ToSnakeCaseRegex(s string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
