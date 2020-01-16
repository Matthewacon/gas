package tests

import (
	"reflect"
	"testing"

	"github.com/Matthewacon/gas"
	"github.com/Matthewacon/gas/internal/metrics"
)

func TestAssertNonNilPanic(t *testing.T) {
	defer metrics.ExpectPanic(t)
	gas.AssertNonNil(nil)
}

func TestAssertNonNilNoPanic(t *testing.T) {
	defer metrics.ExpectNoPanic(t)
	gas.AssertNonNil(struct{}{})
}

func TestAssertTypePanic(t *testing.T) {
	defer metrics.ExpectPanic(t)
	s := struct{}{}
	gas.AssertType(reflect.TypeOf(s), nil)
}

func TestAssertTypeNoPanic(t *testing.T) {
	defer metrics.ExpectNoPanic(t)
	i := 0
	gas.AssertType(reflect.TypeOf(i), i)
}

func TestAssertKindPanic(t *testing.T) {
	defer metrics.ExpectPanic(t)
	gas.AssertKind(reflect.Struct, nil)
}

func TestAssertKindNoPanic(t *testing.T) {
	defer metrics.ExpectNoPanic(t)
	s := struct{}{}
	gas.AssertKind(reflect.Struct, s)
}