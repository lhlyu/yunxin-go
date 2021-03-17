package im

const _API_MSG_UPLOAD_URL = "https://api.netease.im/nimserver/msg/upload.action"

type MsgUploadParam struct {
	// 字符流base64串(Base64.encode(bytes)) ，最大15M的字符流
	// 是否必须: true
	Content string `json:"content"`

	// 上传文件类型
	// 是否必须: false
	Type string `json:"type,omitempty"`

	// 返回的url是否需要为https的url，true或false，默认false
	// 是否必须: false
	Ishttps string `json:"ishttps,omitempty"`

	// 文件过期时长，单位：秒，必须大于等于86400
	// 是否必须: false
	ExpireSec int `json:"expireSec,omitempty"`

	// 文件的应用场景，不超过32个字符
	// 是否必须: false
	Tag string `json:"tag,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0
// 文件上传
func (y *YunxinIM) ApiMsgUpload(param *MsgUploadParam) *ImResp {
	return y.PostFrom(_API_MSG_UPLOAD_URL, param)
}
