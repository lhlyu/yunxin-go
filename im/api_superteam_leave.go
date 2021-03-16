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

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#主动退群
// 主动退群
func (y *YunxinIM) ApiSuperteamLeave(tid string, accid string) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_LEAVE_URL, SuperteamLeaveParam{tid, accid})
}
