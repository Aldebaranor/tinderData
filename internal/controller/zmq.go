package controller

import (
	"context"
	"sync"
	"tinderData/api/v1"
	"tinderData/internal/logic/zmq"
)

var (
	Zmq = cZmq{}
	Wg  = sync.WaitGroup{}
)

type cZmq struct{}

func (c *cZmq) CacheToIotDb(ctx context.Context, req *v1.GetCacheReq) (res *v1.GetCacheRes, err error) {

	go zmq.CleanIotdb(ctx, req.SimId)

	return
}
