package im

const _API_USER_SET_SPECIAL_RELATION_URL = "https://api.netease.im/nimserver/user/setSpecialRelation.action"

type UserSetSpecialRelationParam struct {
	// 用户帐号，最大长度32字符，必须保证一个
	// APP内唯一
	// 是否必须: true
	Accid string `json:"accid"`

	// 被加黑或加静音的帐号
	// 是否必须: true
	TargetAcc string `json:"targetAcc"`

	// 本次操作的关系类型,1:黑名单操作，2:静音列表操作
	// 是否必须: true
	RelationType int `json:"relationType"`

	// 操作值，0:取消黑名单或静音，1:加入黑名单或静音
	// 是否必须: true
	Value int `json:"value"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/用户关系托管?#设置黑名单/静音
// 设置黑名单/静音
func (y *YunxinIM) ApiUserSetSpecialRelation(param *UserSetSpecialRelationParam) *ImResp {
	return y.PostFrom(_API_USER_SET_SPECIAL_RELATION_URL, param)
}
