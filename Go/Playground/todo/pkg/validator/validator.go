package validator

import (
	"fmt"
	"regexp"
)

type ValidationResult struct {
	Validations []string
}

func (v ValidationResult) IsValid() bool {
	return len(v.Validations) == 0
}

func newValidationResult(validations []*string) *ValidationResult {
	v := &ValidationResult{}

	for _, validation := range validations {
		if validation != nil {
			v.Validations = append(v.Validations, *validation)
		}
	}

	return v
}

func Validate(validations ...*string) *ValidationResult {
	return newValidationResult(validations)
}

func ValidateLength(key, value string, max int) *string {
	if len(value) > max {
		r := fmt.Sprintf("%s must be less than or equal to %d characters", key, max)
		return &r
	}
	return nil

}

func ValidateRequired(key string, value string) *string {

	if value == "" {
		r := fmt.Sprintf("%s can't be blank", key)
		return &r
	}
	return nil

}

func ValidateRange(key, value string, min, max int) *string {
	l := len(value)

	if l < min || l > max {
		r := fmt.Sprintf("%s must be between %d and %d", key, min, max)
		return &r
	}

	return nil
}

func ValidateEmail(key, value string) *string {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)

	if ok := re.Match([]byte(value)); ok {
		return nil
	}

	r := key + ": invalid"
	return &r
}
