#pragma once

// 兵力姿态数据，每步长各节点同步
typedef struct ForcePosture
{
	unsigned int id;			// 兵力id
	double lon;					// 经度
	double lat;					// 纬度
	double alt;					// 高度
	double heading;				// 航向角
	double pitch;				// 偏转角
	double roll;				// 滚转角
	double speed;				// 速度
	double life;				// 生命值
	double remainingMileage;	// 剩余续航里程
} ForcePosture;

// 探测到的兵力态势
typedef struct DetectedPostore
{
	unsigned int id;			// 兵力id
	double lon;					// 经度
	double lat;					// 纬度
	double alt;					// 高度
	double heading;				// 航向角
	double pitch;				// 偏转角
	double speed;				// 速度
} DetectedPostore;
