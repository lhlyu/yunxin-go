package im

const _API_SUPERTEAM_REMOVE_MANAGER_URL = "https://api.netease.im/nimserver/superteam/removeManager.action"

type SuperteamRemoveManagerParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 要解除掉管理员的accid(JSONArray)，["aaa","bbb"]，一次最多操作10个
	// 是否必须: true
	ManagerAccids string `json:"managerAccids"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E8%A7%A3%E9%99%A4%E7%AE%A1%E7%90%86%E5%91%98
// 超级群解除管理员
func (y *YunxinIM) ApiSuperteamRemoveManager(tid string, owner string, managerAccids string) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_REMOVE_MANAGER_URL, SuperteamRemoveManagerParam{tid, owner, managerAccids})
}
