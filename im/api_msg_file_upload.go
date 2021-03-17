package im

const _API_MSG_FILE_UPLOAD_URL = "https://api.netease.im/nimserver/msg/fileUpload.action"

type MsgFileUploadParam struct {
	// 最大15M的字符流
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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0%EF%BC%88multipart%E6%96%B9%E5%BC%8F%EF%BC%89
// 文件上传（multipart方式）
func (y *YunxinIM) ApiMsgFileUpload(param *MsgFileUploadParam) *ImResp {
	return y.DoPost(_API_MSG_FILE_UPLOAD_URL, param, _FileContentType)
}
