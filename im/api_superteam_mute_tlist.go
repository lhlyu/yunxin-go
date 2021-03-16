package im

const _API_SUPERTEAM_MUTE_TLIST_URL = "https://api.netease.im/nimserver/superteam/muteTlist.action"

type SuperteamMuteTlistParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主或管理员用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 要禁言的accid(JSONArray)，["aaa","bbb"]，一次最多操作10个
	// 是否必须: true
	MuteAccids string `json:"muteAccids"`

	// 1:禁言，0:解禁
	// 是否必须: true
	Mute string `json:"mute"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#禁言群成员
// 禁言群成员
func (y *YunxinIM) ApiSuperteamMuteTlist(param *SuperteamMuteTlistParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_MUTE_TLIST_URL, param)
}
