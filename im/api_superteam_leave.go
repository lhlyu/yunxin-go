package im

const _API_SUPERTEAM_LEAVE_URL = "https://api.netease.im/nimserver/superteam/leave.action"

type SuperteamLeaveParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 要退群的用户对应的accid
	// 是否必须: true
	Accid string `json:"accid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E4%B8%BB%E5%8A%A8%E9%80%80%E7%BE%A4
// 主动退超级群
func (y *YunxinIM) ApiSuperteamLeave(tid string, accid string) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_LEAVE_URL, SuperteamLeaveParam{tid, accid})
}
