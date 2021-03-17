package im

const _API_SUPERTEAM_SEND_ATTACH_MSG_URL = "https://api.netease.im/nimserver/superteam/sendAttachMsg.action"

type SuperteamSendAttachMsgParam struct {
	// 发送者accid, 用户帐号，最大长度32字符
	// 是否必须: true
	From string `json:"from"`

	// 群tid
	// 是否必须: true
	To string `json:"to"`

	// 通知具体内容，第三方组装的字符串，是JSON串，最大长度4096字符
	// 是否必须: true
	Attach string `json:"attach"`

	// 推送内容，不超过500字符
	// 是否必须: false
	PushContent string `json:"pushContent,omitempty"`

	// 推送对应的payload，必须是JSON，不能超过2k字符
	// 是否必须: false
	PushPayload string `json:"pushPayload,omitempty"`

	// 可以指定为客户端本地的声音文件，长度不要超过30个字符
	// 是否必须: false
	Sound string `json:"sound,omitempty"`

	// 发消息时特殊指定的行为选项，Json格式，可用于指定消息计数等特殊行为；option中字段不填时表示默认值。
	// option示例：
	// {"badge":false,"needPushNick":false,"route":false}
	// 字段说明：
	// 1. badge:该消息是否需要计入到未读计数中，默认true；
	// 2. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认false(注意与sendMsg.action接口有别)；
	// 3. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)；
	// 是否必须: false
	Option string `json:"option,omitempty"`

	// 发自定义通知时，是否强制推送
	// 是否必须: false
	IsForcePush string `json:"isForcePush,omitempty"`

	// 发自定义通知时，强制推送文案，最长500个字符
	// 是否必须: false
	ForcePushContent string `json:"forcePushContent,omitempty"`

	// 发自定义通知时，强推列表是否为群里除发送者外的所有有效成员
	// 是否必须: false
	ForcePushAll string `json:"forcePushAll,omitempty"`

	// 发自定义通知时，强推列表，格式为JSONArray，如"accid1","accid2"
	// 是否必须: false
	ForcePushList string `json:"forcePushList,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E5%8F%91%E9%80%81%E8%87%AA%E5%AE%9A%E4%B9%89%E7%B3%BB%E7%BB%9F%E9%80%9A%E7%9F%A5
// 超级群发送自定义系统通知
func (y *YunxinIM) ApiSuperteamSendAttachMsg(param *SuperteamSendAttachMsgParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_SEND_ATTACH_MSG_URL, param)
}
