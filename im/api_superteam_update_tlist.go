package im

const _API_SUPERTEAM_UPDATE_TLIST_URL = "https://api.netease.im/nimserver/superteam/updateTlist.action"

type SuperteamUpdateTlistParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 要修改的用户对应的accid
	// 是否必须: true
	Accid string `json:"accid"`

	// 1:关闭消息提醒，0:打开消息提醒，其他值无效
	// 是否必须: false
	SilentType string `json:"silentType,omitempty"`

	// 群成员昵称，最大长度32字符
	// 是否必须: false
	Nick string `json:"nick,omitempty"`

	// 自定义扩展字段，最大长度32字符
	// 是否必须: false
	Custom string `json:"custom,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#修改群成员信息
// 修改群成员信息
func (y *YunxinIM) ApiSuperteamUpdateTlist(param *SuperteamUpdateTlistParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_UPDATE_TLIST_URL, param)
}
