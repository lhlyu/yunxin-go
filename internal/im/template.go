package main

const t = `package im

const [[ .ActionName ]] = "[[ .ActionUrl ]]"

[[ .ParamEntity ]]

// doc: [[ .Doc ]]
// [[ .FuncNote ]]
func (y *YunxinIM) [[ .FuncName ]](param *[[ .ParamName ]]) *ImResp{
[[ if .ContentType ]] return y.DoPost([[ .ActionName ]],param,[[ .ContentType ]]) [[ else ]] return y.PostFrom([[ .ActionName ]],param) [[ end ]]
}
`

// 针对参数比较少的模板
const t2 = `package im

const [[ .ActionName ]] = "[[ .ActionUrl ]]"

[[ .ParamEntity ]]

// doc: [[ .Doc ]]
// [[ .FuncNote ]]
func (y *YunxinIM) [[ .FuncName ]]([[ .Args ]]) *ImResp{
[[ if .ContentType ]] return y.DoPost([[ .ActionName ]],[[ .ParamName ]]{[[ .Fill ]]},[[ .ContentType ]]) [[ else ]] return y.PostFrom([[ .ActionName ]],[[ .ParamName ]]{[[ .Fill ]]}) [[ end ]]
}
`

type Model struct {
	// 文件名
	FilePath string
	// 请求地址常量名
	ActionName string
	// 请求地址
	ActionUrl string
	// 参数实体
	ParamEntity string
	// 文档地址
	Doc string
	// 方法注释
	FuncNote string
	// 方法名字
	FuncName string
	// 参数名字
	ParamName string

	// 针对比较少参数的方法直接使用具名
	Args string
	// 填充参数
	Fill string

	// 自定义消息体
	ContentType string
}
