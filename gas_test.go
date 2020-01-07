package gas

import (
	"reflect"
	"testing"
)

func expectPanic(t *testing.T) {
	if r := recover(); r != nil {
		return
	}
	t.Errorf("Assertion did not panic as expected!\n")
}

func expectNoPanic(t *testing.T) {
	r := recover()
	if r == nil {
		return
	}
	t.Errorf("Assertion paniced unexpectedly: \n%v\n", r)
}

func TestAssertNonNilPanic(t *testing.T) {
	defer expectPanic(t)
	AssertNonNil(nil)
}

func TestAssertNonNilNoPanic(t *testing.T) {
	defer expectNoPanic(t)
	AssertNonNil(struct{}{})
}

func TestAssertTypePanic(t *testing.T) {
	defer expectPanic(t)
	s := struct{}{}
	AssertType(reflect.TypeOf(s), nil)
}

func TestAssertTypeNoPanic(t *testing.T) {
	defer expectNoPanic(t)
	i := 0
	AssertType(reflect.TypeOf(i), i)
}

func TestAssertKindPanic(t *testing.T) {
	defer expectPanic(t)
	AssertKind(reflect.Struct, nil)
}

func TestAssertKindNoPanic(t *testing.T) {
	defer expectNoPanic(t)
	s := struct{}{}
	AssertKind(reflect.Struct, s)
}

func BenchmarkAssertNonNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AssertNonNil(i)
	}
}

func BenchmarkAssertTypeConstant(b *testing.B) {
	s := struct{}{}
	_type := reflect.TypeOf(s)
	BuildTypeAssertion(_type)
	for i := 0; i < b.N; i++ {
		AssertType(_type, s)
	}
}

func BenchmarkAssertTypeChange(b *testing.B) {
	_type := reflect.TypeOf(struct{}{})
	BuildTypeAssertion(_type)
	for i := 0; i < b.N; i++ {
		AssertType(_type, struct{}{})
	}
}

func BenchmarkAssertKindConstant(b *testing.B) {
	ii := 0
	BuildKindAssertion(reflect.Int)
	for i := 0; i < b.N; i++ {
		AssertKind(reflect.Int, ii)
	}
}

func BenchmarkAssertKindChange(b *testing.B) {
	BuildKindAssertion(reflect.Int)
	for i := 0; i < b.N; i++ {
		AssertKind(reflect.Int, i)
	}
}