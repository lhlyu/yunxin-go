package im

import (
	"bytes"
	"errors"
	"github.com/lhlyu/yunxin-go/util"
	"github.com/lhlyu/yunxin-go/util/qs"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	_ContentType     = "application/x-www-form-urlencoded;charset=utf-8"
	_FileContentType = "multipart/form-data;charset=utf-8"
)

func (*YunxinIM) Format(v interface{}) ([]byte, error) {
	uv, err := qs.Values(v)
	return []byte(uv.Encode()), err
}

func (y *YunxinIM) PostFrom(uri string, v interface{}) *ImResp {
	return y.DoPost(uri, v, "")
}

func (y *YunxinIM) DoPost(uri string, v interface{}, contentType string) *ImResp {
	imResp := &ImResp{}
	if v == nil {
		err := errors.New("param is nil")
		y.LogHandler(err)
		imResp.Err = err
		return imResp
	}
	if y.LogHandler == nil {
		y.LogHandler = defaultLogHandler
	}
	if y.RandHandler == nil {
		y.RandHandler = defaultRandHandler
	}
	if contentType == "" {
		contentType = _ContentType
	}

	b, err := y.Format(v)
	if err != nil {
		y.LogHandler(err)
		imResp.Err = err
		return imResp
	}
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(b))

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
