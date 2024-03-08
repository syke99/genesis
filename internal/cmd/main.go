package cmd

import (
	"github.com/syke99/genesis/internal/pkg"
	"reflect"
	"unsafe"
)

func Spawn(f reflect.Value, m map[string]any) (any, error) {
	for mn, mv := range m {
		fl := f.FieldByName(mn)

		if fl.CanSet() {
			flT := reflect.TypeOf(fl)

			mvT := reflect.TypeOf(mv)

			if flT.Kind() == reflect.Struct && mvT.Kind() == reflect.Map {
				switch mv.(type) {
				case map[string]any:
					v, err := Spawn(fl, mv.(map[string]any))
					if err != nil {
						return nil, err
					}
					fl.Set(v.(reflect.Value))
				default:
					return nil, pkg.BadMap(mn)
				}
			} else {
				fl.Set(reflect.ValueOf(mv))
			}
		} else {
			if fl.CanAddr() {
				fl = reflect.NewAt(fl.Type(), unsafe.Pointer(fl.UnsafeAddr())).Elem()
				fl.Set(reflect.ValueOf(mv))
			} else {
				return nil, pkg.UnaddressableField(mn)
			}

		}
	}

	return f, nil
}
