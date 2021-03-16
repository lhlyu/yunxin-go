package im

const _API_MSG_RECALL_URL = "https://api.netease.im/nimserver/msg/recall.action"

type MsgRecallParam struct {
	// 要撤回消息的msgid
	// 是否必须: true
	DeleteMsgid string `json:"deleteMsgid"`

	// 要撤回消息的创建时间
	// 是否必须: true
	Timetag int64 `json:"timetag"`

	// 7:表示点对点消息撤回，8:表示群消息撤回，其它为参数错误
	// 是否必须: true
	Type int `json:"type"`

	// 发消息的accid
	// 是否必须: true
	From string `json:"from"`

	// 如果点对点消息，为接收消息的accid,如果群消息，为对应群的tid
	// 是否必须: true
	To string `json:"to"`

	// 可以带上对应的描述
	// 是否必须: false
	Msg string `json:"msg,omitempty"`

	// 1表示绕过撤回时间检测，其它为非法参数，最多撤回近30天内的消息。如果需要撤回时间检测，不填即可。
	// 是否必须: false
	IgnoreTime string `json:"ignoreTime,omitempty"`

	// 推送文案，android以此为推送显示文案；ios若未填写payload，显示文案以pushcontent为准。超过500字符后，会对文本进行截断。
	// 是否必须: false
	Pushcontent string `json:"pushcontent,omitempty"`

	// 推送对应的payload,必须是JSON,不超过2K字符
	// 是否必须: false
	Payload string `json:"payload,omitempty"`

	// 所属环境，根据env可以配置不同的抄送地址
	// 是否必须: false
	Env string `json:"env,omitempty"`

	// 扩展字段，最大5000字符
	// 是否必须: false
	Attach string `json:"attach,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/消息功能?#批量发送点对点普通消息
// 消息撤回
func (y *YunxinIM) ApiMsgRecall(param *MsgRecallParam) *ImResp {
	return y.PostFrom(_API_MSG_RECALL_URL, param)
}
