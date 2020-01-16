package benchmarks

import (
	"reflect"
	"testing"

	"github.com/Matthewacon/gas"
)

func BenchmarkAssertNonNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gas.AssertNonNil(i)
	}
}

func BenchmarkAssertTypeConstant(b *testing.B) {
	s := struct{}{}
	_type := reflect.TypeOf(s)
	for i := 0; i < b.N; i++ {
		gas.AssertType(_type, s)
	}
}

func BenchmarkAssertTypeChange(b *testing.B) {
	_type := reflect.TypeOf(struct{}{})
	for i := 0; i < b.N; i++ {
		gas.AssertType(_type, struct{}{})
	}
}

func BenchmarkAssertKindConstant(b *testing.B) {
	ii := 0
	for i := 0; i < b.N; i++ {
		gas.AssertKind(reflect.Int, ii)
	}
}

func BenchmarkAssertKindChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gas.AssertKind(reflect.Int, i)
	}
}
