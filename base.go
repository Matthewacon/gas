package gas

import (
	"fmt"
	"reflect"

	"github.com/Matthewacon/gas/internal"
)

type Assertion func (interface{}, ...interface{}) ()

//store built assertions for cheap reuse
var arbitraryAssertions map[reflect.Type]Assertion
var kindAssertions map[reflect.Kind]Assertion

func AssertNonNil(i interface{}, formatString ...interface{}) {
	if i == nil {
		err := ""
		if len(formatString) > 0 {
			err = internal.Format(formatString...)
		}
		internal.AssertPanic(AssertNonNil, err)
	}
}

func buildTypeAssertion(_type reflect.Type) Assertion {
	var _f Assertion
	if f, ok := arbitraryAssertions[_type]; !ok {
		_f = func (i interface{}, fArgs ...interface{}) {
			iType := reflect.TypeOf(i)
			if iType != _type {
				var err string
				if len(fArgs) > 0 {
					err = internal.Format(fArgs...)
				} else {
					err = fmt.Sprintf(
						"Type '%s' is not of type '%s'!\n",
						_type.String(),
						iType,
					)
				}
				internal.AssertPanic(AssertType, err)
			}
		}
	} else {
		_f = f
	}
	return _f
}

func AssertType(_type reflect.Type, test interface{}, format ...interface{}) {
	buildTypeAssertion(_type)(test, format...)
}

func buildKindAssertion(kind reflect.Kind) Assertion {
	var _f Assertion
	if f, ok := kindAssertions[kind]; !ok {
		_f = func (i interface{}, fArgs ...interface{}) {
			AssertNonNil(i)
			iKind := reflect.TypeOf(i).Kind()
			if iKind != kind {
				var err string
				if len(fArgs) > 0 {
					err = internal.Format(fArgs...)
				} else {
					err = fmt.Sprintf(
						"Type '%s' is not of type '%s'!\n",
						iKind.String(),
						kind.String(),
					)
				}
				internal.AssertPanic(AssertKind, err)
			}
		}
	} else {
		_f = f
	}
	return _f
}

func AssertKind(kind reflect.Kind, i interface{}, format ...interface{}) {
	buildKindAssertion(kind)(i, format...)
}