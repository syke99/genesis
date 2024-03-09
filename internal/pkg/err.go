package pkg

import (
	"errors"
	"fmt"
)

func BadMap(field string) error {
	return errors.New(fmt.Sprintf("map provided for field %s not of type map[string]any or map[string]func(ctx context.Context) any", field))
}

func UnaddressableField(field string) error {
	return errors.New(fmt.Sprintf("field %s is unaddressable", field))
}
