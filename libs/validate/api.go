package validate

import (
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

func ValidateApi(i interface{}) (error, map[string]string) {
	errMap := make(map[string]string)

	if err := validate.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			ss := strings.SplitN(e.Translate(trans), "-", 2)
			errMap[ss[0]] = ss[1]
		}

		return err, errMap
	}
	return nil, nil
}
