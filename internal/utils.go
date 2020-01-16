package internal

import (
	"fmt"
	"reflect"
)

type AssertionError string
//error
func (e AssertionError) Error() string { return string(e) }

func Format(format ...interface{}) string {
	if format != nil {
		if len(format) >= 1 {
			return fmt.Sprintf(format[0].(string), format[1:]...)
		} else if len(format) == 1 {
			return fmt.Sprintf(format[0].(string))
		}
	}
	return ""
}

func AssertPanic(assertion interface{}, message string) {
	typ := reflect.TypeOf(assertion).Name()
	pMessage := []interface{}{
		"%s Assertion failed!\n",
		typ,
	}
	if len(message) > 0 {
		pMessage = []interface{}{
			"%s Assertion failed: %s\n",
			typ,
			message,
		}
	}
	panic(AssertionError(Format(pMessage...)))
}