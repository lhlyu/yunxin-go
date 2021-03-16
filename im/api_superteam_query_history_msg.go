package im

const _API_SUPERTEAM_QUERY_HISTORY_MSG_URL = "https://api.netease.im/nimserver/superteam/queryHistoryMsg.action"

type SuperteamQueryHistoryMsgParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 查询用户对应的accid
	// 是否必须: true
	Accid string `json:"accid"`

	// 开始时间，ms
	// 是否必须: true
	Begintime string `json:"begintime"`

	// 截止时间，ms
	// 是否必须: true
	Endtime string `json:"endtime"`

	// 本次查询的消息条数上限(最多100条)，小于等于0，或者大于100，会提示参数错误
	// 是否必须: true
	Limit int `json:"limit"`

	// 1按时间正序排列，2按时间降序排列，其它返回参数414错误，默认是按降序排列
	// 是否必须: false
	Reverse int `json:"reverse,omitempty"`

	// 查询指定的多个消息类型，类型之间用","分割，不设置该参数则查询全部类型消息。
	// 类型支持，1:图片，2:语音，3:视频，4:地理位置，5:通知，6:文件，10:提示，100:自定义
	// 是否必须: false
	Type string `json:"type,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#查询云端历史消息
// 查询云端历史消息
func (y *YunxinIM) ApiSuperteamQueryHistoryMsg(param *SuperteamQueryHistoryMsgParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_QUERY_HISTORY_MSG_URL, param)
}
