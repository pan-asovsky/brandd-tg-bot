package service

import (
	"fmt"
	"regexp"
	"strings"

	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
)

type phoneService struct {
	normalizeRegex *regexp.Regexp
	detectRegex    *regexp.Regexp
}

func NewPhoneNormalizingService() i.PhoneService {
	return &phoneService{normalizeRegex: regexp.MustCompile(`\d`), detectRegex: regexp.MustCompile(`\D`)}
}

func (p *phoneService) Normalize(phone string) (string, error) {
	matches := p.normalizeRegex.FindAllString(phone, -1)
	if matches == nil {
		return "", fmt.Errorf("[phone_normalize] no numbers in phone: %s", phone)
	}

	digits := strings.Join(matches, "")

	return p.validateAndFormat(digits)
}

func (p *phoneService) Detect(text string) (string, bool) {
	digits := p.detectRegex.ReplaceAllString(text, "")

	if len(digits) == 10 && digits[0] == '9' {
		return "8" + digits, true
	}

	if len(digits) == 11 && (digits[0] == '7' || digits[0] == '8') {
		if digits[0] == '7' {
			return "8" + digits[1:], true
		}
		return digits, true
	}

	return "", false
}

func (p *phoneService) validateAndFormat(digits string) (string, error) {
	length := len(digits)

	if length == 10 {
		return p.formatTenDigit(digits)
	} else if length == 11 {
		return p.formatElevenDigit(digits)
	} else {
		return "", fmt.Errorf("[phone_normalize] invalid numbers count: %d (excepted 10 or 11)", length)
	}
}

func (p *phoneService) formatTenDigit(digits string) (string, error) {
	if digits[0] != '9' {
		return "", fmt.Errorf("[phone_normalize] 10-digit number must start from 9: %s", digits)
	}
	return "8" + digits, nil
}

func (p *phoneService) formatElevenDigit(digits string) (string, error) {
	firstDigit := digits[0]

	if firstDigit == '8' {
		return digits, nil
	} else if firstDigit == '7' {
		return "8" + digits[1:], nil
	} else {
		return "", fmt.Errorf("[phone_normalize] 11-digit number must start with 7 or 8: %s", digits)
	}
}
