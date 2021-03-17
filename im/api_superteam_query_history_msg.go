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

	// 查询指定的多个消息类型，类型之间用","分割，不设置该参数则查询全部类型消息。 类型支持，1:图片，2:语音，3:视频，4:地理位置，5:通知，6:文件，10:提示，100:自定义
	// 是否必须: false
	Type string `json:"type,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E6%9F%A5%E8%AF%A2%E4%BA%91%E7%AB%AF%E5%8E%86%E5%8F%B2%E6%B6%88%E6%81%AF
// 超级群查询云端历史消息
func (y *YunxinIM) ApiSuperteamQueryHistoryMsg(param *SuperteamQueryHistoryMsgParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_QUERY_HISTORY_MSG_URL, param)
}
