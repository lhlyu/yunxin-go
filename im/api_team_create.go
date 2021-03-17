package im

const _API_TEAM_CREATE_URL = "https://api.netease.im/nimserver/team/create.action"

type TeamCreateParam struct {
	// 群名称，最大长度64字符
	// 是否必须: true
	Tname string `json:"tname"`

	// 群主用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 邀请的群成员列表。["aaa","bbb"](JSONArray对应的accid，如果解析出错会报414)，members与owner总和上限为200。members中无需再加owner自己的账号。
	// 是否必须: true
	Members string `json:"members"`

	// 群公告，最大长度1024字符
	// 是否必须: false
	Announcement string `json:"announcement,omitempty"`

	// 群描述，最大长度512字符
	// 是否必须: false
	Intro string `json:"intro,omitempty"`

	// 邀请发送的文字，最大长度150字符
	// 是否必须: true
	Msg string `json:"msg"`

	// 管理后台建群时，0不需要被邀请人同意加入群，1需要被邀请人同意才可以加入群。其它会返回414
	// 是否必须: true
	Magree int `json:"magree"`

	// 群建好后，sdk操作时，0不用验证，1需要验证,2不允许任何人加入。其它返回414
	// 是否必须: true
	Joinmode int `json:"joinmode"`

	// 自定义高级群扩展属性，第三方可以跟据此属性自定义扩展自己的群属性。（建议为json）,最大长度1024字符
	// 是否必须: false
	Custom string `json:"custom,omitempty"`

	// 群头像，最大长度1024字符
	// 是否必须: false
	Icon string `json:"icon,omitempty"`

	// 被邀请人同意方式，0-需要同意(默认),1-不需要同意。其它返回414
	// 是否必须: false
	Beinvitemode int `json:"beinvitemode,omitempty"`

	// 谁可以邀请他人入群，0-管理员(默认),1-所有人。其它返回414
	// 是否必须: false
	Invitemode int `json:"invitemode,omitempty"`

	// 谁可以修改群资料，0-管理员(默认),1-所有人。其它返回414
	// 是否必须: false
	Uptinfomode int `json:"uptinfomode,omitempty"`

	// 谁可以更新群自定义属性，0-管理员(默认),1-所有人。其它返回414
	// 是否必须: false
	Upcustommode int `json:"upcustommode,omitempty"`

	// 该群最大人数(包含群主)，范围：2至应用定义的最大群人数(默认:200)。其它返回414
	// 是否必须: false
	TeamMemberLimit int `json:"teamMemberLimit,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E5%88%9B%E5%BB%BA%E7%BE%A4
// 创建群
func (y *YunxinIM) ApiTeamCreate(param *TeamCreateParam) *ImResp {
	return y.PostFrom(_API_TEAM_CREATE_URL, param)
}
