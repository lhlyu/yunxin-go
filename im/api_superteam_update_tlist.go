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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E4%BF%AE%E6%94%B9%E7%BE%A4%E6%88%90%E5%91%98%E4%BF%A1%E6%81%AF
// 修改群成员信息
func (y *YunxinIM) ApiSuperteamUpdateTlist(param *SuperteamUpdateTlistParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_UPDATE_TLIST_URL, param)
}
