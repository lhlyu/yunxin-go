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

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/消息功能?#批量发送点对点普通消息
// 文件上传（multipart方式）
func (y *YunxinIM) ApiMsgFileUpload(param *MsgFileUploadParam) *ImResp {
	return y.DoPost(_API_MSG_FILE_UPLOAD_URL, param, _FileContentType)
}
