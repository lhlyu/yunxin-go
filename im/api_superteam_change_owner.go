package im

const _API_SUPERTEAM_CHANGE_OWNER_URL = "https://api.netease.im/nimserver/superteam/changeOwner.action"

type SuperteamChangeOwnerParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 新群主的用户对应的accid
	// 是否必须: true
	Accid string `json:"accid"`

	// 1:群主移交群主后离开此群，2:群主移交群主后成为普通成员，其它会提示参数错误
	// 是否必须: true
	Leave string `json:"leave"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E7%A7%BB%E4%BA%A4%E7%BE%A4%E4%B8%BB
// 移交群主
func (y *YunxinIM) ApiSuperteamChangeOwner(param *SuperteamChangeOwnerParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_CHANGE_OWNER_URL, param)
}
