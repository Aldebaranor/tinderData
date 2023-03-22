package main

import (
	"context"
	"tinderData/internal/logic/iotDb"
	"tinderData/internal/logic/zmq"
	_ "tinderData/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"tinderData/internal/cmd"
)

func main() {
	//建立数据库连接
	var ctx context.Context
	iotDb.ConnectIotDb()

	go getZmq(ctx)
	go getDetect(ctx)

	cmd.Main.Run(gctx.New())
}

func getZmq(ctx context.Context) {
	for {
		zmq.RecvMsg(ctx)
	}
}

func getDetect(ctx context.Context) {
	for {
		zmq.RecvDetect(ctx)
	}
}
