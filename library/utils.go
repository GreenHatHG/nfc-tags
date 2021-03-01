package library

import "github.com/gogf/gf/frame/g"

func AsyncPrintStack(err error) {
	g.Log().Async().Stack(false).Errorf("%+s", err)
}
