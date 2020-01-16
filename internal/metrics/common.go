package metrics

//both testing.T and testing.B implement this interface
type Common interface {
	Errorf(format string, args ...interface{})
}

func ExpectPanic(c Common) {
	if recover() != nil {
		return
	}
	c.Errorf("Assertion did not panic as expected!\n")
}

func ExpectNoPanic(c Common) {
	r := recover()
	if r == nil {
		return
	}
	c.Errorf(
		"Assertion paniced unexpectedly: \n%v\n",
		r,
	)
}
