package genesis

import (
	"github.com/syke99/genesis/internal/cmd"
	"reflect"
)

func Spawn[T any](m map[string]any) (*T, error) {
	s := new(T)

	v := reflect.ValueOf(s).Elem()

	rV, err := cmd.Spawn(v, m)
	if err != nil {
		return nil, err
	}

	v = rV.(reflect.Value)

	s = v.Addr().Interface().(*T)

	return s, nil
}
