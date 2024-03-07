package genesis

import (
	"errors"
	"fmt"
	"reflect"
)

func nested(f reflect.Value, m map[string]any) (any, error) {
	for mn, mv := range m {
		fl := f.FieldByName(mn)

		if fl.CanSet() {
			flT := reflect.TypeOf(fl)

			mvT := reflect.TypeOf(mv)

			if flT.Kind() == reflect.Struct && mvT.Kind() == reflect.Map {
				switch mv.(type) {
				case map[string]any:
					v, err := nested(fl, mv.(map[string]any))
					if err != nil {
						return nil, err
					}
					fl.Set(v.(reflect.Value))
				default:
					return nil, errors.New(fmt.Sprintf("map provided for nested struct %s not of type map[string]any", flT.Name()))
				}
			} else {
				fl.Set(reflect.ValueOf(mv))
			}
		} else {
			return nil, errors.New(fmt.Sprintf("field %s is either unexported or unaddressable", mn))
		}
	}

	return f, nil
}

func Spawn[T any](m map[string]any) (*T, error) {
	s := new(T)

	v := reflect.ValueOf(s).Elem()

	rV, err := nested(v, m)
	if err != nil {
		return nil, err
	}

	v = rV.(reflect.Value)

	intermediate := v.Interface().(T)

	return &intermediate, err
}
