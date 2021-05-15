package problem

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dop251/goja"
	"github.com/traefik/yaegi/interp"
)

const solution = "solution"

func evalGo(cases []Case, code string) (time.Duration, error) {
	t := time.Duration(0)

	for _, c := range cases {
		vm := interp.New(interp.Options{})
		s := time.Now()
		v, err := vm.Eval(code + fmt.Sprintf("\n%s(%s)", solution, strings.Join(c.Input, ", ")))
		e := time.Now()
		if err != nil {
			return 0, err
		}
		if fmt.Sprint(v.Interface()) != strings.Join(c.Output, ", ") {
			return 0, errors.New("failed")
		}
		t = time.Duration((int64(t) + int64(e.Sub(s))) / 2)
	}

	return t, nil
}

func evalJS(cases []Case, code string) (time.Duration, error) {
	t := time.Duration(0)

	for _, c := range cases {
		vm := goja.New()
		s := time.Now()
		v, err := vm.RunString(code + fmt.Sprintf("\n%s(%s)", solution, strings.Join(c.Input, ", ")))
		e := time.Now()
		if err != nil {
			return 0, err
		}
		if fmt.Sprint(v.String()) != strings.Join(c.Output, ", ") {
			return 0, errors.New("failed")
		}
		t = time.Duration((int64(t) + int64(e.Sub(s))) / 2)
	}

	return t, nil
}
