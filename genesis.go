package genesis

import (
	"context"
	"reflect"

	"github.com/syke99/genesis/internal/cmd"
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

func SpawnWithContext[T any](ctx context.Context, m map[string]func(ctx context.Context) any) (*T, error) {
	s := new(T)

	v := reflect.ValueOf(s).Elem()

	rV, err := cmd.SpawnWithContext(ctx, v, m)
	if err != nil {
		return nil, err
	}

	v = rV.(reflect.Value)

	s = v.Addr().Interface().(*T)

	return s, nil
}

// TODO: implement lazy spawner
//type LazySpawner[F any] interface {
//	GetField(field string, ctx context.Context, args ...any) F
//}
//
//func LazySpawn[L LazySpawner[any], F any](m map[string]func(ctx context.Context, args ...any) F) (*L, error) {
//	s := new(L)
//
//	v := reflect.ValueOf(s).Elem()
//
//	rV, err := cmd.LazySpawn(v, m)
//	if err != nil {
//		return nil, err
//	}
//
//	v = rV.(reflect.Value)
//
//	s = v.Addr().Interface().(*L)
//
//	return s, nil
//}
