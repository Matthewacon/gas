package gas

import (
	"fmt"
	"reflect"
)

func AssertPointer(i interface{}, format ...interface{}) {
	buildKindAssertion(reflect.Ptr)(i, format...)
}

func AssertStruct(i interface{}, format ...interface{}) {
	buildKindAssertion(reflect.Struct)(i, format...)
}

func AssertInterface(i interface{}, format ...interface{}) {
	buildKindAssertion(reflect.Interface)(i, format...)
}

func AssertSlice(i interface{}, format ...interface{}) {
	buildKindAssertion(reflect.Slice)(i, format...)
}

func AssertArray(i interface{}, format ...interface{}) {
	buildKindAssertion(reflect.Array)(i, format...)
}

func AssertMap(i interface{}, format ...interface{}) {
	buildKindAssertion(reflect.Map)(i, format...)
}

func AssertFunc(i interface{}, format ...interface{}) {
	buildKindAssertion(reflect.Func)(i, format...)
}

type AssertionCluster func()

func RunAssertionCluster(cluster AssertionCluster) error {
	var err error = nil
	func () {
		defer func() {
			//TODO once go-away is done replace this recover
			if r := recover(); r != nil {
				err = fmt.Errorf("%v", r)
			}
		}()
		cluster()
	}()
	return err
}