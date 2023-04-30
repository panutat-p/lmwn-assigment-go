package src

import (
	"encoding/json"
	"reflect"
	"strconv"
)

type NanString string

func (s *NanString) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	items := reflect.ValueOf(v)
	switch items.Kind() {
	case reflect.String:
		if items.String() == "Nan" {
			return nil
		}
		*s = NanString(items.String())
	case reflect.Float64:
		// number field in JSON
		*s = NanString(strconv.FormatInt(int64(items.Float()), 10))
	default:
		// null
	}
	return nil
}

type NullInt int

func (i *NullInt) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	items := reflect.ValueOf(v)
	switch items.Kind() {
	case reflect.Float64:
		// number field in JSON
		*i = NullInt(items.Float())
	default:
		// null, Nan, string
		*i = -1
	}
	return nil
}
