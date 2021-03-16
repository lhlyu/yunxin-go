package im

import (
	"bytes"
	"github.com/lhlyu/yunxin-go/util"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	_ContentType     = "application/x-www-form-urlencoded;charset=utf-8"
	_FileContentType = "multipart/form-data;charset=utf-8"
)

func (*YunxinIM) Format(v interface{}) ([]byte, error) {
	var uv url.Values
	uv = cast.ToStringMapStringSlice(util.ToJsonStr(v))
	return []byte(uv.Encode()), nil
}

func (y *YunxinIM) PostFrom(uri string, v interface{}) *ImResp {
	return y.DoPost(uri, v, "")
}

func (y *YunxinIM) DoPost(uri string, v interface{}, contentType string) *ImResp {
	if contentType == "" {
		contentType = _ContentType
	}
	b, _ := y.Format(v)
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(b))
	imResp := &ImResp{}

	if err != nil {
		y.LogHandler(err)
		imResp.Err = err
		return imResp
	}

	nonce := y.RandHandler()
	now := strconv.FormatInt(time.Now().Unix(), 10)
	checksum := util.SHA1toStr(y.AppSecret + nonce + now)

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("AppKey", y.AppKey)
	req.Header.Set("Nonce", nonce)
	req.Header.Set("CurTime", now)
	req.Header.Set("CheckSum", checksum)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		y.LogHandler(err)
		imResp.Err = err
		return imResp
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		y.LogHandler(err)
		imResp.Err = err
		return imResp
	}

	imResp.Result = gjson.ParseBytes(body)
	imResp.Body = body
	return imResp
}
