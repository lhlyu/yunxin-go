package im

const _API_SUPERTEAM_UPDATE_TINFO_URL = "https://api.netease.im/nimserver/superteam/updateTinfo.action"

type SuperteamUpdateTinfoParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主或管理员用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 群名称，最大长度64字符
	// 是否必须: false
	Tname string `json:"tname,omitempty"`

	// 群描述，最大长度512字符
	// 是否必须: false
	Intro string `json:"intro,omitempty"`

	// 群公告，最大长度1024字符
	// 是否必须: false
	Announcement string `json:"announcement,omitempty"`

	// 自定义群扩展属性，第三方可以根据此属性自定义扩展自己的群属性，最大长度1024字符
	// 是否必须: false
	ServerCustom string `json:"serverCustom,omitempty"`

	// 群头像，最大长度1024字符
	// 是否必须: false
	Icon string `json:"icon,omitempty"`

	// 申请入群模式，0-入群不需要申请，1-入群需要申请，2-不允许申请入群。其它返回414
	// 是否必须: false
	Joinmode string `json:"joinmode,omitempty"`

	// 谁可以邀请他人入群，0-管理员(默认),1-所有人。其它返回414
	// 是否必须: false
	Invitemode string `json:"invitemode,omitempty"`

	// 谁可以修改群资料，0-管理员(默认),1-所有人。其它返回414
	// 是否必须: false
	Uptinfomode string `json:"uptinfomode,omitempty"`

	// 谁可以更新群自定义属性，0-管理员(默认),1-所有人。其它返回414
	// 是否必须: false
	Upcustommode string `json:"upcustommode,omitempty"`

	// 邀请同意模式，0-邀请需要同意(默认)，1-邀请不需要同意。其它返回414
	// 是否必须: false
	Beinvitemode string `json:"beinvitemode,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#修改群信息
// 修改群信息
func (y *YunxinIM) ApiSuperteamUpdateTinfo(param *SuperteamUpdateTinfoParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_UPDATE_TINFO_URL, param)
}
