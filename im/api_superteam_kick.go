package im

const _API_SUPERTEAM_KICK_URL = "https://api.netease.im/nimserver/superteam/kick.action"

type SuperteamKickParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主或管理员用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 被踢出群的accid(JSONArray)，["aaa","bbb"]，一次最多操作200个
	// 是否必须: true
	KickAccids string `json:"kickAccids"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#踢人出群
// 踢人出群
func (y *YunxinIM) ApiSuperteamKick(tid string, owner string, kickaccids string) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_KICK_URL, SuperteamKickParam{tid, owner, kickaccids})
}
