package im

const _API_MSG_DEL_ROAM_SESSION_URL = "https://api.netease.im/nimserver/msg/delRoamSession.action"

type MsgDelRoamSessionParam struct {
	// 会话类型，1-p2p会话，2-群会话，其他返回414
	// 是否必须: true
	Type int `json:"type"`

	// 发送者accid,
	// 用户帐号，最大长度32字节
	// 是否必须: true
	From string `json:"from"`

	// type=1表示对端accid，type=2表示tid
	// 是否必须: true
	To string `json:"to"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/消息功能?#批量发送点对点普通消息
// 删除会话漫游
func (y *YunxinIM) ApiMsgDelRoamSession(kind int, from string, to string) *ImResp {
	return y.PostFrom(_API_MSG_DEL_ROAM_SESSION_URL, MsgDelRoamSessionParam{kind, from, to})
}
