package im

const _API_MSG_SEND_BATCH_MSG_URL = "https://api.netease.im/nimserver/msg/sendBatchMsg.action"

type MsgSendBatchMsgParam struct {
	// 发送者accid，用户帐号，最大32字符，
	// 必须保证一个APP内唯一
	// 是否必须: true
	FromAccid string `json:"fromAccid"`

	// ["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414错误），限500人
	// 是否必须: true
	ToAccids string `json:"toAccids"`

	// 0
	// 表示文本消息,
	// 1 表示图片，
	// 2 表示语音，
	// 3 表示视频，
	// 4 表示地理位置信息，
	// 6 表示文件，
	// 10 表示提示消息，
	// 100 自定义消息类型
	// 是否必须: true
	Type int `json:"type"`

	// 最大长度5000字符，JSON格式。
	// 具体请参考： 消息格式示例
	// 是否必须: true
	Body string `json:"body"`

	// 发消息时特殊指定的行为选项,Json格式，可用于指定消息的漫游，存云端历史，发送方多端同步，推送，消息抄送等特殊行为;option中字段不填时表示默认值
	// option示例:
	// {"push":false,"roam":true,"history":false,"sendersync":true,"route":false,"badge":false,"needPushNick":true}
	// 字段说明：
	// 1. roam: 该消息是否需要漫游，默认true（需要app开通漫游消息功能）；
	// 2. history: 该消息是否存云端历史，默认true；
	// 3. sendersync: 该消息是否需要发送方多端同步，默认true；
	// 4. push: 该消息是否需要APNS推送或安卓系统通知栏推送，默认true；
	// 5. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能);
	// 6. badge:该消息是否需要计入到未读计数中，默认true;
	// 7. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认true;
	// 8. persistent: 是否需要存离线消息，不设置该参数时默认true;
	// 9. sessionUpdate: 是否将本消息更新到会话列表服务里本会话的lastmsg，默认true。
	// 是否必须: false
	Option string `json:"option,omitempty"`

	// 推送文案，最长500个字符。具体参见
	// 推送配置参数详解。
	// 是否必须: false
	Pushcontent string `json:"pushcontent,omitempty"`

	// 必须是JSON,不能超过2k字符。该参数与APNs推送的payload含义不同。具体参见
	// 推送配置参数详解。
	// 是否必须: false
	Payload string `json:"payload,omitempty"`

	// 开发者扩展字段，长度限制1024字符
	// 是否必须: false
	Ext string `json:"ext,omitempty"`

	// 可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
	// 是否必须: false
	Bid string `json:"bid,omitempty"`

	// 可选，单条消息是否使用易盾反垃圾，可选值为0。
	// 0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。
	// 若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
	// 是否必须: false
	UseYidun int `json:"useYidun,omitempty"`

	// 可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
	// 是否必须: false
	YidunAntiCheating string `json:"yidunAntiCheating,omitempty"`

	// 是否需要返回消息ID
	// false：不返回消息ID（默认值）
	// true：返回消息ID（toAccids包含的账号数量不可以超过100个）
	// 是否必须: false
	ReturnMsgid bool `json:"returnMsgid,omitempty"`

	// 所属环境，根据env可以配置不同的抄送地址
	// 是否必须: false
	Env string `json:"env,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/消息功能?#批量发送点对点普通消息
// 批量发送点对点普通消息
func (y *YunxinIM) ApiMsgSendBatchMsg(param *MsgSendBatchMsgParam) *ImResp {
	return y.PostFrom(_API_MSG_SEND_BATCH_MSG_URL, param)
}
