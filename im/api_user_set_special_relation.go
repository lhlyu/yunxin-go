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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E8%AE%BE%E7%BD%AE%E9%BB%91%E5%90%8D%E5%8D%95/%E9%9D%99%E9%9F%B3
// 设置黑名单/静音
func (y *YunxinIM) ApiUserSetSpecialRelation(param *UserSetSpecialRelationParam) *ImResp {
	return y.PostFrom(_API_USER_SET_SPECIAL_RELATION_URL, param)
}
