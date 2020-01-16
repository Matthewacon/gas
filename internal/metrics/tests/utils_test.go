package tests

import (
	"github.com/Matthewacon/gas"
	"testing"

	"github.com/Matthewacon/gas/internal/metrics"
)

func TestAssertPointerPanic(t *testing.T) {
	defer metrics.ExpectPanic(t)
	gas.AssertPointer(0)
}

func TestAssertPointerNoPanic(t *testing.T) {
	defer metrics.ExpectNoPanic(t)
	gas.AssertPointer(&struct{}{})
}

func TestAssertStructPanic(t *testing.T) {
	defer metrics.ExpectPanic(t)
	gas.AssertStruct(0)
}

func TestAssertStructNoPanic(t *testing.T) {
 defer metrics.ExpectNoPanic(t)
 gas.AssertStruct(struct{}{})
}

func TestAssertInterfacePanic(t *testing.T) {
	defer metrics.ExpectPanic(t)
	gas.AssertInterface(0)
}

//func TestAssertInterfaceNoPanic(t *testing.T) {
//	defer metrics.ExpectNoPanic(t)
//	i := func () interface{} {
//		return struct{}{}
//	}()
//	gas.AssertInterface(i)
//}

func TestSlicePanic(t *testing.T) {
	defer metrics.ExpectPanic(t)
	gas.AssertSlice([0]struct{}{})
}

func TestSliceNoPanic(t *testing.T) {
	defer metrics.ExpectNoPanic(t)
	gas.AssertSlice([]struct{}{})
}

func TestArrayPanic(t *testing.T) {
	defer metrics.ExpectPanic(t)
	gas.AssertArray([]struct{}{})
}

func TestArrayNoPanic(t *testing.T) {
	defer metrics.ExpectNoPanic(t)
	gas.AssertArray([0]struct{}{})
}

func TestAssertMapPanic(t *testing.T) {
	defer metrics.ExpectPanic(t)
	gas.AssertMap(0)
}

func TestAssertMapNoPanic(t *testing.T) {
	defer metrics.ExpectNoPanic(t)
	gas.AssertMap(map[int]int{})
}

func TestAssertFuncPanic(t *testing.T) {
	defer metrics.ExpectPanic(t)
	gas.AssertFunc(0)
}

func TestAssertFuncNoPanic(t *testing.T) {
	defer metrics.ExpectNoPanic(t)
	gas.AssertFunc(func(){})
}

func TestRunAssertionClusterError(t *testing.T) {
	defer metrics.ExpectPanic(t)
	err := gas.RunAssertionCluster(func() {
		gas.AssertArray([0]int{})
		gas.AssertFunc(func() {})
		gas.AssertNonNil(nil)
	})
	if err != nil {
		panic(err)
	}
}

func TestRunAssertionClusterNoError(t *testing.T) {
	defer metrics.ExpectNoPanic(t)
	err := gas.RunAssertionCluster(func() {
		gas.AssertArray([0]int{})
		gas.AssertFunc(func() {})
		gas.AssertNonNil(0)
	})
	if err != nil {
		panic(err)
	}
}