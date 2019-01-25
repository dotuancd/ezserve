package validation

import (
	. "gopkg.in/go-playground/validator.v8"
	"reflect"
)

func RequiredWithout(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {

	compareWithField, compareKind, ok := v.GetStructFieldOK(currentStruct, param)

	if !ok {
		return false
	}

	hasValue := HasValue(v, topStruct, currentStruct, compareWithField, compareWithField.Type(), compareKind, param)
	targetHasValue := HasValue(v, topStruct, currentStruct, field, fieldtype, fieldKind, param)

	return hasValue || targetHasValue
}