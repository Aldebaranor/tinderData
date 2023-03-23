package zmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
	"strings"
	"tinderData/internal/logic/gcache"
	"tinderData/internal/logic/iotDb"
	"tinderData/internal/logic/softbus"
	"tinderData/internal/model/entity"
	"tinderData/internal/model/message"
	proto_message "tinderData/internal/model/proto"
)

type Situation struct {
	simId            uint32      // 推演任务id
	ctlSubscriber    *zmq.Socket // 订阅controller发布的消息
	detectSubscriber *zmq.Socket //订阅dectect信息
	msgOffset        int         // zmq消息过滤标识长度
}

var TotalNum = make(map[uint32]int)

func RecvMsg(ctx context.Context) {
	s := &Situation{}

	subscriberUrl, _ := g.Cfg().Get(ctx, "zmq.server")
	serverUrl := subscriberUrl.String()

	s.ctlSubscriber, _ = softbus.Ctx().NewSocket(zmq.SUB)
	err := s.ctlSubscriber.Connect(serverUrl)
	defer s.ctlSubscriber.Close()

	if err != nil {
		glog.Printf("zmq controller订阅地址出错", err)
		return
	}
	glog.Println("zmq controller消息订阅: ", serverUrl)
	err = s.ctlSubscriber.SetSubscribe("")
	if err != nil {
		return
	}

	for {
		if resp, err := s.ctlSubscriber.RecvBytes(0); err == nil {
			//根据msgOffset过滤消息标识
			lenInt := len(resp)
			//topicMaxInt := gconv.String(resp[:11])
			topicMaxInt := gconv.String(resp[:lenInt-1])
			//下划线分割simId和message编号
			simIdIndex := strings.Index(topicMaxInt, "_")
			topicSimId := gconv.String(resp[:simIdIndex])
			s.simId = gconv.Uint32(topicSimId)
			s.msgOffset = len(topicSimId) + 1

			topicStr := gconv.String(resp[s.msgOffset : s.msgOffset+3])
			topicId := gconv.Uint32(topicStr)
			resp = resp[s.msgOffset+3:]
			if topicId == message.MESSAGE_TASK_NEXT_STEP_SYNC {
				Msg := &proto_message.NextStepSync{}
				err := proto.Unmarshal(resp, Msg)
				if err != nil {
					glog.Println("解析protobuf出错")
				}
				//glog.Println("controller Msg:", Msg)
				//态势数据存入cache和levelDb
				gcache.SaveAsCacheAndLevelDb(ctx, topicSimId, Msg)
			}
		}
	}
}

func RecvDetect(ctx context.Context) {

	s := &Situation{}

	getDetectUrl, _ := g.Cfg().Get(ctx, "zmq.detect")
	detectUrl := getDetectUrl.String()
	s.detectSubscriber, _ = softbus.Ctx().NewSocket(zmq.SUB)
	detectErr := s.detectSubscriber.Connect(detectUrl)
	defer s.detectSubscriber.Close()
	if detectErr != nil {
		glog.Printf("zmq detect订阅地址出错", detectErr)
		return
	}
	glog.Println("zmq detect消息订阅: ", detectUrl)
	detectErr = s.detectSubscriber.SetSubscribe("")
	if detectErr != nil {
		return
	}

	for {
		topic, detectErr := s.detectSubscriber.RecvBytes(0)
		if detectErr != nil {
			glog.Println("detect解析topic出错")
			continue
		}

		msg, detectErr := s.detectSubscriber.RecvBytes(0)
		if detectErr != nil {
			glog.Println("detect接收报文消息失败")
			continue
		}
		lenInt := len(topic)
		//根据msgOffset过滤消息标识
		//topicMaxInt := gconv.String(topic[:11])
		topicMaxInt := gconv.String(topic[:lenInt-1])

		//下划线分割simId和message编号
		simIdIndex := strings.Index(topicMaxInt, "_")
		topicSimId := gconv.String(topic[:simIdIndex])
		s.simId = gconv.Uint32(topicSimId)
		s.msgOffset = len(topicSimId) + 1

		topicStr := gconv.String(topic[s.msgOffset:])
		topicId := gconv.Uint32(topicStr)
		topic = topic[:s.msgOffset+1]

		//MESSAGE_FUSION_DETECT
		if topicId == message.MESSAGE_FUSION_DETECT {
			fd := &proto_message.FusionDetected{}
			err := proto.Unmarshal(msg, fd)
			if err != nil {
				glog.Println("解析protobuf出错")
			}
			//glog.Println("detect Msg:", fd)
			//融合探测数据存入cache和levelDb
			if fd.Data != nil {
				gcache.SaveDetect(ctx, topicSimId, fd)
			}
		}
	}

}

