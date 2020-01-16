package benchmarks

import (
	"testing"
	
	"github.com/Matthewacon/gas"
	"github.com/Matthewacon/gas/internal/metrics"
)

func BenchmarkAssertPointerNoPanicConstant(b *testing.B) {
	var p *struct{}
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertPointer(p)
	}
}

func BenchmarkAssertPointerPanicConstant(b *testing.B) {
	var p struct{}
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertPointer(p)
	}
}

func BenchmarkAssertPointerNoPanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertPointer(&struct{}{})
	}
}

func BenchmarkAssertPointerPanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertPointer(i)
	}
}

func BenchmarkAssertStructNoPanicConstant(b *testing.B) {
	var s struct{}
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertStruct(s)
	}
}

func BenchmarkAssertStructPanicConstant(b *testing.B) {
	var p *struct{}
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertStruct(p)
	}
}

func BenchmarkAssertStructNoPanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertStruct(struct{}{})
	}
}

func BenchmarkAssertStructPanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertStruct(i)
	}
}

func BenchmarkAssertSliceNoPanicConstant(b *testing.B) {
	var s []int
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertSlice(s)
	}
}

func BenchmarkAssertSlicePanicConstant(b *testing.B) {
	var p int
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertSlice(p)
	}
}

func BenchmarkAssertSliceNoPanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertSlice([]int{})
	}
}

func BenchmarkAssertSlicePanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertSlice(i)
	}
}

func BenchmarkAssertArrayNoPanicConstant(b *testing.B) {
	var s [0]int
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertArray(s)
	}
}

func BenchmarkAssertArrayPanicConstant(b *testing.B) {
	var p int
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertArray(p)
	}
}

func BenchmarkAssertArrayNoPanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertArray([0]int{})
	}
}

func BenchmarkAssertArrayPanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertArray(i)
	}
}

func BenchmarkAssertMapNoPanicConstant(b *testing.B) {
	var s map[int]int
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertMap(s)
	}
}

func BenchmarkAssertMapPanicConstant(b *testing.B) {
	var p int
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertMap(p)
	}
}

func BenchmarkAssertMapNoPanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertMap(map[int]int{})
	}
}

func BenchmarkAssertMapPanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertMap(i)
	}
}

func BenchmarkAssertFuncNoPanicConstant(b *testing.B) {
	var s func()
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertFunc(s)
	}
}

func BenchmarkAssertFuncPanicConstant(b *testing.B) {
	var p int
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertFunc(p)
	}
}

func BenchmarkAssertFuncNoPanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectNoPanic(b)
		gas.AssertFunc(func() {
			//generate distinct functions
			_ = i
		})
	}
}

func BenchmarkAssertFuncPanicChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer metrics.ExpectPanic(b)
		gas.AssertFunc(i)
	}
}


func BenchmarkRunAssertionClusterNoErrorConstant(b *testing.B) {
	cluster := func() {}
	for i := 0; i < b.N; i++ {
		if err := gas.RunAssertionCluster(cluster); err != nil {
			b.Errorf("Cluster returned unexpected error: %s\n", err.Error())
			return
		}
	}
}

func BenchmarkRunAssertionClusterErrorConstant(b *testing.B) {
	cluster := func() {
		panic(&struct{}{})
	}
	for i := 0; i < b.N; i++ {
		if err := gas.RunAssertionCluster(cluster); err == nil {
			b.Errorf("Cluster did not return an error!\n")
		}
	}
}

func BenchmarkRunAssertionClusterNoErrorChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := gas.RunAssertionCluster(func() {
			//generate distinct functions
			_ = i
		})
		if err != nil {
			b.Errorf("Cluster returned unexpected error: %s\n", err.Error())
			return
		}
	}
}

func BenchmarkRunAssertionClusterErrorChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := gas.RunAssertionCluster(func () {
			panic(string(i))
		})
		if err == nil {
			b.Errorf("Cluster did not return an error!\n")
			return
		}
	}
}
