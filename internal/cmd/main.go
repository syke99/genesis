package cmd

import (
	"context"
	"reflect"
	"unsafe"

	"github.com/syke99/genesis/internal/pkg"
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

func SpawnWithContext(ctx context.Context, f reflect.Value, m map[string]func(ctx context.Context) any) (any, error) {
	for mn, mv := range m {
		fl := f.FieldByName(mn)

		if fl.CanSet() {
			flT := reflect.TypeOf(fl)

			res := mv(ctx)

			mvT := reflect.TypeOf(res)

			if flT.Kind() == reflect.Struct && mvT.Kind() == reflect.Map {
				switch res.(type) {
				case map[string]any:
					v, err := Spawn(fl, res.(map[string]any))
					if err != nil {
						return nil, err
					}
					fl.Set(v.(reflect.Value))
				case map[string]func(ctx context.Context) any:
					v, err := SpawnWithContext(ctx, fl, res.(map[string]func(ctx context.Context) any))
					if err != nil {
						return nil, err
					}
					fl.Set(v.(reflect.Value))
				default:
					return nil, pkg.BadMap(mn)
				}
			} else {
				fl.Set(reflect.ValueOf(mv(ctx)))
			}
		} else {
			if fl.CanAddr() {
				fl = reflect.NewAt(fl.Type(), unsafe.Pointer(fl.UnsafeAddr())).Elem()
				fl.Set(reflect.ValueOf(mv(ctx)))
			} else {
				return nil, pkg.UnaddressableField(mn)
			}
		}
	}

	return f, nil
}

// TODO: implement lazy spawner
//func LazySpawn[F any](f reflect.Value, m map[string]func(ctx context.Context, args ...any) F) (any, error) {
//	for mn, mv := range m {
//		fl := f.FieldByName(mn)
//
//		if fl.CanSet() {
//			fl.Set(reflect.ValueOf(mv))
//		} else {
//			if fl.CanAddr() {
//				fl = reflect.NewAt(fl.Type(), unsafe.Pointer(fl.UnsafeAddr())).Elem()
//				fl.Set(reflect.ValueOf(mv))
//			} else {
//				return nil, pkg.UnaddressableField(mn)
//			}
//		}
//	}
//
//	return f, nil
//}
