package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func fn() error {
	e1 := errors.New("error")
	return e1
}

func logError(x interface{}) {
	if err, ok := x.(error); ok {
		err1, ok := errors.Cause(err).(stackTracer)
		if !ok {
			panic("oops, err does not implement stackTracer")
		}

		st := err1.StackTrace()
		a := fmt.Sprintf("%+v", st[0])
		a = strings.Replace(a, "\n\t", " at ", 1)
		fmt.Println(a)
	}
}

func main() {
	err := fn()
	logError(err)
}
