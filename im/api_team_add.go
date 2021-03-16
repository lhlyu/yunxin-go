package im

const _API_TEAM_ADD_URL = "https://api.netease.im/nimserver/team/add.action"

type TeamAddParam struct {
	// 网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 用户帐号，最大长度32字符，按照群属性invitemode传入
	// 是否必须: true
	Owner string `json:"owner"`

	// ["aaa","bbb"](JSONArray对应的accid，如果解析出错会报414)，一次最多拉200个成员
	// 是否必须: true
	Members string `json:"members"`

	// 管理后台建群时，0不需要被邀请人同意加入群，1需要被邀请人同意才可以加入群。其它会返回414
	// 是否必须: true
	Magree int `json:"magree"`

	// 邀请发送的文字，最大长度150字符
	// 是否必须: true
	Msg string `json:"msg"`

	// 自定义扩展字段，最大长度512
	// 是否必须: false
	Attach string `json:"attach,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#拉人入群
// 拉人入群
func (y *YunxinIM) ApiTeamAdd(param *TeamAddParam) *ImResp {
	return y.PostFrom(_API_TEAM_ADD_URL, param)
}
