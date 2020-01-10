package gas

import (
	"fmt"
	"reflect"
)

func format(format ...interface{}) string {
	if format != nil {
		if len(format) >= 1 {
			return fmt.Sprintf(format[0].(string), format[1:]...)
		} else if len(format) == 1 {
			return fmt.Sprintf(format[0].(string))
		}
	}
	return ""
}

type Assertion func (interface{}, ...interface{}) ()

//store built assertions for cheap reuse
var arbitraryAssertions map[reflect.Type]Assertion
var kindAssertions map[reflect.Kind]Assertion

func AssertNonNil(i interface{}, formatString ...interface{}) {
	if i == nil {
		if len(formatString) > 0 {
			panic(fmt.Errorf(format(formatString...)))
		} else {
			panic(fmt.Errorf("Nil assertion failed!\n"))
		}
	}
}

func BuildTypeAssertion(_type reflect.Type) Assertion {
	var _f Assertion
	if f, ok := arbitraryAssertions[_type]; !ok {
		_f = func (i interface{}, fArgs ...interface{}) {
			iType := reflect.TypeOf(i)
			if iType != _type {
				var err error
				if str := format(fArgs...); len(str) > 0 {
					err = fmt.Errorf(str)
				} else {
					err = fmt.Errorf("Assertion failed: Type '%s' is not of type '%s'!\n", _type.String(), iType)
				}
				panic(err)
			}
		}
	} else {
		_f = f
	}
	return _f
}

func AssertType(_type reflect.Type, test interface{}, format ...interface{}) {
	BuildTypeAssertion(_type)(test, format...)
}

func BuildKindAssertion(kind reflect.Kind) Assertion {
	var _f Assertion
	if f, ok := kindAssertions[kind]; !ok {
		_f = func (i interface{}, fArgs ...interface{}) {
			AssertNonNil(i)
			iKind := reflect.TypeOf(i).Kind()
			if iKind != kind {
				var err error
				if str := format(fArgs...); len(str) > 0 {
					err = fmt.Errorf(str)
				} else {
					err = fmt.Errorf("Assertion failed: Type '%s' is not of type '%s'!\n", iKind.String(), kind.String())
				}
				panic(err)
			}
		}
	} else {
		_f = f
	}
	return _f
}

func AssertKind(kind reflect.Kind, i interface{}, format ...interface{}) {
	BuildKindAssertion(kind)(i, format...)
}

func AssertPointer(i interface{}, format ...interface{}) {
	BuildKindAssertion(reflect.Ptr)(i, format...)
}

func AssertStruct(i interface{}, format ...interface{}) {
	BuildKindAssertion(reflect.Struct)(i, format...)
}

func AssertInterface(i interface{}, format ...interface{}) {
	BuildKindAssertion(reflect.Interface)(i, format...)
}

func AssertSlice(i interface{}, format ...interface{}) {
	BuildKindAssertion(reflect.Slice)(i, format...)
}

func AssertArray(i interface{}, format ...interface{}) {
	BuildKindAssertion(reflect.Array)(i, format...)
}

func AssertMap(i interface{}, format ...interface{}) {
	BuildKindAssertion(reflect.Map)(i, format...)
}

func AssertFunc(i interface{}, format ...interface{}) {
	BuildKindAssertion(reflect.Func)(i, format...)
}