package script

import (
	"fmt"
	"time"

	"github.com/traefik/yaegi/interp"
)

func RunGo(code string) (string, string) {
	vm := interp.New(interp.Options{})
	s := time.Now()
	v, err := vm.Eval(code)
	e := time.Now()
	r := e.Sub(s).String()
	if err != nil {
		return err.Error(), r
	}
	return fmt.Sprint(v.Interface()), r
}
