package controller

import (
	"context"
	v1 "tinderData/api/v1"
	"tinderData/internal/logic/iotDb"
)

var (
	IotDb = cIotDb{}
)

type cIotDb struct{}

func (c *cIotDb) Query(ctx context.Context, req *v1.QueryReq) (res *v1.QueryRes, err error) {
	data := make(map[string]interface{})
	data["sql"] = req.Sql
	data["rowLimit"] = req.RowLimit
	temp := iotDb.Query(data)
	res = &v1.QueryRes{
		//Code: gconv.Int(temp["code"]),
		Msg: temp,
	}
	return
}
