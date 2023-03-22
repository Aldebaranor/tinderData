package entity

type IotDbForce struct {
	Stage                uint32               `protobuf:"varint,1,opt,name=stage,proto3" json:"stage,omitempty"`
	SimTime              int64                `protobuf:"varint,2,opt,name=sim_time,json=simTime,proto3" json:"sim_time,omitempty"`
	StepRatio            float64              `protobuf:"fixed64,3,opt,name=step_ratio,json=stepRatio,proto3" json:"step_ratio,omitempty"`
	Forces               []*IotDbForcePosture `protobuf:"bytes,4,rep,name=forces,proto3" json:"forces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

type IotDbForcePosture struct {
	ForceId          uint32  `protobuf:"varint,1,opt,name=force_id,json=forceId,proto3" json:"force_id,omitempty"`
	Life             float64 `protobuf:"fixed64,6,opt,name=life,proto3" json:"life,omitempty"`
	Lon              float64 `protobuf:"fixed64,7,opt,name=lon,proto3" json:"lon,omitempty"`
	Lat              float64 `protobuf:"fixed64,8,opt,name=lat,proto3" json:"lat,omitempty"`
	Alt              float64 `protobuf:"fixed64,9,opt,name=alt,proto3" json:"alt,omitempty"`
	Heading          float64 `protobuf:"fixed64,10,opt,name=heading,proto3" json:"heading,omitempty"`
	Pitch            float64 `protobuf:"fixed64,11,opt,name=pitch,proto3" json:"pitch,omitempty"`
	Roll             float64 `protobuf:"fixed64,12,opt,name=roll,proto3" json:"roll,omitempty"`
	Speed            float64 `protobuf:"fixed64,13,opt,name=speed,proto3" json:"speed,omitempty"`
	RemainingMileage float64 `protobuf:"fixed64,14,opt,name=remaining_mileage,json=remainingMileage,proto3" json:"remaining_mileage,omitempty"`

	Life2                float64  `protobuf:"fixed64,6,opt,name=life2,proto3" json:"life2,omitempty"`
	Lon2                 float64  `protobuf:"fixed64,7,opt,name=lon2,proto3" json:"lon2,omitempty"`
	Lat2                 float64  `protobuf:"fixed64,8,opt,name=lat2,proto3" json:"lat2,omitempty"`
	Alt2                 float64  `protobuf:"fixed64,9,opt,name=alt2,proto3" json:"alt2,omitempty"`
	Heading2             float64  `protobuf:"fixed64,10,opt,name=heading2,proto3" json:"heading2,omitempty"`
	Pitch2               float64  `protobuf:"fixed64,11,opt,name=pitch2,proto3" json:"pitch2,omitempty"`
	Roll2                float64  `protobuf:"fixed64,12,opt,name=roll2,proto3" json:"roll2,omitempty"`
	Speed2               float64  `protobuf:"fixed64,13,opt,name=speed2,proto3" json:"speed2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
