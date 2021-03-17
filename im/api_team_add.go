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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E6%8B%89%E4%BA%BA%E5%85%A5%E7%BE%A4
// 拉人入群
func (y *YunxinIM) ApiTeamAdd(param *TeamAddParam) *ImResp {
	return y.PostFrom(_API_TEAM_ADD_URL, param)
}
