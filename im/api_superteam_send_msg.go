package im

const _API_SUPERTEAM_SEND_MSG_URL = "https://api.netease.im/nimserver/superteam/sendMsg.action"

type SuperteamSendMsgParam struct {
	// 群tid
	// 是否必须: true
	Tid string `json:"tid"`

	// 消息发送者accid，必须是群成员
	// 是否必须: true
	FromAccid string `json:"fromAccid"`

	// 0 表示文本消息,
	// 1 表示图片，
	// 2 表示语音，
	// 3 表示视频，
	// 4 表示地理位置信息，
	// 6 表示文件，
	// 100 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
	// 是否必须: true
	Type string `json:"type"`

	// 最大长度5000字符，JSON格式。
	// 具体请参考： 消息格式示例
	// 是否必须: true
	Body string `json:"body"`

	// 对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。
	// true或false, 默认false。
	// 只对消息类型为：100 自定义消息类型 的消息生效。
	// 是否必须: false
	Antispam string `json:"antispam,omitempty"`

	// 在antispam参数为true时生效。
	// 自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：
	// {"type":1,"data":"custom content"}
	// 字段说明：
	// 1. type: 1：文本，2：图片。
	// 2. data: 文本内容or图片地址。
	// 是否必须: false
	AntispamCustom string `json:"antispamCustom,omitempty"`

	// 可选，单条消息是否使用易盾反垃圾，可选值为0。
	// 0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。
	// 若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断。
	// 是否必须: false
	UseYidun string `json:"useYidun,omitempty"`

	// 可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
	// 是否必须: false
	YidunAntiCheating string `json:"yidunAntiCheating,omitempty"`

	// 发消息时特殊指定的行为选项,JSON格式，可用于指定消息的漫游，存云端历史，发送方多端同步，消息抄送等特殊行为;option中字段不填时表示默认值 ，option示例:
	// {"roam":true,"history":false,"sendersync":true,"route":false}
	// 字段说明：
	// 1. roam: 该消息是否需要漫游，默认true（需要app开通漫游消息功能）；
	// 2. history: 该消息是否存云端历史，默认true；
	// 3. sendersync: 该消息是否需要发送方多端同步，默认true；
	// 4. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)；
	// 5. persistent: 是否需要存离线消息，不设置该参数时默认true；
	// 6. push: 该消息是否需要推送，默认true；
	// 7. badge: 该消息是否需要计入到未读计数中，默认true；
	// 8. needPushNick: 推送文案是否需要带上昵称，默认true；
	// 是否必须: false
	Option string `json:"option,omitempty"`

	// 开发者扩展字段，长度限制1024字符
	// 是否必须: false
	Ext string `json:"ext,omitempty"`

	// 推送内容，不超过500字符
	// 是否必须: false
	PushContent string `json:"pushContent,omitempty"`

	// 推送对应的payload，必须是JSON，不能超过2k字符
	// 是否必须: false
	PushPayload string `json:"pushPayload,omitempty"`

	// 发送消息时，是否强制推送
	// 是否必须: false
	IsForcePush string `json:"isForcePush,omitempty"`

	// 发送消息时，强制推送的内容
	// 是否必须: false
	ForcePushContent string `json:"forcePushContent,omitempty"`

	// 发送消息时，强推（@操作）列表是否为群里除发送者外的所有有效成员
	// 是否必须: false
	ForcePushAll string `json:"forcePushAll,omitempty"`

	// 发送消息时，强推（@操作）列表，格式为JSONArray，如"accid1","accid2"
	// 是否必须: false
	ForcePushList string `json:"forcePushList,omitempty"`

	// 所属环境，根据env可以配置不同的抄送地址
	// 是否必须: false
	Env string `json:"env,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E5%8F%91%E9%80%81%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF
// 超级群发送普通消息
func (y *YunxinIM) ApiSuperteamSendMsg(param *SuperteamSendMsgParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_SEND_MSG_URL, param)
}
