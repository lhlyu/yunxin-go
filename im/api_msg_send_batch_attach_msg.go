package im

const _API_MSG_SEND_BATCH_ATTACH_MSG_URL = "https://api.netease.im/nimserver/msg/sendBatchAttachMsg.action"

type MsgSendBatchAttachMsgParam struct {
	// 发送者accid，用户帐号，最大32字符，APP内唯一
	// 是否必须: true
	FromAccid string `json:"fromAccid"`

	// ["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414错误），最大限500人
	// 是否必须: true
	ToAccids string `json:"toAccids"`

	// 自定义通知内容，第三方组装的字符串，建议是JSON串，最大长度4096字符
	// 是否必须: true
	Attach string `json:"attach"`

	// 推送文案，最长500个字符。具体参见 推送配置参数详解。
	// 是否必须: false
	Pushcontent string `json:"pushcontent,omitempty"`

	// 必须是JSON,不能超过2k字符。该参数与APNs推送的payload含义不同。具体参见 推送配置参数详解。
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
	// 2. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认false(ps:注意与sendBatchMsg.action接口有别)。
	// 3. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)
	// 是否必须: false
	Option string `json:"option,omitempty"`

	// 发自定义通知时，是否强制推送
	// 是否必须: false
	IsForcePush string `json:"isForcePush,omitempty"`

	// 发自定义通知时，强制推送文案，最长500个字符
	// 是否必须: false
	ForcePushContent string `json:"forcePushContent,omitempty"`

	// 所属环境，根据env可以配置不同的抄送地址
	// 是否必须: false
	Env string `json:"env,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E8%87%AA%E5%AE%9A%E4%B9%89%E7%B3%BB%E7%BB%9F%E9%80%9A%E7%9F%A5
// 批量发送点对点自定义系统通知
func (y *YunxinIM) ApiMsgSendBatchAttachMsg(param *MsgSendBatchAttachMsgParam) *ImResp {
	return y.PostFrom(_API_MSG_SEND_BATCH_ATTACH_MSG_URL, param)
}
