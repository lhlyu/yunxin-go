package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

const requestGoFile = `package im

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
	uv,err := qs.Values(v)
	return []byte(uv.Encode()), err
}

func (y *YunxinIM) PostFrom(uri string, v interface{}) *ImResp {
	return y.DoPost(uri, v, "")
}

func (y *YunxinIM) DoPost(uri string, v interface{}, contentType string) *ImResp {
	imResp := &ImResp{}
	if v == nil{
		err := errors.New("param is nil")
		y.LogHandler(err)
		imResp.Err = err
		return imResp
	}
	if y.LogHandler == nil{
		y.LogHandler = defaultLogHandler
	}
	if y.RandHandler == nil{
		y.RandHandler = defaultRandHandler
	}
	if contentType == "" {
		contentType = _ContentType
	}
	
	b,err := y.Format(v)
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
}`

const responseGoFile = `package im

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
`

const imGoFile = `package im

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

type IM interface {
	[[ range .Docs ]] [[ . ]] 
	[[ end ]]
}

type YunxinIM struct {
	AppKey    string
	AppSecret string

	// 自定义处理日志(错误信息)
	LogHandler func(err error)
	// 自定义随机数
	RandHandler func() string
}

func NewIM(appKey, appSecret string) IM {
	return &YunxinIM{
		AppKey:      appKey,
		AppSecret:   appSecret,
		LogHandler:  defaultLogHandler,
		RandHandler: defaultRandHandler,
	}
}

var defaultLogHandler = func(err error) {
	log.Println(err)
}

var defaultRandHandler = func() string {
	return uuid.NewV4().String()
}`

func genFiles(dir string, m map[string]interface{}) {
	genFile(path.Join(dir, "request.go"), requestGoFile)
	genFile(path.Join(dir, "response.go"), responseGoFile)
	tl := render(imGoFile, m)
	genFile(path.Join(dir, "im.go"), tl)
}

func genFile(fpath, content string) {
	file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
}