// cache存入IotDb
func CleanIotdb(ctx context.Context, id uint32) {
	//存IotDb之前清空同名数据
	//TODO:修复清除数据的bug
	//iotDb.CleanCache(id)
	SaveIotDb(ctx, id)
}

func SaveIotDb(ctx context.Context, sId uint32) {

	//timer := time.NewTimer(time.Second * 2)

	////设置空闲时间，一段时间没有数据就关闭持久化进程
	//stopT := time.NewTimer(time.Second * 60)
	//flag := false

	for {
		//select {
		//case <-timer.C:
		//取出所有key
		simList, _ := gcache.MsgCache.Keys(ctx)
		if len(simList) != 0 {
			//筛选出所有Id为sId的key
			var keyList []string
			for _, sim := range simList {
				if strings.Contains(gconv.String(sim), gconv.String(sId)) {
					keyList = append(keyList, gconv.String(sim))
				}
			}

			if len(keyList) != 0 {
				fmt.Println(gconv.String(sId) + " has " + gconv.String(len(keyList)))
				//flag = true
				var id uint32
				for _, key := range keyList {
					//取出来msg是gvar.var类型,消费掉
					msgVar, _ := gcache.MsgCache.Get(ctx, key)
					if msgVar.Val() != nil {
						id = gconv.Uint32(strings.Split(gconv.String(key), "@")[0])
						//if _, ok := TotalNum[id]; ok {
						//	TotalNum[id]++
						//} else {
						//	TotalNum[id] = 0
						//}
						gcache.MsgCache.Remove(ctx, key)
						//类型转换,gvar类型的msg转换为结构体数组进行遍历存入IotDb
						msgBytes, _ := json.Marshal(msgVar)
						var msg entity.IotDbForce
						json.Unmarshal(msgBytes, &msg)
						simId := gconv.String(id)
						simTime := gconv.Int64(msg.SimTime)
						dataTime := simTime * 1e6
						//fmt.Println("---start insert table---"+gconv.String(simId)+" simTime "+gconv.String(simTime)+" forceSize " + gconv.String(len(msg.Forces)))
						iotDb.InsertTabel(msg.Forces, simId, dataTime, simTime)
						//fmt.Println("---insert over---"+gconv.String(simId)+" simTime "+gconv.String(simTime)+" forceSize ")
						//for _, force := range msg.Forces {
						//	//iotDb.Wg.Add(1)
						//	dataTime++
						//	iotDb.SaveToIotDb(force, simId, dataTime, simTime)
						//	//go iotDb.SaveToIotDb(force, simId, simTime)
						//}
						glog.Println("[" + gconv.String(id) + "] " + "-heartbeat->" + gconv.String(simTime))

					}
				}
				//glog.Println("[" + gconv.String(id) + "] " + "-total->" + gconv.String(TotalNum[id]))
				//iotDb.Wg.Wait()
			}
		}
		//设置定时for循环5s一次
		//timer.Reset(time.Second * 5)
		//fmt.Println(sId)
		//case <-stopT.C:
		//	//60秒内都没有数据，flag一直不变都是false
		//	if flag {
		//		//60s内有数据就再续60s
		//		stopT.Reset(time.Second * 60)
		//	} else {
		//		//没有数据就return
		//		glog.Println("--------------" + gconv.String(sId) + "--------超过1分钟没有数据，IotDb持久化结束")
		//		return
		//	}
		//}
	}

}
