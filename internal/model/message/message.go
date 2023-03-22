package message

import "encoding/json"

const (
	MESSAGE_NEW_TASK                           = 0x0101 // 新建任务
	MESSAGE_TASK_CMD                           = 0x0102 // 任务控制
	MESSAGE_NEW_TASK_TEST                      = 0x0103 // 新建任务
	MESSAGE_CLOSE_TASK_TEST                    = 0x0104 // 关闭任务
	MESSAGE_TASK_NEXT_STEP                     = 0x0105 // 推动任务节点单步计算
	MESSAGE_TASK_CREATED                       = 0x0201 // task节点创建
	MESSAGE_TASK_REQUEST_SCENARIO              = 0x0202 // task向controller请求想定信息
	MESSAGE_TASK_PREPARED                      = 0x0203 // task节点准备完成
	MESSAGE_TASK_INIT_SYNC                     = 0x0204 // task节点初始同步
	MESSAGE_TASK_INIT_SYNC_FINISHED            = 0x0205 // task节点初始同步完成
	MESSAGE_TASK_NEXT_STEP_FINISHED            = 0x0206 // task节点单步计算完成
	MESSAGE_TASK_NEXT_STEP_SYNC                = 0x0207 // task节点单步计算后同步
	MESSAGE_TASK_NEXT_STEP_SYNC_FINISHED       = 0x0208 // task节点单步计算后同步完成
	MESSAGE_TASK_POST_MESSAGE                  = 0x0209 // task节点对外抛出消息事件
	MESSAGE_TASK_POST_MESSAGE_FINISHED         = 0x0210 // task节点对外抛出消息事件完成
	MESSAGE_TASK_MESSAGE_POST_PROCESS          = 0x0211 // task节点后处理消息事件
	MESSAGE_TASK_MESSAGE_POST_PROCESS_FINISHED = 0x0212 // task节点后处理消息事件完成
	MESSAGE_TASK_MESSAGE_NEXT_STEP             = 0x0213 // task节点下步长消息事件处理
	MESSAGE_TASK_MESSAGE_NEXT_STEP_FINISHED    = 0x0214 // task节点下步长消息事件处理完成
	MESSAGE_NODE_CREATED                       = 0x0301 // node节点创建
	MESSAGE_REQUEST_SCENARIO_SUMMARY           = 0x0401 // 请求想定摘要信息
	MESSAGE_RESPONSE_SCENARIO_SUMMARY          = 0x0402 // 回复想定摘要信息
	MESSAGE_SIM_FORCE_JOINED                   = 0x0501 // 仿真兵力加入至场景
	MESSAGE_SIM_FORCE_RESIGNED                 = 0x0502 // 仿真兵力从场景退出
	MESSAGE_SIM_EVENT                          = 0x0503 // 仿真事件
	MESSAGE_FUSION_CREATED                     = 0x0601 // 融合节点创建
	MESSAGE_FUSION_TO_MERGE                    = 0x0602 // 单个兵力每步长上传到融合的信息
	MESSAGE_FUSION_OUT_MERGE                   = 0x0603 // 融合模块每步长广播给兵力的消息
	MESSAGE_FUSION_INIT_SYNC_FINISHED          = 0x0604 // 融合节点初始同步完成
	MESSAGE_FUSION_NEXT_STEP_SYNC_FINISHED     = 0x0605 // 融合节点单步计算后同步完成
	MESSAGE_FUSION_DETECT                      = 0x0606 // 融合模块每步长广播探测结果
	MESSAGE_SCENARIO_EVENT                     = 0x1001 // 想定重要的变化事件
	MESSAGE_SCENARIO_TIME                      = 0x1002 // 想定时间倍率信息
)

const (
	CMD_START      = 1 // 启动
	CMD_PAUSE      = 2 // 暂停
	CMD_CONTINUE   = 3 // 继续
	CMD_STOP       = 4 // 停止
	CMD_ACCELERATE = 5 // 加速
	CMD_DECELERATE = 6 // 减速
)

type HttpResponse struct {
	HasError bool            `json:"hasError"` // 是否包含错误
	Message  string          `json:"message"`  // 若有错误，则说明错误信息
	Result   json.RawMessage `json:"result"`   // 实际返回结果
}
