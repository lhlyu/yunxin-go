package im

import "github.com/tidwall/gjson"

type ImResp struct {
	gjson.Result

	Err  error
	Body []byte
}

// 没有错误 并且 状态码等于200才会成功
func (r *ImResp) IsSuccess() bool {
	if r.Err != nil {
		return false
	}
	return r.Get("code").Int() == 200
}

func (r *ImResp) BodyString() string {
	return string(r.Body)
}
