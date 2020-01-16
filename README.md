# gas
A simple Go runtime assertion library.

## Getting started
Add `github.com/Matthewacon/gas v0.0.3` to the require section in your `go.mod`.  

### Prebuilt assertions
gas provides a set of prebuilt assertion functions:
```go
AssertNonNil
AssertPointer
AssertStruct
AssertInterface
AssertSlice
AssertArray
AssertMap
AssertFunc
```

All prebuilt assertions conform to the function prototype specified by `gas.Assertion`:
```go
type Assertion func(i interface{}, format ...interface{})
```
where `i` is the value that you want to test and `format...` are the formatting directive and subsequent
formatting values, should you want to provide a message for an assertion failure.

#### Example
```go
package main

import "github.com/Matthewacon/gas"

func main() {
 gas.AssertNonNil(nil, "Uh oh, we didn't want that to be nil!\n")
}
```

### Custom assertions
Additionally, gas provides two custom assertion functions for building custom type and kind assertions:
```go
package main

import "reflect"

func AssertType(k reflect.Kind, i interface{}, format ...interface{})
func AssertKind(t reflect.Type, i interface{}, format ...interface{})
```

### Assertion clusters
You may have multiple assertions that you want to run, and you may want to handle a failure. Assertion 
clusters allow you to group all of your assertions into a handler and produce an error upon assertion
failure, rather than panic.

```go
package main

import "github.com/Matthewacon/gas"

func main() {
 err := gas.RunAssertionCluster(func() {
  gas.AssertNonNil(3.14)
  gas.AssertPointer(&struct{}{})
  gas.AssertFunc("am I a function?", "no I am not!\n")
 })
 if err != nil {
  //handle failure
 }
}
```
 
## Running the tests and benchmarks
Test:
```sh
go test ./...
```
Bench:
```sh
go test -bench=. ./...
```

## How expensive is it?
Assertions are cached when built so reuse is cheap. See the benchmarks:
```
BenchmarkAssertNonNil-8                         	545119610	        2.20 ns/op
BenchmarkAssertTypeConstant-8                   	25229216	       46.2 ns/op
BenchmarkAssertTypeChange-8                     	25249456	       46.3 ns/op
BenchmarkAssertKindConstant-8                   	40084676	       29.6 ns/op
BenchmarkAssertKindChange-8                     	28726497	       41.9 ns/op
BenchmarkAssertPointerNoPanicConstant-8         	4603818	      287 ns/op
BenchmarkAssertPointerPanicConstant-8           	1000000000	        0.000003 ns/op
BenchmarkAssertPointerNoPanicChange-8           	3768321	      292 ns/op
BenchmarkAssertPointerPanicChange-8             	1000000000	        0.000006 ns/op
BenchmarkAssertStructNoPanicConstant-8          	5278958	      247 ns/op
BenchmarkAssertStructPanicConstant-8            	1000000000	        0.000006 ns/op
BenchmarkAssertStructNoPanicChange-8            	3755145	      293 ns/op
BenchmarkAssertStructPanicChange-8              	1000000000	        0.000006 ns/op
BenchmarkAssertSliceNoPanicConstant-8           	3755778	      295 ns/op
BenchmarkAssertSlicePanicConstant-8             	1000000000	        0.000005 ns/op
BenchmarkAssertSliceNoPanicChange-8             	3253950	      391 ns/op
BenchmarkAssertSlicePanicChange-8               	1000000000	        0.000007 ns/op
BenchmarkAssertArrayNoPanicConstant-8           	3740067	      293 ns/op
BenchmarkAssertArrayPanicConstant-8             	1000000000	        0.000006 ns/op
BenchmarkAssertArrayNoPanicChange-8             	5167531	      248 ns/op
BenchmarkAssertArrayPanicChange-8               	1000000000	        0.000006 ns/op
BenchmarkAssertMapNoPanicConstant-8             	3774405	      286 ns/op
BenchmarkAssertMapPanicConstant-8               	1000000000	        0.000004 ns/op
BenchmarkAssertMapNoPanicChange-8               	3108926	      439 ns/op
BenchmarkAssertMapPanicChange-8                 	1000000000	        0.000004 ns/op
BenchmarkAssertFuncNoPanicConstant-8            	3716983	      294 ns/op
BenchmarkAssertFuncPanicConstant-8              	1000000000	        0.000004 ns/op
BenchmarkAssertFuncNoPanicChange-8              	3411025	      311 ns/op
BenchmarkAssertFuncPanicChange-8                	1000000000	        0.000003 ns/op
BenchmarkRunAssertionClusterNoErrorConstant-8   	28611800	       42.1 ns/op
BenchmarkRunAssertionClusterErrorConstant-8     	4980104	      241 ns/op
BenchmarkRunAssertionClusterNoErrorChange-8     	29019114	       41.3 ns/op
BenchmarkRunAssertionClusterErrorChange-8       	5455513	      218 ns/op
```

## License
This project is licensed under the [M.I.T. License](https://github.com/Matthewacon/gas/blob/master/LICENSE).
