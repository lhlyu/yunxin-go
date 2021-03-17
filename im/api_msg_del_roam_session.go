package im

const _API_MSG_DEL_ROAM_SESSION_URL = "https://api.netease.im/nimserver/msg/delRoamSession.action"

type MsgDelRoamSessionParam struct {
	// 会话类型，1-p2p会话，2-群会话，其他返回414
	// 是否必须: true
	Type int `json:"type"`

	// 发送者accid, 用户帐号，最大长度32字节
	// 是否必须: true
	From string `json:"from"`

	// type=1表示对端accid，type=2表示tid
	// 是否必须: true
	To string `json:"to"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E5%88%A0%E9%99%A4%E4%BC%9A%E8%AF%9D%E6%BC%AB%E6%B8%B8
// 删除会话漫游
func (y *YunxinIM) ApiMsgDelRoamSession(kind int, from string, to string) *ImResp {
	return y.PostFrom(_API_MSG_DEL_ROAM_SESSION_URL, MsgDelRoamSessionParam{kind, from, to})
}
