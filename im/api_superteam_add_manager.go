package im

const _API_SUPERTEAM_ADD_MANAGER_URL = "https://api.netease.im/nimserver/superteam/addManager.action"

type SuperteamAddManagerParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 要添加为管理员的accid(JSONArray)，["aaa","bbb"]，一次最多操作10个
	// 是否必须: true
	ManagerAccids string `json:"managerAccids"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#添加管理员
// 添加管理员
func (y *YunxinIM) ApiSuperteamAddManager(tid string, owner string, manageraccids string) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_ADD_MANAGER_URL, SuperteamAddManagerParam{tid, owner, manageraccids})
}
