package helpers

import (
	"reflect"
)

func CopyStructByFields(src interface{}, fields ...string) interface{} {
	srcValue := reflect.ValueOf(src)
	srcType := srcValue.Type()

	var destFields []reflect.StructField

	for _, fieldName := range fields {
		field, found := srcType.FieldByName(fieldName)
		if found {
			destFields = append(destFields, reflect.StructField{
				Name: fieldName,
				Type: field.Type,
				Tag:  field.Tag,
			})
		}
	}

	destType := reflect.StructOf(destFields)

	dest := reflect.New(destType).Elem()

	for _, fieldName := range fields {
		_, found := srcType.FieldByName(fieldName)
		if found {
			destField := dest.FieldByName(fieldName)
			destField.Set(srcValue.FieldByName(fieldName))
		}
	}
	return dest.Interface()
}
