package errors

import (
	"fmt"
	"github.com/dotuancd/ezserve/app/supports/str"
	v "gopkg.in/go-playground/validator.v8"
	"strings"
)

type validationFormatter func(f *v.FieldError) interface{}

var validationFormatters = map[string]validationFormatter{
	"required_without": FormatRequiredWithout,
}

var defaultValidationFormatter = func(f *v.FieldError) interface{} {
	return replacements("The :attribute failed validation :tag", f)
}

func FormatRequiredWithout(f *v.FieldError) interface{} {
	return replacements("The :attribute required when :param not represent", f)
}

func formatValidationErrors(err v.ValidationErrors) map[string]interface{} {
	fields := map[string]interface{}{}

	for _, fError := range err {
		// fields: {email: "The email fail validation required"}
		name := strings.ToLower(fError.Name)
		
		if sField, ok := fError.Type.FieldByName(fError.Name); ok {
			fmt.Print(sField)
		}

		var formatter validationFormatter

		formatter, hasFormatter := validationFormatters[fError.Tag]

		if !hasFormatter {
			formatter = defaultValidationFormatter
		}

		fields[name] = formatter(fError)
	}

	return fields
}

func replacements(template string, f *v.FieldError) string {
	return str.Replacements(template, map[string]interface{}{
		"attribute": strings.ToLower(f.Name),
		"Attribute": str.UpperFirst(f.Name),
		"ATTRIBUTE": strings.ToUpper(f.Name),
		"TAG": strings.ToUpper(f.Tag),
		"tag": strings.ToLower(f.Tag),
		"Tag": str.UpperFirst(f.Tag),
		"PARAM": strings.ToUpper(f.Param),
		"param": strings.ToLower(f.Param),
		"Param": str.UpperFirst(f.Param),
	})
}


