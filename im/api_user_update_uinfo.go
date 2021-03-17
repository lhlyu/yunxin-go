package im

const _API_USER_UPDATE_UINFO_URL = "https://api.netease.im/nimserver/user/updateUinfo.action"

type UserUpdateUinfoParam struct {
	// 用户帐号，最大长度32字符，必须保证一个APP内唯一
	// 是否必须: true
	Accid string `json:"accid"`

	// 用户昵称，最大长度64字符，可设置为空字符串
	// 是否必须: false
	Name string `json:"name,omitempty"`

	// 用户头像，最大长度1024字节，可设置为空字符串
	// 是否必须: false
	Icon string `json:"icon,omitempty"`

	// 用户签名，最大长度256字符，可设置为空字符串
	// 是否必须: false
	Sign string `json:"sign,omitempty"`

	// 用户email，最大长度64字符，可设置为空字符串
	// 是否必须: false
	Email string `json:"email,omitempty"`

	// 用户生日，最大长度16字符，可设置为空字符串
	// 是否必须: false
	Birth string `json:"birth,omitempty"`

	// 用户mobile，最大长度32字符，非中国大陆手机号码需要填写国家代码(如美国：+1-xxxxxxxxxx)或地区代码(如香港：+852-xxxxxxxx)，可设置为空字符串
	// 是否必须: false
	Mobile string `json:"mobile,omitempty"`

	// 用户性别，0表示未知，1表示男，2女表示女，其它会报参数错误
	// 是否必须: false
	Gender int `json:"gender,omitempty"`

	// 用户名片扩展字段，最大长度1024字符，用户可自行扩展，建议封装成JSON字符串，也可以设置为空字符串
	// 是否必须: false
	Ex string `json:"ex,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%90%8D%E7%89%87
// 更新用户名片
func (y *YunxinIM) ApiUserUpdateUinfo(param *UserUpdateUinfoParam) *ImResp {
	return y.PostFrom(_API_USER_UPDATE_UINFO_URL, param)
}
