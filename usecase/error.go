package usecase

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	tags = map[string]string{
		"required":   "is required",
		"gt":         "must be greater than",
		"gte":        "must be greater or equal to",
		"lt":         "must be smaller than",
		"ltecsfield": "must be smaller than",
		"oneof":      "must be one of",
	}
)

func getMessage(errors []validator.FieldError) (msg string) {

	for _, er := range errors {
		tag := tags[er.Tag()]
		field := er.StructField()
		param := er.Param()

		if len(msg) > 0 {
			msg = strings.TrimSpace(fmt.Sprintf("%s, %s %s %s", msg, field, tag, param))
			continue
		}
		msg = strings.TrimSpace(fmt.Sprintf("%s %s %s", field, tag, param))
	}

	return msg
}
