package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// 当参数不大于LessNum时，采用具名参数
const LessNum = 3

// 用于生成im的sdk
// 读取网站的数据，解析，生成
func main() {
	dir := "./im"
	clean(dir)
	docs := make([]string, 0)
	var total int
	for _, data := range datas {
		if data.Action == "" {
			continue
		}
		total++
		name := extractName(data.Action)
		entity, fields := parse(name.ParamName, data.Param)
		args, fill := LessArgs(fields)
		model := Model{
			FilePath:    path.Join(dir, name.FileName),
			ActionName:  name.ActionName,
			ActionUrl:   data.Action,
			ParamEntity: entity,
			Doc:         data.Doc,
			FuncNote:    data.Note,
			FuncName:    name.FuncName,
			ParamName:   name.ParamName,
			Args:        args,
			Fill:        fill,
			ContentType: data.ContentType,
		}
		fnNote, fn := getFunction(write(model))
		docs = append(docs, fnNote, fn)
	}
	if len(docs)/2 != total {
		log.Fatal("生成数量不对", len(docs), total)
	}
	genDoc(dir, docs)
	// 生成其他文件
	genFiles(dir, map[string]interface{}{
		"Docs": docs,
	})
}

func LessArgs(fields []*Field) (args, fill string) {
	if len(fields) > LessNum {
		return
	}
	argsItems := make([]string, 0)
	fillItems := make([]string, 0)
	for _, field := range fields {
		argsItems = append(argsItems, fmt.Sprintf("%s %s", littleCamelCase(handlerKeyWord(field.Name)), strings.ToLower(field.Kind)))
		fillItems = append(fillItems, littleCamelCase(handlerKeyWord(field.Name)))
	}
	return strings.Join(argsItems, ","), strings.Join(fillItems, ",")
}

// 处理go关键字
func handlerKeyWord(keyword string) string {
	switch keyword {
	case "type":
		return "kind"
	}
	return keyword
}

const fnlength = len("func (y *YunxinIM) ")

// 提取方法
func getFunction(s string) (string, string) {
	lines := strings.Split(s, "\n")
	var note string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Index(line, "func") != 0 {
			note = line
			continue
		}
		length := len(line)
		return note, line[fnlength : length-1]
	}
	return "", ""
}

// 生成文档
func genDoc(dir string, docs []string) {
	file, err := os.OpenFile(path.Join(dir, "README.md"), os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString("# IM \n\n")
	write.WriteString("### 例子\n\n")
	write.WriteString("```go\n")
	write.WriteString(demo)
	write.WriteString("```\n\n")
	write.WriteString("### 注意\n\n")
	write.WriteString("> 消息抄送 和 第三方回调 没有实现(作者没用到)\n\n")
	write.WriteString("### 方法\n\n")
	write.WriteString("```go\n")
	for _, doc := range docs {
		write.WriteString(doc + "\n")
	}
	write.WriteString("```")
	write.Flush()
}

// 例子
const demo = `package main

import (
	"github.com/lhlyu/yunxin-go/im"
	"log"
)

func main() {
	var appKey,appSecret,accid string
	yx := im.NewIM(appKey,appSecret)
	resp := yx.ApiUserUnblock(accid)
	if resp.IsSuccess() {
		log.Println("err:",resp.Err)
		return
	}
	log.Println(resp.BodyString())
}
`
