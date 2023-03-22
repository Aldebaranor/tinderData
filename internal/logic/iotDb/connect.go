package iotDb

import (
	"github.com/apache/iotdb-client-go/client"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	host     string
	port     string
	user     string
	password string
)

var IotDbSessionPool client.SessionPool

func ConnectIotDb() {
	//连接IotDb

	host, _ := g.Cfg().Get(nil, "iotDb.host")
	port, _ := g.Cfg().Get(nil, "iotDb.port")
	user, _ := g.Cfg().Get(nil, "iotDb.user")
	password, _ := g.Cfg().Get(nil, "iotDb.password")

	config := &client.PoolConfig{
		Host:     host.String(),
		Port:     port.String(),
		UserName: user.String(),
		Password: password.String(),
	}
	IotDbSessionPool = client.NewSessionPool(config, 10, 60000, 60000, false)

	//defer IotDbSessionPool.Close()

}
