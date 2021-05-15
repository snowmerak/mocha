package script

import (
	"time"

	"github.com/dop251/goja"
)

func RunJS(code string) (string, string) {
	vm := goja.New()
	s := time.Now()
	rs, err := vm.RunString(code)
	e := time.Now()
	r := e.Sub(s).String()
	if err != nil {
		return err.Error(), r
	}
	return rs.String(), r
}
