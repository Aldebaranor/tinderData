package iotDb

import (
	"github.com/apache/iotdb-client-go/client"
	"github.com/apache/iotdb-client-go/rpc"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"log"
	"tinderData/internal/model/entity"
)

//var Wg = sync.WaitGroup{}

func SaveToIotDb(force *entity.IotDbForcePosture, simId string, dataTime int64, simTime int64) {
	ts := dataTime
	var (
		deviceId    = "root.tinder." + simId
		measurement = [][]string{
			{"simTime", "forceId", "life", "lon", "lat", "alt", "heading", "pitch", "roll", "speed", "remainingMileage",
				"life2", "lon2", "lat2", "alt2", "heading2", "pitch2", "roll2", "speed2"},
		}
		dataTypes = [][]client.TSDataType{
			{client.INT64, client.INT64, client.FLOAT, client.FLOAT, client.FLOAT, client.FLOAT, client.FLOAT, client.FLOAT, client.FLOAT, client.FLOAT, client.FLOAT,
				client.FLOAT, client.FLOAT, client.FLOAT, client.FLOAT, client.FLOAT, client.FLOAT, client.FLOAT, client.FLOAT},
		}
		values = [][]interface{}{
			{simTime, gconv.Int64(force.ForceId), gconv.Float32(force.Life), gconv.Float32(force.Lon), gconv.Float32(force.Lat), gconv.Float32(force.Alt), gconv.Float32(force.Heading), gconv.Float32(force.Pitch), gconv.Float32(force.Roll), gconv.Float32(force.Speed), gconv.Float32(force.RemainingMileage),
				gconv.Float32(force.Life2), gconv.Float32(force.Lon2), gconv.Float32(force.Lat2), gconv.Float32(force.Alt2), gconv.Float32(force.Heading2), gconv.Float32(force.Pitch2), gconv.Float32(force.Roll2), gconv.Float32(force.Speed2)},
		}
		timestamps = []int64{ts}
	)
	session, err := IotDbSessionPool.GetSession()
	defer IotDbSessionPool.PutBack(session)
	if err == nil {
		checkError(session.InsertRecordsOfOneDevice(deviceId, timestamps, measurement, dataTypes, values, false))
	}
	//defer Wg.Done()
}

func InsertTabel(forces []*entity.IotDbForcePosture, simId string, dataTime int64, simTime int64) {
	//fmt.Println(gconv.String(simId)+" start get session from sessionPool ")
	session, err := IotDbSessionPool.GetSession()
	defer IotDbSessionPool.PutBack(session)
	if err == nil {
		//fmt.Println(gconv.String(simId)+" start get tablet ")
		if tablet, err := CreatForceTablet(forces, simId, dataTime, simTime); err == nil {
			//fmt.Println(gconv.String(simId)+" start use Api to insert iotdb ")
			status, err := session.InsertTablet(tablet, false)
			checkError(status, err)
		} else {
			log.Fatal(err)
		}
	}
}

func CreatForceTablet(forces []*entity.IotDbForcePosture, simId string, dataTime int64, simTime int64) (*client.Tablet, error) {
	deviceId := "root.tinder." + simId
	tablet, err := client.NewTablet(deviceId, []*client.MeasurementSchema{
		{
			Measurement: "simTime",
			DataType:    client.INT64,
		},
		{
			Measurement: "forceId",
			DataType:    client.INT64,
		},
		{
			Measurement: "life",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "lon",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "lat",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "alt",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "heading",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "pitch",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "roll",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "speed",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "remainingMileage",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "life2",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "lon2",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "lat2",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "alt2",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "heading2",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "pitch2",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "roll2",
			DataType:    client.FLOAT,
		},
		{
			Measurement: "speed2",
			DataType:    client.FLOAT,
		},
	}, len(forces))

	if err != nil {
		return nil, err
	}
	ts := dataTime
	for row := 0; row < int(len(forces)); row++ {
		ts++
		tablet.SetTimestamp(ts, row)
		tablet.SetValueAt(gconv.Int64(simTime), 0, row)
		tablet.SetValueAt(gconv.Int64(forces[row].ForceId), 1, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Life), 2, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Lon), 3, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Lat), 4, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Alt), 5, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Heading), 6, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Pitch), 7, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Roll), 8, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Speed), 9, row)
		tablet.SetValueAt(gconv.Float32(forces[row].RemainingMileage), 10, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Life2), 11, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Lon2), 12, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Lat2), 13, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Alt2), 14, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Heading2), 15, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Pitch2), 16, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Roll2), 17, row)
		tablet.SetValueAt(gconv.Float32(forces[row].Speed2), 18, row)

	}
	return tablet, nil
}

func CleanCache(simId uint32) {
	glog.Println("收到想定" + gconv.String(simId) + "激活信号")
	sql := "delete storage group root.tinder." + gconv.String(simId)
	session, err := IotDbSessionPool.GetSession()
	defer IotDbSessionPool.PutBack(session)
	if err != nil {
		log.Println(err)
		return
	}
	sessionDataSet, err := session.ExecuteStatement(sql)
	if err == nil {
		glog.Println("已删除编号" + gconv.String(simId) + "同名想定")
		sessionDataSet.Close()
	} else {
		glog.Println("删除想定" + gconv.String(simId) + "失败")
		glog.Println(err)
	}

	sql = "set storage group to root.tinder." + gconv.String(simId)
	session1, err := IotDbSessionPool.GetSession()
	defer IotDbSessionPool.PutBack(session)
	if err != nil {
		log.Println(err)
		return
	}
	sessionDataSet, err = session1.ExecuteStatement(sql)
	if err == nil {
		glog.Println("想定" + gconv.String(simId) + "存储组创建成功！开始进行持久化------")
		sessionDataSet.Close()
	} else {
		glog.Println("想定" + gconv.String(simId) + "创建存储组失败！")
		glog.Println(err)
	}

}

func checkError(status *rpc.TSStatus, err error) {
	if err != nil {
		log.Println("Iotdb Api Error")
		log.Fatal(err)
	}
	if status != nil {
		if err = client.VerifySuccess(status); err != nil {
			log.Println(err)
			log.Println("Iotdb rpc Error")
		}
	}
}
