package im

const _API_SUPERTEAM_CHANGE_LEVEL_URL = "https://api.netease.im/nimserver/superteam/changeLevel.action"

type SuperteamChangeLevelParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 群人数级别，[2,200(默认)]
	// 是否必须: true
	Tlevel string `json:"tlevel"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E5%8F%98%E6%9B%B4%E7%BE%A4%E4%BA%BA%E6%95%B0%E7%BA%A7%E5%88%AB
// 变更群人数级别
func (y *YunxinIM) ApiSuperteamChangeLevel(tid string, owner string, tlevel string) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_CHANGE_LEVEL_URL, SuperteamChangeLevelParam{tid, owner, tlevel})
}
