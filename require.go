package require_field

import (
	"errors"
	"log"
	"reflect"
	"strings"
)

func Check(obj interface{}) error {
	var val reflect.Value
	val = reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	if typ.Kind() != reflect.Struct {
		log.Printf("WARN: %T is not a struct", obj)
		// it is not a struct
		return nil
	}

	emptyReqFields := checkFields(obj)

	if len(emptyReqFields) > 0 {
		return errors.New("required fields are empty: " + strings.Join(emptyReqFields, ", "))
	}

	return nil
}

func checkFields(obj interface{}) []string {
	var val reflect.Value
	val = reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	if typ.Kind() != reflect.Struct {
		// it is not a struct
		return nil
	}

	emptyReqFields := make([]string, 0)

	for i := 0; i < val.NumField(); i++ {
		if typ.Field(i).Type.Kind() == reflect.Struct {
			emptyReqFields = append(emptyReqFields, checkFields(val.Field(i).Interface())...)
			continue
		}

		tag := typ.Field(i).Tag.Get("required")
		if tag != "" {
			if typ.Field(i).Type.Kind() == reflect.Slice || typ.Field(i).Type.Kind() == reflect.Map {
				if val.Field(i).Len() == 0 {
					emptyReqFields = append(emptyReqFields, typ.Field(i).Name)
				}
				continue
			}

			if val.Field(i).Interface() == reflect.Zero(typ.Field(i).Type).Interface() {
				emptyReqFields = append(emptyReqFields, typ.Field(i).Name)
			}
		}
	}

	return emptyReqFields
}
