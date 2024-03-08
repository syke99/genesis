package genesis

import (
	"github.com/stretchr/testify/assert"
	"github.com/syke99/genesis/internal/pkg"
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

func TestSpawnTestType(t *testing.T) {
	v, err := Spawn[TestType](map[string]any{
		pkg.FieldOneStr: pkg.FieldOne,
		pkg.FieldTwoStr: pkg.FieldTwo,
	})
	assert.NoError(t, err)
	assert.Equal(t, pkg.FieldOne, v.FieldOne)
	assert.Equal(t, pkg.FieldTwo, v.FieldTwo)
}

func TestSpawnUnexportedField(t *testing.T) {
	v, err := Spawn[TestType](map[string]any{
		pkg.FieldOneStr:   pkg.GoodBye,
		pkg.FieldTwoStr:   pkg.FieldFive,
		pkg.FieldThreeStr: pkg.FieldThree,
	})
	assert.NoError(t, err)
	assert.Equal(t, pkg.GoodBye, v.FieldOne)
	assert.Equal(t, pkg.FieldFive, v.FieldTwo)
	assert.Equal(t, pkg.FieldThree, v.fieldThree)
}

func TestSpawnNestedTestType(t *testing.T) {
	v, err := Spawn[TestType](map[string]any{
		pkg.FieldOneStr: pkg.FieldOne,
		pkg.FieldTwoStr: pkg.FieldTwo,
		pkg.Nested: NestedTestType{
			FieldFour: pkg.FieldFour,
			FieldFive: pkg.FieldFive,
			FieldSix:  pkg.FieldSix,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, pkg.FieldOne, v.FieldOne)
	assert.Equal(t, pkg.FieldTwo, v.FieldTwo)
	assert.Equal(t, pkg.FieldFour, v.Nested.FieldFour)
	assert.Equal(t, pkg.FieldFive, v.Nested.FieldFive)
	assert.Equal(t, pkg.FieldSix, v.Nested.FieldSix)
}

func TestSpawnNestedTestTypeFromMap(t *testing.T) {
	v, err := Spawn[TestType](map[string]any{
		pkg.FieldOneStr: pkg.FieldOne,
		pkg.FieldTwoStr: pkg.FieldTwo,
		pkg.Nested: map[string]any{
			pkg.FieldFourStr: pkg.FieldFour,
			pkg.FieldFiveStr: pkg.FieldFive,
			pkg.FieldSixStr:  pkg.FieldSix,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, pkg.FieldOne, v.FieldOne)
	assert.Equal(t, pkg.FieldTwo, v.FieldTwo)
	assert.Equal(t, pkg.FieldFour, v.Nested.FieldFour)
	assert.Equal(t, pkg.FieldFive, v.Nested.FieldFive)
	assert.Equal(t, pkg.FieldSix, v.Nested.FieldSix)
}

func TestSpawnUnexportedFieldNestedTestType(t *testing.T) {
	v, err := Spawn[TestType](map[string]any{
		pkg.FieldOneStr:   pkg.FieldOne,
		pkg.FieldTwoStr:   pkg.FieldTwo,
		pkg.FieldThreeStr: pkg.FieldThree,
		pkg.Nested: NestedTestType{
			FieldFour: pkg.FieldFour,
			FieldFive: pkg.FieldFive,
			FieldSix:  pkg.FieldSix,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, pkg.FieldOne, v.FieldOne)
	assert.Equal(t, pkg.FieldTwo, v.FieldTwo)
	assert.Equal(t, pkg.FieldThree, v.fieldThree)
	assert.Equal(t, pkg.FieldFour, v.Nested.FieldFour)
	assert.Equal(t, pkg.FieldFive, v.Nested.FieldFive)
	assert.Equal(t, pkg.FieldSix, v.Nested.FieldSix)
}

func TestSpawnUnexportedFieldNestedTestTypeFromMap(t *testing.T) {
	v, err := Spawn[TestType](map[string]any{
		pkg.FieldOneStr:   pkg.FieldOne,
		pkg.FieldTwoStr:   pkg.FieldTwo,
		pkg.FieldThreeStr: pkg.FieldThree,
		pkg.Nested: map[string]any{
			pkg.FieldFourStr: pkg.FieldFour,
			pkg.FieldFiveStr: pkg.FieldFive,
			pkg.FieldSixStr:  pkg.FieldSix,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, pkg.FieldOne, v.FieldOne)
	assert.Equal(t, pkg.FieldTwo, v.FieldTwo)
	assert.Equal(t, pkg.FieldThree, v.fieldThree)
	assert.Equal(t, pkg.FieldFour, v.Nested.FieldFour)
	assert.Equal(t, pkg.FieldFive, v.Nested.FieldFive)
	assert.Equal(t, pkg.FieldSix, v.Nested.FieldSix)
}

func TestSpawnNestedTestTypeWithWrongMapType(t *testing.T) {
	_, err := Spawn[TestType](map[string]any{
		pkg.FieldOneStr:   pkg.FieldOne,
		pkg.FieldTwoStr:   pkg.FieldTwo,
		pkg.FieldThreeStr: pkg.FieldThree,
		pkg.Nested: map[string]string{
			pkg.FieldFourStr: pkg.FieldFour,
			pkg.FieldFiveStr: "pkg.FieldFive",
			pkg.FieldSixStr:  "true",
		},
	})
	assert.Error(t, err)
	assert.Equal(t, pkg.BadMap(pkg.Nested).Error(), err.Error())
}
