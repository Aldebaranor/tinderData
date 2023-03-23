package iotDb

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"io/ioutil"
	"net/http"
)

func Query(data map[string]interface{}) map[string]interface{} {
	host, _ := g.Cfg().Get(nil, "iotDb.host")
	restPort, _ := g.Cfg().Get(nil, "iotDb.restPort")
	url := gconv.String(host) + ":" + gconv.String(restPort) + "/rest/v1/query"
	username, _ := g.Cfg().Get(nil, "iotDb.user")
	password, _ := g.Cfg().Get(nil, "iotDb.password")
	//restAPI鉴权
	//'Authorization': 'Basic ' + base64.encode(username + ':' + password)
	c := g.Client()
	authData := gconv.String(username) + ":" + gconv.String(password)
	authDataByte := []byte(authData)
	c.SetHeader("Authorization", "Basic "+gconv.String(base64.StdEncoding.EncodeToString(authDataByte)))
	c.SetContentType("application/json")
	//c.SetHeader("Authorization","Basic cm9vdDpyb2901")
	r, err := c.Post(nil, url, data)
	if err != nil {
		glog.Error(err)
	}
	defer r.Close()
	returnMap, _ := ParseResponse(r.Response)
	return returnMap
}

// 解析响应体json
func ParseResponse(response *http.Response) (map[string]interface{}, error) {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	return result, err
}
