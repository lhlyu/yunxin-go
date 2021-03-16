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

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#移交群主
// 移交群主
func (y *YunxinIM) ApiSuperteamChangeOwner(param *SuperteamChangeOwnerParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_CHANGE_OWNER_URL, param)
}
