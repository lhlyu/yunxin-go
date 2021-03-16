package im

const _API_JOB_NOS_DEL_URL = "https://api.netease.im/nimserver/job/nos/del.action"

type JobNosDelParam struct {
	// 被清理文件的开始时间(毫秒级)
	// 是否必须: true
	StartTime int64 `json:"startTime"`

	// 被清理文件的结束时间(毫秒级)
	// 是否必须: true
	EndTime int64 `json:"endTime"`

	// 被清理的文件类型，文件类型包含contentType则被清理
	// 如原始文件类型为"image/png"，contentType参数为"image",则满足被清理条件
	// 是否必须: false
	ContentType string `json:"contentType,omitempty"`

	// 被清理文件的应用场景，完全相同才被清理
	// 如上传文件时知道场景为"usericon",tag参数为"usericon"，则满足被清理条件
	// 是否必须: false
	Tag string `json:"tag,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/消息功能?#批量发送点对点普通消息
// 上传NOS文件清理任务
func (y *YunxinIM) ApiJobNosDel(param *JobNosDelParam) *ImResp {
	return y.PostFrom(_API_JOB_NOS_DEL_URL, param)
}
