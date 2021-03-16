package im

const _API_MSG_SEND_ATTACH_MSG_URL = "https://api.netease.im/nimserver/msg/sendAttachMsg.action"

type MsgSendAttachMsgParam struct {
	// 发送者accid，用户帐号，最大32字符，APP内唯一
	// 是否必须: true
	From string `json:"from"`

	// 0：点对点自定义通知，1：群消息自定义通知，其他返回414
	// 是否必须: true
	Msgtype int `json:"msgtype"`

	// msgtype==0是表示accid即用户id，msgtype==1表示tid即群id
	// 是否必须: true
	To string `json:"to"`

	// 自定义通知内容，第三方组装的字符串，建议是JSON串，最大长度4096字符
	// 是否必须: true
	Attach string `json:"attach"`

	// 推送文案，最长500个字符。具体参见
	// 推送配置参数详解。
	// 是否必须: false
	Pushcontent string `json:"pushcontent,omitempty"`

	// 必须是JSON,不能超过2k字符。该参数与APNs推送的payload含义不同。具体参见
	// 推送配置参数详解。
	// 是否必须: false
	Payload string `json:"payload,omitempty"`

	// 如果有指定推送，此属性指定为客户端本地的声音文件名，长度不要超过30个字符，如果不指定，会使用默认声音
	// 是否必须: false
	Sound string `json:"sound,omitempty"`

	// 1表示只发在线，2表示会存离线，其他会报414错误。默认会存离线
	// 是否必须: false
	Save int `json:"save,omitempty"`

	// 发消息时特殊指定的行为选项,Json格式，可用于指定消息计数等特殊行为;option中字段不填时表示默认值。
	// option示例：
	// {"badge":false,"needPushNick":false,"route":false}
	// 字段说明：
	// 1. badge:该消息是否需要计入到未读计数中，默认true;
	// 2. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认false(ps:注意与sendMsg.action接口有别);
	// 3. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)
	// 是否必须: false
	Option string `json:"option,omitempty"`

	// 发自定义通知时，是否强制推送
	// 是否必须: false
	IsForcePush string `json:"isForcePush,omitempty"`

	// 发自定义通知时，强制推送文案，最长500个字符
	// 是否必须: false
	ForcePushContent string `json:"forcePushContent,omitempty"`

	// 发群自定义通知时，强推列表是否为群里除发送者外的所有有效成员
	// 是否必须: false
	ForcePushAll string `json:"forcePushAll,omitempty"`

	// 发群自定义通知时，强推列表，格式为JSONArray，如"accid1","accid2"
	// 是否必须: false
	ForcePushList string `json:"forcePushList,omitempty"`

	// 所属环境，根据env可以配置不同的抄送地址
	// 是否必须: false
	Env string `json:"env,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/消息功能?#批量发送点对点普通消息
// 发送自定义系统通知
func (y *YunxinIM) ApiMsgSendAttachMsg(param *MsgSendAttachMsgParam) *ImResp {
	return y.PostFrom(_API_MSG_SEND_ATTACH_MSG_URL, param)
}
