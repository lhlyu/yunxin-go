package im

const _API_SUPERTEAM_INVITE_URL = "https://api.netease.im/nimserver/superteam/invite.action"

type SuperteamInviteParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主或管理员用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 被拉入群的accid(JSONArray)，["aaa","bbb"]，一次最多操作200个
	// 是否必须: true
	InviteAccids string `json:"inviteAccids"`

	// 邀请发送的文字，最大长度150字符
	// 是否必须: true
	Msg string `json:"msg"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#拉人入群
// 拉人入群
func (y *YunxinIM) ApiSuperteamInvite(param *SuperteamInviteParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_INVITE_URL, param)
}
