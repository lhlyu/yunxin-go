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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E8%B8%A2%E4%BA%BA%E5%87%BA%E7%BE%A4
// 踢人出超级群
func (y *YunxinIM) ApiSuperteamKick(tid string, owner string, kickAccids string) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_KICK_URL, SuperteamKickParam{tid, owner, kickAccids})
}
