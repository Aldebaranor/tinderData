package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

//type ZmqReq struct {
//	g.Meta `path:"/recv" tags:"receive" method:"get" summary:"Scenario"`
//}
//type ZmqRes struct {
//	g.Meta `mime:"recv/html" example:"string"`
//}

type GetCacheReq struct {
	g.Meta `path:"/cache" tags:"cache" method:"get" summary:"cache"`
	SimId  uint32
}
type GetCacheRes struct {
	g.Meta `mime:"cache/html" example:"string"`
}
