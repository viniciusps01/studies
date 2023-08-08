package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	Values url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		Values: data,
		Errors: errors{},
	}
}

func (f *Form) Has(field string) bool {
	value := f.Values.Get(field)

	return value != ""
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Values.Get(field)
		trimmedValue := strings.TrimSpace(value)

		if trimmedValue == "" {
			f.Errors.Add(field, "This field shouldn't be blank")
		}
	}
}

func (f *Form) MinLength(field string, length int) bool {
	value := f.Values.Get(field)
	size := len(value)

	hasMinLength := size >= length

	if !hasMinLength {
		message := fmt.Sprintf("This field should have at least %v characters", length)
		f.Errors.Add(field, message)
		return false
	}

	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) ValidateEmail(field string) {
	email := f.Values.Get(field)

	if !govalidator.IsEmail(email) {
		f.Errors.Add(field, "Invalid email address")
	}
}
