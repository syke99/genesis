package pkg

import (
	"errors"
	"fmt"
)

func BadMap(field string) error {
	return errors.New(fmt.Sprintf("map provided for field %s not of type map[string]any or map[string]func(ctx context.Context) any", field))
}

func WrongSpawnForMap(field string) error {
	return errors.New(fmt.Sprintf("Spawn called on field %s with provided map map[string]func(ctx context.Context) any, should use SpawnWithContext instead", field))
}

func UnaddressableField(field string) error {
	return errors.New(fmt.Sprintf("field %s is unaddressable", field))
}
