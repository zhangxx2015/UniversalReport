package Exception

import (
	"errors"
	"fmt"
)

func Throw(format string, a ...interface{}) error {
	if a == nil {
		panic(errors.New(format))
	}
	panic(errors.New(fmt.Sprintf(format, a)))
}

// func (c *BaseController) Return() error {
// 	c.JsonData("")
// 	panic(errors.New(""))
// }

type CatchHandler func(error)

func Catch(fun func(error)) CatchHandler {
	return fun
}

func Try(fun func(), handler CatchHandler) {
	defer func() {
		if e := recover(); e != nil {
			var err error = e.(error)
			handler(err)
		}
	}()
	fun()
}

//func main() {
//	Try(func() {
//		Throw("err:foo")
//		print("next")
//	}, func(err error) {
//		println(err.Error())
//		println("catch")
//	})
//	println("done")
//}
