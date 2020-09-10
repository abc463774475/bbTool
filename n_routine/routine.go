package n_routine

import (
	"reflect"
	"github.com/abc463774475/bbtool/n_log"
	"runtime/debug"
)

//  i []interface   is the args
func RoutineFun(f interface{}, i ...interface{}) {
	go func(f interface{}, i ...interface{}) {
		defer func() {
			if err := recover() ; err != nil {
				r := reflect.TypeOf(f)
				n_log.Erro("have fun erro  %v   %v",err,r)
				debug.PrintStack()
			}
		}()

		j := reflect.ValueOf(f)

		param := []reflect.Value{}
		for _, v := range i {
			param = append(param, reflect.ValueOf(v))
		}
		j.Call(param)
	}(f, i...)
}

type FStart func() error
type FEnd func() error

func WithPanicResolveFull (fs FStart, fe FEnd, f interface{}, i ...interface{})  {
	if fs != nil {
		fs()
	}

	go func(f interface{}, i ...interface{}) {
		defer func() {
			if err := recover() ; err != nil {
				r := reflect.TypeOf(f)
				n_log.Erro("have fun erro  %v   %v",err,r)
				debug.PrintStack()
			}

			fe ()
		}()

		j := reflect.ValueOf(f)

		param := []reflect.Value{}
		for _, v := range i {
			param = append(param, reflect.ValueOf(v))
		}
		j.Call(param)
	}(f, i...)
}
