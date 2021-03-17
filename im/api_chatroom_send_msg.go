package im

const _API_CHATROOM_SEND_MSG_URL = "https://api.netease.im/nimserver/chatroom/sendMsg.action"

type ChatroomSendMsgParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 客户端消息id，使用uuid等随机串，msgId相同的消息会被客户端去重
	// 是否必须: true
	MsgId string `json:"msgId"`

	// 消息发出者的账号accid
	// 是否必须: true
	FromAccid string `json:"fromAccid"`

	// 消息类型：
	// 0: 表示文本消息，
	// 1: 表示图片，
	// 2: 表示语音，
	// 3: 表示视频，
	// 4: 表示地理位置信息，
	// 6: 表示文件，
	// 10: 表示Tips消息，
	// 100: 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
	// 是否必须: true
	MsgType int `json:"msgType"`

	// 重发消息标记，0：非重发消息，1：重发消息，如重发消息会按照msgid检查去重逻辑
	// 是否必须: false
	ResendFlag int `json:"resendFlag,omitempty"`

	// 文本消息：填写消息文案;
	// 其它类型消息，请参考 消息格式示例；
	// 长度限制4096字符
	// 是否必须: false
	Attach string `json:"attach,omitempty"`

	// 消息扩展字段，内容可自定义，请使用JSON格式，长度限制4096字符
	// 是否必须: false
	Ext string `json:"ext,omitempty"`

	// 是否跳过存储云端历史，0：不跳过，即存历史消息；1：跳过，即不存云端历史；默认0
	// 是否必须: false
	SkipHistory int `json:"skipHistory,omitempty"`

	// 可选，消息丢弃的概率。取值范围[0-9999]；
	// 其中0代表不丢弃消息，9999代表99.99%的概率丢弃消息，默认不丢弃；
	// 注意如果填写了此参数，下面的highPriority参数则会无效；
	// 此参数可用于流控特定业务类型的消息。
	// 是否必须: false
	AbandonRatio int `json:"abandonRatio,omitempty"`

	// 可选，true表示是高优先级消息，云信会优先保障投递这部分消息；false表示低优先级消息。默认false。
	// 强烈建议应用恰当选择参数，以便在必要时，优先保障应用内的高优先级消息的投递。若全部设置为高优先级，则等于没有设置，单个聊天室最多支持每秒10条的高优先级消息，超过的会转为普通消息。 高优先级消息可以设置进入后重发，见needHighPriorityMsgResend参数
	// 是否必须: false
	HighPriority bool `json:"highPriority,omitempty"`

	// 可选，true表示会重发消息，false表示不会重发消息。默认true。注:若设置为true， 用户离开聊天室之后重新加入聊天室，在有效期内还是会收到发送的这条消息，目前有效期默认30s。在没有配置highPriority时needHighPriorityMsgResend不生效。
	// 是否必须: false
	NeedHighPriorityMsgResend bool `json:"needHighPriorityMsgResend,omitempty"`

	// 可选，单条消息是否使用易盾反垃圾，可选值为0。
	// 0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。
	// 若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
	// 是否必须: false
	UseYidun int `json:"useYidun,omitempty"`

	// 可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
	// 是否必须: false
	YidunAntiCheating string `json:"yidunAntiCheating,omitempty"`

	// 可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
	// 是否必须: false
	Bid string `json:"bid,omitempty"`

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

	// 所属环境，根据env可以配置不同的抄送地址
	// 是否必须: false
	Env string `json:"env,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E5%8F%91%E9%80%81%E8%81%8A%E5%A4%A9%E5%AE%A4%E6%B6%88%E6%81%AF
// 发送聊天室消息
func (y *YunxinIM) ApiChatroomSendMsg(param *ChatroomSendMsgParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_SEND_MSG_URL, param)
}
