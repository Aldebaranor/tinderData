package softbus

import (
	"github.com/gogf/gf/os/glog"
	zmq "github.com/pebbe/zmq4"
)

type context struct {
	ctx *zmq.Context
}

var contextDefault context

func init() {
	ctx, err := zmq.NewContext()
	if err != nil {
		glog.Error(nil, "Failed to create zmq context")
		glog.Error(err)
	}

	// 设置sockets最大数量
	//err = ctx.SetMaxSockets(10000)
	//if err != nil {
	//	glog.Error("Failed to set max sockets: ", err)
	//	panic(err)
	//}

	contextDefault.ctx = ctx
}

func Ctx() *zmq.Context {
	return contextDefault.ctx
}
