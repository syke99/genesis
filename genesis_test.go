package genesis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestType struct {
	FieldOne   string
	FieldTwo   int
	fieldThree bool
	Nested     struct {
		FieldFour string
		FieldFive int
		FieldSix  bool
	}
}

type NestedTestType struct {
	FieldFour string
	FieldFive int
	FieldSix  bool
}

func TestSpawnNestedTestType(t *testing.T) {
	v, err := Spawn[TestType](map[string]any{
		"FieldOne": "hello world",
		"FieldTwo": 1234,
		"Nested": NestedTestType{
			FieldFour: "what's up!",
			FieldFive: 5678,
			FieldSix:  true,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "hello world", v.FieldOne)
	assert.Equal(t, 1234, v.FieldTwo)
	assert.Equal(t, "what's up!", v.Nested.FieldFour)
	assert.Equal(t, 5678, v.Nested.FieldFive)
	assert.Equal(t, true, v.Nested.FieldSix)
}

func TestSpawnNestedTestTypeFromMap(t *testing.T) {
	v, err := Spawn[TestType](map[string]any{
		"FieldOne": "hello world",
		"FieldTwo": 1234,
		"Nested": map[string]any{
			"FieldFour": "what's up!",
			"FieldFive": 5678,
			"FieldSix":  true,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "hello world", v.FieldOne)
	assert.Equal(t, 1234, v.FieldTwo)
	assert.Equal(t, "what's up!", v.Nested.FieldFour)
	assert.Equal(t, 5678, v.Nested.FieldFive)
	assert.Equal(t, true, v.Nested.FieldSix)
}

func TestSpawnUnexportedField(t *testing.T) {
	_, err := Spawn[TestType](map[string]any{
		"FieldOne":   "goodbye everyone",
		"FieldTwo":   5678,
		"fieldThree": true,
	})
	assert.Error(t, err)
	assert.Equal(t, "field fieldThree is either unexported or unaddressable", err.Error())
}
