package im

const _API_SUPERTEAM_MUTE_URL = "https://api.netease.im/nimserver/superteam/mute.action"

type SuperteamMuteParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 0:解除禁言，1:禁言普通成员，3:禁言整个群(包括群主)
	// 是否必须: true
	MuteType string `json:"muteType"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#禁言群
// 禁言群
func (y *YunxinIM) ApiSuperteamMute(tid string, owner string, mutetype string) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_MUTE_URL, SuperteamMuteParam{tid, owner, mutetype})
}
