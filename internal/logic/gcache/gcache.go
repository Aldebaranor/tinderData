package gcache

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"time"
	"tinderData/internal/model/entity"
	proto_message "tinderData/internal/model/proto"
)

var (
	MsgCache = gcache.New()
	IdCache  = gcache.New()
)

func SaveAsCacheAndLevelDb(ctx context.Context, msgSimId string, msg *proto_message.NextStepSync) {

	simTime := gconv.String(msg.SimTime)
	fKey := msgSimId + "@" + simTime

	//这里是维护一个所有想定Id的缓存
	//IdCache存在simId的key？
	containId, _ := IdCache.Contains(ctx, msgSimId)
	if !containId {
		//simId存入IdCache，设置过期时间60s
		IdCache.Set(ctx, msgSimId, msgSimId, time.Second*60)
	} else {
		//simId作为Key存入IdCache，60s超期
		IdCache.Set(ctx, msgSimId, msgSimId, time.Second*60)
	}

	//这里是对zmq可能把数据分两个包发送进行合并处理
	contain, _ := MsgCache.Contains(ctx, fKey)
	if contain {
		//取出数据
		halfVar, _ := MsgCache.Get(ctx, fKey)
		msgBytes, _ := json.Marshal(halfVar)
		var forces entity.IotDbForce
		json.Unmarshal(msgBytes, &forces)
		//msg.Data转Force
		//合并forces数据存msgMache
		forces = addForce(msg, forces)
		//fKey存入MsgCache，60s超期，用于进行持久化
		MsgCache.Set(ctx, fKey, forces, time.Second*60)
	} else {

		forces := msgToCache(msg)
		//MsgCache存入第一批数据
		MsgCache.Set(ctx, fKey, forces, time.Second*60)
	}

	//存入levelDb，文件名为simId
	//levelDb暂时没有合并detect
	//LevelDbSession, err := leveldb.OpenFile("./"+msgSimId+".db", nil)
	//if err != nil {
	//	log.Printf("%+v", err)
	//}
	//defer LevelDbSession.Close()
	//
	////序列化
	//jKey, _ := json.Marshal(simTime)
	//jValue, _ := json.Marshal(msg)
	////存到levelDb
	//LevelDbSession.Put([]byte(jKey), []byte(jValue), nil)

}

func SaveDetect(ctx context.Context, msgSimId string, msg *proto_message.FusionDetected) {

	simTime := gconv.String(msg.SimTime)
	fKey := msgSimId + simTime
	//cache存在simId_Time的key？
	contain, _ := MsgCache.Contains(ctx, fKey)
	if contain {
		//取出数据
		halfVar, _ := MsgCache.Get(ctx, fKey)
		msgBytes, _ := json.Marshal(halfVar)
		var forces entity.IotDbForce
		json.Unmarshal(msgBytes, &forces)

		//msg.Data转Force
		//合并存cache
		forces = addDetect(msg.Data, forces)

		//simId_Time作为Key存入cache，60s超期
		MsgCache.Set(ctx, fKey, forces, time.Second*60)
	} else {
		//simId存入cache，设置过期时间60s
		MsgCache.Set(ctx, msgSimId, nil, time.Second*60)
		forces := detectToCache(msg)
		MsgCache.Set(ctx, fKey, forces, time.Second*60)
	}

	//检查levelDb覆盖问题？
	//存入levelDb，文件名为simId
	LevelDbSession, err := leveldb.OpenFile("./"+msgSimId+"detect.db", nil)
	if err != nil {
		log.Printf("%+v", err)
	}
	defer LevelDbSession.Close()

	//序列化
	jKey, _ := json.Marshal(simTime)
	jValue, _ := json.Marshal(msg)
	//存到levelDb
	LevelDbSession.Put([]byte(jKey), []byte(jValue), nil)
}

func detectToCache(data *proto_message.FusionDetected) *entity.IotDbForce {
	var forces []*entity.IotDbForcePosture
	for _, datum := range data.Data {
		force := &entity.IotDbForcePosture{
			ForceId:  datum.Id,
			Lon2:     datum.Lon,
			Lat2:     datum.Lat,
			Alt2:     datum.Alt,
			Heading2: datum.Heading,
			Pitch2:   datum.Pitch,
			Speed2:   datum.Speed,
		}
		forces = append(forces, force)
	}
	res := &entity.IotDbForce{
		SimTime: data.SimTime,
		Forces:  forces,
	}

	return res
}

func msgToCache(data *proto_message.NextStepSync) *entity.IotDbForce {
	var forces []*entity.IotDbForcePosture
	for _, datum := range data.Forces {
		force := &entity.IotDbForcePosture{
			ForceId: datum.ForceId,
			Lon:     datum.Lon,
			Lat:     datum.Lat,
			Alt:     datum.Alt,
			Heading: datum.Heading,
			Pitch:   datum.Pitch,
			Speed:   datum.Speed,
		}
		forces = append(forces, force)
	}
	res := &entity.IotDbForce{
		Stage:     data.Stage,
		SimTime:   data.SimTime,
		StepRatio: data.StepRatio,
		Forces:    forces,
	}

	return res
}

func addForce(data *proto_message.NextStepSync, msg entity.IotDbForce) entity.IotDbForce {
	for _, datum := range data.Forces {
		for _, force := range msg.Forces {
			if datum.ForceId == force.ForceId {
				force.Lon = datum.Lon
				force.Lat = datum.Lat
				force.Alt = datum.Alt
				force.Heading = datum.Heading
				force.Pitch = datum.Pitch
				force.Speed = datum.Speed
			}
		}
	}
	msg.Stage = data.Stage
	msg.StepRatio = data.StepRatio
	msg.SimTime = data.SimTime
	return msg
}

func addDetect(data []*proto_message.DetectedData, msg entity.IotDbForce) entity.IotDbForce {

	for _, datum := range data {
		for _, force := range msg.Forces {
			if datum.Id == force.ForceId {
				force.Lon2 = datum.Lon
				force.Lat2 = datum.Lat
				force.Alt2 = datum.Alt
				force.Heading2 = datum.Heading
				force.Pitch2 = datum.Pitch
				force.Speed2 = datum.Speed
			}
		}
	}
	return msg
}
