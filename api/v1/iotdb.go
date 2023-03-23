package v1

import "github.com/gogf/gf/v2/frame/g"

type QueryReq struct {
	g.Meta   `path:"/query" tags:"query" method:"post" summary:"Query"`
	Sql      string `json:"sql"`
	RowLimit int    `json:"row_limit"`
}

type QueryRes struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"data"`
}
