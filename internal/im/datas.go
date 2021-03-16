package main

// 数据
type Data struct {
	// 请求地址
	Action string
	// 参数数据
	Param string
	// API说明
	Note string
	// 文档地址
	Doc string
	// 请求体类型，默认是 application/x-www-form-urlencoded;charset=utf-8
	ContentType string
}

var datas = []Data{
	{
		Action: "https://api.netease.im/nimserver/user/create.action",
		Param: `accid	String	是	网易云信IM账号，最大长度32字符，必须保证一个
APP内唯一（只允许字母、数字、半角下划线_、
@、半角点以及半角-组成，不区分大小写，
会统一小写处理，请注意以此接口返回结果中的accid为准）。
name	String	否	网易云信IM账号昵称，最大长度64字符。
props	String	否	json属性，开发者可选填，最大长度1024字符。该参数已不建议使用。
icon	String	否	网易云信IM账号头像URL，开发者可选填，最大长度1024
token	String	否	网易云信IM账号可以指定登录IM token值，最大长度128字符，
并更新，如果未指定，会自动生成token，并在
创建成功后返回
sign	String	否	用户签名，最大长度256字符
email	String	否	用户email，最大长度64字符
birth	String	否	用户生日，最大长度16字符
mobile	String	否	用户mobile，最大长度32字符，非中国大陆手机号码需要填写国家代码(如美国：+1-xxxxxxxxxx)或地区代码(如香港：+852-xxxxxxxx)
gender	int	否	用户性别，0表示未知，1表示男，2女表示女，其它会报参数错误
ex	String	否	用户名片扩展字段，最大长度1024字符，用户可自行扩展，建议封装成JSON字符串`,
		Note: "创建网易云信IM账号",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BD%91%E6%98%93%E4%BA%91%E9%80%9A%E4%BF%A1ID",
	},
	{
		Action: "https://api.netease.im/nimserver/user/update.action",
		Param: `accid	String	是	网易云信IM账号，最大长度32字符，必须保证一个
APP内唯一
props	String	否	该参数已不建议使用。
token	String	否	网易云信IM账号可以指定登录token值，最大长度128字符`,
		Note: "更新网易云信IM token",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BD%91%E6%98%93%E4%BA%91%E9%80%9A%E4%BF%A1ID?#%E6%9B%B4%E6%96%B0%E7%BD%91%E6%98%93%E4%BA%91%E4%BF%A1IM%20token",
	},
	{
		Action: "https://api.netease.im/nimserver/user/refreshToken.action",
		Param: `accid	String	是	网易云信IM账号，最大长度32字符，必须保证一个
APP内唯一`,
		Note: "重置网易云信IM token",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BD%91%E6%98%93%E4%BA%91%E9%80%9A%E4%BF%A1ID?#%E9%87%8D%E7%BD%AE%E7%BD%91%E6%98%93%E4%BA%91%E4%BF%A1IM%20token",
	},
	{
		Action: "https://api.netease.im/nimserver/user/block.action",
		Param: `accid	String	是	网易云信IM账号，最大长度32字符，必须保证一个
APP内唯一
needkick	String	否	是否踢掉被禁用户，true或false，默认false
kickNotifyExt	String	否	踢人时的扩展字段，SDK版本需要大于等于v7.7.0`,
		Note: "封禁网易云信IM账号",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BD%91%E6%98%93%E4%BA%91%E9%80%9A%E4%BF%A1ID?#%E5%B0%81%E7%A6%81%E7%BD%91%E6%98%93%E4%BA%91%E4%BF%A1IM%E8%B4%A6%E5%8F%B7",
	},
	{
		Action: "https://api.netease.im/nimserver/user/unblock.action",
		Param: `accid	String	是	网易云信IM账号，最大长度32字符，必须保证一个
APP内唯一`,
		Note: "解禁网易云信IM账号",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BD%91%E6%98%93%E4%BA%91%E9%80%9A%E4%BF%A1ID?#%E8%A7%A3%E7%A6%81%E7%BD%91%E6%98%93%E4%BA%91%E4%BF%A1IM%E8%B4%A6%E5%8F%B7",
	},
	{
		Action: "https://api.netease.im/nimserver/user/updateUinfo.action",
		Param: `accid	String	是	用户帐号，最大长度32字符，必须保证一个APP内唯一
name	String	否	用户昵称，最大长度64字符，可设置为空字符串
icon	String	否	用户头像，最大长度1024字节，可设置为空字符串
sign	String	否	用户签名，最大长度256字符，可设置为空字符串
email	String	否	用户email，最大长度64字符，可设置为空字符串
birth	String	否	用户生日，最大长度16字符，可设置为空字符串
mobile	String	否	用户mobile，最大长度32字符，非中国大陆手机号码需要填写国家代码(如美国：+1-xxxxxxxxxx)或地区代码(如香港：+852-xxxxxxxx)，可设置为空字符串
gender	int	否	用户性别，0表示未知，1表示男，2女表示女，其它会报参数错误
ex	String	否	用户名片扩展字段，最大长度1024字符，用户可自行扩展，建议封装成JSON字符串，也可以设置为空字符串`,
		Note: "更新用户名片",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%90%8D%E7%89%87",
	},
	{
		Action: "https://api.netease.im/nimserver/user/getUinfos.action",
		Param: `accids	String	是	用户帐号（例如：JSONArray对应的accid串，如：["zhangsan"]，如果解析出错，会报414）（一次查询最多为200）`,
		Note: "获取用户名片",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%90%8D%E7%89%87?#%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7%E5%90%8D%E7%89%87",
	},
	{
		Action: "https://api.netease.im/nimserver/user/setDonnop.action",
		Param: `accid	String	是	用户帐号
donnopOpen	String	是	桌面端在线时，移动端是否不推送：
true:移动端不需要推送，false:移动端需要推送`,
		Note: "设置桌面端在线时，移动端是否需要推送",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E8%AE%BE%E7%BD%AE?#%E8%AE%BE%E7%BD%AE%E6%A1%8C%E9%9D%A2%E7%AB%AF%E5%9C%A8%E7%BA%BF%E6%97%B6%EF%BC%8C%E7%A7%BB%E5%8A%A8%E7%AB%AF%E6%98%AF%E5%90%A6%E4%B8%8D%E6%8E%A8%E9%80%81",
	},
	{
		Action: "https://api.netease.im/nimserver/user/mute.action",
		Param: `accid	String	是	用户帐号
mute	Boolean	是	是否全局禁言：
true：全局禁言，false:取消全局禁言`,
		Note: "账号全局禁言",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E8%AE%BE%E7%BD%AE?#%E8%B4%A6%E5%8F%B7%E5%85%A8%E5%B1%80%E7%A6%81%E8%A8%80",
	},
	{
		Action: "https://api.netease.im/nimserver/friend/add.action",
		Param: `accid	String	是	加好友发起者accid
faccid	String	是	加好友接收者accid
type	int	是	1直接加好友，2请求加好友，3同意加好友，4拒绝加好友
msg	String	否	加好友对应的请求消息，第三方组装，最长256字符
serverex	String	否	服务器端扩展字段，限制长度256
此字段client端只读，server端读写`,
		Note: "加好友",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E5%8A%A0%E5%A5%BD%E5%8F%8B",
	},
	{
		Action: "https://api.netease.im/nimserver/friend/update.action",
		Param: `accid	String	是	发起者accid
faccid	String	是	要修改朋友的accid
alias	String	否	给好友增加备注名，限制长度128，可设置为空字符串
ex	String	否	修改ex字段，限制长度256，可设置为空字符串
serverex	String	否	修改serverex字段，限制长度256，可设置为空字符串
此字段client端只读，server端读写`,
		Note: "更新好友相关信息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E6%9B%B4%E6%96%B0%E5%A5%BD%E5%8F%8B%E7%9B%B8%E5%85%B3%E4%BF%A1%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/friend/delete.action",
		Param: `accid	String	是	发起者accid
faccid	String	是	要删除朋友的accid
isDeleteAlias	Boolean	否	是否需要删除备注信息
默认false:不需要，true:需要`,
		Note: "删除好友",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E5%88%A0%E9%99%A4%E5%A5%BD%E5%8F%8B",
	},
	{
		Action: "https://api.netease.im/nimserver/friend/get.action",
		Param: `accid	String	是	发起者accid
updatetime	Long	是	更新时间戳，接口返回该时间戳之后有更新的好友列表
createtime	Long	否	【Deprecated】定义同updatetime`,
		Note: "获取好友关系",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E8%8E%B7%E5%8F%96%E5%A5%BD%E5%8F%8B%E5%85%B3%E7%B3%BB",
	},
	{
		Action: "https://api.netease.im/nimserver/user/setSpecialRelation.action",
		Param: `accid	String	是	用户帐号，最大长度32字符，必须保证一个
APP内唯一
targetAcc	String	是	被加黑或加静音的帐号
relationType	int	是	本次操作的关系类型,1:黑名单操作，2:静音列表操作
value	int	是	操作值，0:取消黑名单或静音，1:加入黑名单或静音`,
		Note: "设置黑名单/静音",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E8%AE%BE%E7%BD%AE%E9%BB%91%E5%90%8D%E5%8D%95/%E9%9D%99%E9%9F%B3",
	},
	{
		Action: "https://api.netease.im/nimserver/user/listBlackAndMuteList.action",
		Param: `accid	String	是	用户帐号，最大长度32字符，必须保证一个
APP内唯一`,
		Note: "查看指定用户的黑名单和静音列表",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E6%9F%A5%E7%9C%8B%E6%8C%87%E5%AE%9A%E7%94%A8%E6%88%B7%E7%9A%84%E9%BB%91%E5%90%8D%E5%8D%95%E5%92%8C%E9%9D%99%E9%9F%B3%E5%88%97%E8%A1%A8",
	},
	// 消息功能
	{
		Action: "https://api.netease.im/nimserver/msg/sendMsg.action",
		Param: `from	String	是	发送者accid，用户帐号，最大32字符，
必须保证一个APP内唯一
ope	int	是	0：点对点个人消息，1：群消息（高级群），其他返回414
to	String	是	ope==0是表示accid即用户id，ope==1表示tid即群id
type	int	是	0 表示文本消息,
1 表示图片，
2 表示语音，
3 表示视频，
4 表示地理位置信息，
6 表示文件，
10 表示提示消息，
100 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
body	String	是	最大长度5000字符，JSON格式。
具体请参考： 消息格式示例
antispam	String	否	对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。
true或false, 默认false。
只对消息类型为：100 自定义消息类型 的消息生效。
antispamCustom	String	否	在antispam参数为true时生效。
自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：

{"type":1,"data":"custom content"}

字段说明：
1. type: 1：文本，2：图片。
2. data: 文本内容or图片地址。
option	String	否	发消息时特殊指定的行为选项,JSON格式，可用于指定消息的漫游，存云端历史，发送方多端同步，推送，消息抄送等特殊行为;option中字段不填时表示默认值 ，option示例:

{"push":false,"roam":true,"history":false,"sendersync":true,"route":false,"badge":false,"needPushNick":true}

字段说明：
1. roam: 该消息是否需要漫游，默认true（需要app开通漫游消息功能）；
2. history: 该消息是否存云端历史，默认true；
3. sendersync: 该消息是否需要发送方多端同步，默认true；
4. push: 该消息是否需要APNS推送或安卓系统通知栏推送，默认true；
5. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能);
6. badge:该消息是否需要计入到未读计数中，默认true;
7. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认true;
8. persistent: 是否需要存离线消息，不设置该参数时默认true;
9. sessionUpdate: 是否将本消息更新到会话列表服务里本会话的lastmsg，默认true。
pushcontent	String	否	推送文案,最长500个字符。具体参见 推送配置参数详解。
payload	String	否	必须是JSON,不能超过2k字符。该参数与APNs推送的payload含义不同。具体参见 推送配置参数详解。
ext	String	否	开发者扩展字段，长度限制1024字符
forcepushlist	String	否	发送群消息时的强推用户列表（云信demo中用于承载被@的成员），格式为JSONArray，如["accid1","accid2"]。若forcepushall为true，则forcepushlist为除发送者外的所有有效群成员
forcepushcontent	String	否	发送群消息时，针对强推列表forcepushlist中的用户，强制推送的内容
forcepushall	String	否	发送群消息时，强推列表是否为群里除发送者外的所有有效成员，true或false，默认为false
bid	String	否	可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
useYidun	int	否	可选，单条消息是否使用易盾反垃圾，可选值为0。
0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。

若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
yidunAntiCheating	String	否	可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
markRead	int	否	可选，群消息是否需要已读业务（仅对群消息有效），0:不需要，1:需要
checkFriend	boolean	否	是否为好友关系才发送消息，默认否
注：使用该参数需要先开通功能服务
subType	int	否	自定义消息子类型，大于0
msgSenderNoSense	int	否	发送方是否无感知。0-有感知，1-无感知。若无感知，则消息发送者无该消息的多端、漫游、历史记录等。
msgReceiverNoSense	int	否	接受方是否无感知。0-有感知，1-无感知。若无感知，则消息接收者者无该消息的多端、漫游、历史记录等
env	String	否	所属环境，根据env可以配置不同的抄送地址`,
		Note: "给用户或者高级群发送普通消息，包括文本，图片，语音，视频和地理位置",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/msg/sendBatchMsg.action",
		Param: `fromAccid	String	是	发送者accid，用户帐号，最大32字符，
必须保证一个APP内唯一
toAccids	String	是	["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414错误），限500人
type	int	是	0 表示文本消息,
1 表示图片，
2 表示语音，
3 表示视频，
4 表示地理位置信息，
6 表示文件，
10 表示提示消息，
100 自定义消息类型
body	String	是	最大长度5000字符，JSON格式。
具体请参考： 消息格式示例
option	String	否	发消息时特殊指定的行为选项,Json格式，可用于指定消息的漫游，存云端历史，发送方多端同步，推送，消息抄送等特殊行为;option中字段不填时表示默认值 option示例:

{"push":false,"roam":true,"history":false,"sendersync":true,"route":false,"badge":false,"needPushNick":true}

字段说明：
1. roam: 该消息是否需要漫游，默认true（需要app开通漫游消息功能）；
2. history: 该消息是否存云端历史，默认true；
3. sendersync: 该消息是否需要发送方多端同步，默认true；
4. push: 该消息是否需要APNS推送或安卓系统通知栏推送，默认true；
5. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能);
6. badge:该消息是否需要计入到未读计数中，默认true;
7. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认true;
8. persistent: 是否需要存离线消息，不设置该参数时默认true;
9. sessionUpdate: 是否将本消息更新到会话列表服务里本会话的lastmsg，默认true。
pushcontent	String	否	推送文案，最长500个字符。具体参见 推送配置参数详解。
payload	String	否	必须是JSON,不能超过2k字符。该参数与APNs推送的payload含义不同。具体参见 推送配置参数详解。
ext	String	否	开发者扩展字段，长度限制1024字符
bid	String	否	可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
useYidun	int	否	可选，单条消息是否使用易盾反垃圾，可选值为0。
0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。

若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
yidunAntiCheating	String	否	可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
returnMsgid	Boolean	否	是否需要返回消息ID
false：不返回消息ID（默认值）
true：返回消息ID（toAccids包含的账号数量不可以超过100个）
env	String	否	所属环境，根据env可以配置不同的抄送地址`,
		Note: "批量发送点对点普通消息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/msg/sendAttachMsg.action",
		Param: `from	String	是	发送者accid，用户帐号，最大32字符，APP内唯一
msgtype	int	是	0：点对点自定义通知，1：群消息自定义通知，其他返回414
to	String	是	msgtype==0是表示accid即用户id，msgtype==1表示tid即群id
attach	String	是	自定义通知内容，第三方组装的字符串，建议是JSON串，最大长度4096字符
pushcontent	String	否	推送文案，最长500个字符。具体参见 推送配置参数详解。
payload	String	否	必须是JSON,不能超过2k字符。该参数与APNs推送的payload含义不同。具体参见 推送配置参数详解。
sound	String	否	如果有指定推送，此属性指定为客户端本地的声音文件名，长度不要超过30个字符，如果不指定，会使用默认声音
save	int	否	1表示只发在线，2表示会存离线，其他会报414错误。默认会存离线
option	String	否	发消息时特殊指定的行为选项,Json格式，可用于指定消息计数等特殊行为;option中字段不填时表示默认值。
option示例：
{"badge":false,"needPushNick":false,"route":false}
字段说明：
1. badge:该消息是否需要计入到未读计数中，默认true;
2. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认false(ps:注意与sendMsg.action接口有别);
3. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)
isForcePush	String	否	发自定义通知时，是否强制推送
forcePushContent	String	否	发自定义通知时，强制推送文案，最长500个字符
forcePushAll	String	否	发群自定义通知时，强推列表是否为群里除发送者外的所有有效成员
forcePushList	String	否	发群自定义通知时，强推列表，格式为JSONArray，如"accid1","accid2"
env	String	否	所属环境，根据env可以配置不同的抄送地址`,
		Note: "发送自定义系统通知",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/msg/sendBatchAttachMsg.action",
		Param: `fromAccid	String	是	发送者accid，用户帐号，最大32字符，APP内唯一
toAccids	String	是	["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414错误），最大限500人
attach	String	是	自定义通知内容，第三方组装的字符串，建议是JSON串，最大长度4096字符
pushcontent	String	否	推送文案，最长500个字符。具体参见 推送配置参数详解。
payload	String	否	必须是JSON,不能超过2k字符。该参数与APNs推送的payload含义不同。具体参见 推送配置参数详解。
sound	String	否	如果有指定推送，此属性指定为客户端本地的声音文件名，长度不要超过30个字符，如果不指定，会使用默认声音
save	int	否	1表示只发在线，2表示会存离线，其他会报414错误。默认会存离线
option	String	否	发消息时特殊指定的行为选项,Json格式，可用于指定消息计数等特殊行为;option中字段不填时表示默认值。
option示例：
{"badge":false,"needPushNick":false,"route":false}
字段说明：
1. badge:该消息是否需要计入到未读计数中，默认true;
2. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认false(ps:注意与sendBatchMsg.action接口有别)。
3. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)
isForcePush	String	否	发自定义通知时，是否强制推送
forcePushContent	String	否	发自定义通知时，强制推送文案，最长500个字符
env	String	否	所属环境，根据env可以配置不同的抄送地址`,
		Note: "批量发送点对点自定义系统通知",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/msg/upload.action",
		Param: `content	String	是	字符流base64串(Base64.encode(bytes)) ，最大15M的字符流
type	String	否	上传文件类型
ishttps	String	否	返回的url是否需要为https的url，true或false，默认false
expireSec	Integer	否	文件过期时长，单位：秒，必须大于等于86400
tag	String	否	文件的应用场景，不超过32个字符`,
		Note: "文件上传",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/msg/fileUpload.action",
		Param: `content	Multipart	是	最大15M的字符流
type	String	否	上传文件类型
ishttps	String	否	返回的url是否需要为https的url，true或false，默认false
expireSec	Integer	否	文件过期时长，单位：秒，必须大于等于86400
tag	String	否	文件的应用场景，不超过32个字符`,
		Note:        "文件上传（multipart方式）",
		Doc:         "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
		ContentType: "_FileContentType",
	},
	{
		Action: "https://api.netease.im/nimserver/job/nos/del.action",
		Param: `startTime	Long	是	被清理文件的开始时间(毫秒级)
endTime	Long	是	被清理文件的结束时间(毫秒级)
contentType	String	否	被清理的文件类型，文件类型包含contentType则被清理 如原始文件类型为"image/png"，contentType参数为"image",则满足被清理条件
tag	String	否	被清理文件的应用场景，完全相同才被清理 如上传文件时知道场景为"usericon",tag参数为"usericon"，则满足被清理条件`,
		Note: "上传NOS文件清理任务",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/msg/recall.action",
		Param: `deleteMsgid	String	是	要撤回消息的msgid
timetag	long	是	要撤回消息的创建时间
type	int	是	7:表示点对点消息撤回，8:表示群消息撤回，其它为参数错误
from	String	是	发消息的accid
to	String	是	如果点对点消息，为接收消息的accid,如果群消息，为对应群的tid
msg	String	否	可以带上对应的描述
ignoreTime	String	否	1表示绕过撤回时间检测，其它为非法参数，最多撤回近30天内的消息。如果需要撤回时间检测，不填即可。
pushcontent	String	否	推送文案，android以此为推送显示文案；ios若未填写payload，显示文案以pushcontent为准。超过500字符后，会对文本进行截断。
payload	String	否	推送对应的payload,必须是JSON,不超过2K字符
env	String	否	所属环境，根据env可以配置不同的抄送地址
attach	String	否	扩展字段，最大5000字符`,
		Note: "消息撤回",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/msg/broadcastMsg.action",
		Param: `body	String	是	广播消息内容，最大4096字符
from	String	否	发送者accid, 用户帐号，最大长度32字符，必须保证一个APP内唯一
isOffline	String	否	是否存离线，true或false，默认false
ttl	int	否	存离线状态下的有效期，单位小时，默认7天
targetOs	String	否	目标客户端，默认所有客户端，jsonArray，格式：["ios","aos","pc","web","mac"]`,
		Note: "发送广播消息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/msg/delMsgOneWay.action",
		Param: `deleteMsgid	String	是	要撤回消息的msgid
timetag	long	是	要撤回消息的创建时间
type	int	是	13:表示点对点消息撤回，14:表示群消息撤回，其它为参数错误
from	String	是	发消息的accid
to	String	是	如果点对点消息，为接收消息的accid,如果群消息，为对应群的tid
msg	String	否	可以带上对应的描述`,
		Note: "单向撤回消息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/msg/delRoamSession.action",
		Param: `type	int	是	会话类型，1-p2p会话，2-群会话，其他返回414
from	String	是	发送者accid, 用户帐号，最大长度32字节
to	String	是	type=1表示对端accid，type=2表示tid`,
		Note: "删除会话漫游",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E6%89%B9%E9%87%8F%E5%8F%91%E9%80%81%E7%82%B9%E5%AF%B9%E7%82%B9%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
	},
	// 群组功能（高级群）
	{
		Action: "https://api.netease.im/nimserver/team/create.action",
		Param: `tname	String	是	群名称，最大长度64字符
owner	String	是	群主用户帐号，最大长度32字符
members	String	是	邀请的群成员列表。["aaa","bbb"](JSONArray对应的accid，如果解析出错会报414)，members与owner总和上限为200。members中无需再加owner自己的账号。
announcement	String	否	群公告，最大长度1024字符
intro	String	否	群描述，最大长度512字符
msg	String	是	邀请发送的文字，最大长度150字符
magree	int	是	管理后台建群时，0不需要被邀请人同意加入群，1需要被邀请人同意才可以加入群。其它会返回414
joinmode	int	是	群建好后，sdk操作时，0不用验证，1需要验证,2不允许任何人加入。其它返回414
custom	String	否	自定义高级群扩展属性，第三方可以跟据此属性自定义扩展自己的群属性。（建议为json）,最大长度1024字符
icon	String	否	群头像，最大长度1024字符
beinvitemode	int	否	被邀请人同意方式，0-需要同意(默认),1-不需要同意。其它返回414
invitemode	int	否	谁可以邀请他人入群，0-管理员(默认),1-所有人。其它返回414
uptinfomode	int	否	谁可以修改群资料，0-管理员(默认),1-所有人。其它返回414
upcustommode	int	否	谁可以更新群自定义属性，0-管理员(默认),1-所有人。其它返回414
teamMemberLimit	int	否	该群最大人数(包含群主)，范围：2至应用定义的最大群人数(默认:200)。其它返回414`,
		Note: "创建群",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E5%88%9B%E5%BB%BA%E7%BE%A4",
	},
	{
		Action: "https://api.netease.im/nimserver/team/add.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	用户帐号，最大长度32字符，按照群属性invitemode传入
members	String	是	["aaa","bbb"](JSONArray对应的accid，如果解析出错会报414)，一次最多拉200个成员
magree	int	是	管理后台建群时，0不需要被邀请人同意加入群，1需要被邀请人同意才可以加入群。其它会返回414
msg	String	是	邀请发送的文字，最大长度150字符
attach	String	否	自定义扩展字段，最大长度512`,
		Note: "拉人入群",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E6%8B%89%E4%BA%BA%E5%85%A5%E7%BE%A4",
	},
	{
		Action: "https://api.netease.im/nimserver/team/kick.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	管理员的accid，用户帐号，最大长度32字符
member	String	否	被移除人的accid，用户账号，最大长度32字符;注：member或members任意提供一个，优先使用member参数
members	String	否	["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414）一次最多操作200个accid; 注：member或members任意提供一个，优先使用member参数
attach	String	否	自定义扩展字段，最大长度512`,
		Note: "踢人出群",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E8%B8%A2%E4%BA%BA%E5%87%BA%E7%BE%A4",
	},
	{
		Action: "https://api.netease.im/nimserver/team/remove.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主用户帐号，最大长度32字符`,
		Note: "解散群",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E8%A7%A3%E6%95%A3%E7%BE%A4",
	},
	{
		Action: "https://api.netease.im/nimserver/team/update.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回
tname	String	否	群名称，最大长度64字符
owner	String	是	群主用户帐号，最大长度32字符
announcement	String	否	群公告，最大长度1024字符
intro	String	否	群描述，最大长度512字符
joinmode	int	否	群建好后，sdk操作时，0不用验证，1需要验证,2不允许任何人加入。其它返回414
custom	String	否	自定义高级群扩展属性，第三方可以跟据此属性自定义扩展自己的群属性。（建议为json）,最大长度1024字符
icon	String	否	群头像，最大长度1024字符
beinvitemode	int	否	被邀请人同意方式，0-需要同意(默认),1-不需要同意。其它返回414
invitemode	int	否	谁可以邀请他人入群，0-管理员(默认),1-所有人。其它返回414
uptinfomode	int	否	谁可以修改群资料，0-管理员(默认),1-所有人。其它返回414
upcustommode	int	否	谁可以更新群自定义属性，0-管理员(默认),1-所有人。其它返回414
teamMemberLimit	int	否	该群最大人数(包含群主)，范围：2至应用定义的最大群人数(默认:200)。其它返回414`,
		Note: "编辑群资料",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E7%BC%96%E8%BE%91%E7%BE%A4%E8%B5%84%E6%96%99",
	},
	{
		Action: "https://api.netease.im/nimserver/team/query.action",
		Param: `tids	String	是	群id列表，如["3083","3084"]
ope	int	是	1表示带上群成员列表，0表示不带群成员列表，只返回群信息
ignoreInvalid	Boolean	否	是否忽略无效的tid，默认为false。设置为true时将忽略无效tid，并在响应结果中返回无效的tid`,
		Note: "群信息与成员列表查询",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E7%BE%A4%E4%BF%A1%E6%81%AF%E4%B8%8E%E6%88%90%E5%91%98%E5%88%97%E8%A1%A8%E6%9F%A5%E8%AF%A2",
	},
	{
		Action: "https://api.netease.im/nimserver/team/queryDetail.action",
		Param: `tid	long	是	群id，群唯一标识，创建群时会返回`,
		Note: "获取群组详细信息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E7%BE%A4%E7%BB%84%E8%AF%A6%E7%BB%86%E4%BF%A1%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/team/getMarkReadInfo.action",
		Param: `tid	long	是	群id，群唯一标识，创建群时会返回
msgid	long	是	发送群已读业务消息时服务器返回的消息ID
fromAccid	String	是	消息发送者账号
snapshot	Boolean	否	是否返回已读、未读成员的accid列表，默认为false`,
		Note: "获取群组已读消息的已读详情信息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E7%BE%A4%E7%BB%84%E5%B7%B2%E8%AF%BB%E6%B6%88%E6%81%AF%E7%9A%84%E5%B7%B2%E8%AF%BB%E8%AF%A6%E6%83%85%E4%BF%A1%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/team/changeOwner.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主用户帐号，最大长度32字符
newowner	String	是	新群主帐号，最大长度32字符
leave	int	是	1:群主解除群主后离开群，2：群主解除群主后成为普通成员。其它414`,
		Note: "移交群主",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E7%A7%BB%E4%BA%A4%E7%BE%A4%E4%B8%BB",
	},
	{
		Action: "https://api.netease.im/nimserver/team/addManager.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主用户帐号，最大长度32字符
members	String	是	["aaa","bbb"](JSONArray对应的accid，如果解析出错会报414)，长度最大1024字符（一次添加最多10个管理员）`,
		Note: "任命管理员",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E4%BB%BB%E5%91%BD%E7%AE%A1%E7%90%86%E5%91%98",
	},
	{
		Action: "https://api.netease.im/nimserver/team/removeManager.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主用户帐号，最大长度32字符
members	String	是	["aaa","bbb"](JSONArray对应的accid，如果解析出错会报414)，长度最大1024字符（一次解除最多10个管理员）`,
		Note: "移除管理员",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E7%A7%BB%E9%99%A4%E7%AE%A1%E7%90%86%E5%91%98",
	},
	{
		Action: "https://api.netease.im/nimserver/team/joinTeams.action",
		Param: `accid	String	是	要查询用户的accid`,
		Note: "获取某用户所加入的群信息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E6%9F%90%E7%94%A8%E6%88%B7%E6%89%80%E5%8A%A0%E5%85%A5%E7%9A%84%E7%BE%A4%E4%BF%A1%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/team/updateTeamNick.action",
		Param: `tid	String	是	群唯一标识，创建群时网易云信服务器产生并返回
owner	String	是	群主 accid
accid	String	是	要修改群昵称的群成员 accid
nick	String	否	accid 对应的群昵称，最大长度32字符
custom	String	否	自定义扩展字段，最大长度1024字节`,
		Note: "修改群昵称",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E4%BF%AE%E6%94%B9%E7%BE%A4%E6%98%B5%E7%A7%B0",
	},
	{
		Action: "https://api.netease.im/nimserver/team/muteTeam.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回
accid	String	是	要操作的群成员accid
ope	int	是	1：关闭消息提醒，2：打开消息提醒，其他值无效`,
		Note: "修改消息提醒开关",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E4%BF%AE%E6%94%B9%E6%B6%88%E6%81%AF%E6%8F%90%E9%86%92%E5%BC%80%E5%85%B3",
	},
	{
		Action: "https://api.netease.im/nimserver/team/muteTlist.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回
owner	String	是	群主accid
accid	String	是	禁言对象的accid
mute	int	是	1-禁言，0-解禁`,
		Note: "禁言群成员",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E7%A6%81%E8%A8%80%E7%BE%A4%E6%88%90%E5%91%98",
	},
	{
		Action: "https://api.netease.im/nimserver/team/leave.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回
accid	String	是	退群的accid`,
		Note: "主动退群",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E4%B8%BB%E5%8A%A8%E9%80%80%E7%BE%A4",
	},
	{
		Action: "https://api.netease.im/nimserver/team/muteTlistAll.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回
owner	String	是	群主的accid
mute	String	否	true:禁言，false:解禁(mute和muteType至少提供一个，都提供时按mute处理)
muteType	int	否	禁言类型 0:解除禁言，1:禁言普通成员 3:禁言整个群(包括群主)`,
		Note: "将群组整体禁言",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E5%B0%86%E7%BE%A4%E7%BB%84%E6%95%B4%E4%BD%93%E7%A6%81%E8%A8%80",
	},
	{
		Action: "https://api.netease.im/nimserver/team/listTeamMute.action",
		Param: `tid	String	是	网易云信服务器产生，群唯一标识，创建群时会返回
owner	String	是	群主的accid`,
		Note: "获取群组禁言列表",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E7%BE%A4%E7%BB%84%E7%A6%81%E8%A8%80%E5%88%97%E8%A1%A8",
	},
	// 群组功能（超大群）
	{
		Action: "https://api.netease.im/nimserver/superteam/create.action",
		Param: `owner	String	是	群主用户帐号，最大长度32字符
inviteAccids	String	是	邀请的群成员列表。["aaa","bbb"](JSONArray对应的accid，如果解析出错会报414)，inviteAccids与owner总和上限为200。inviteAccids中无需再加owner自己的账号。
tname	String	是	群名称，最大长度64字符
intro	String	否	群描述，最大长度512字符
announcement	String	否	群公告，最大长度1024字符
serverCustom	String	否	自定义群扩展属性，第三方可以根据此属性自定义扩展自己的群属性，最大长度1024字符
icon	String	否	群头像，最大长度1024字符
msg	String	是	邀请发送的文字，最大长度150字符
joinmode	String	是	申请入群模式，0-入群不需要申请，1-入群需要申请，2-不允许申请入群。其它返回414
beinvitemode	String	否	邀请同意模式，0-邀请需要同意(默认)，1-邀请不需要同意。其它返回414
invitemode	String	否	谁可以邀请他人入群，0-管理员(默认),1-所有人。其它返回414
uptinfomode	String	否	谁可以修改群资料，0-管理员(默认),1-所有人。其它返回414
upcustommode	String	否	谁可以更新群自定义属性，0-管理员(默认),1-所有人。其它返回414
tlevel	String	否	群人数级别，[2,200(默认)]`,
		Note: "创建群",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/dismiss.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主用户帐号，最大长度32字符`,
		Note: "解散群",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E8%A7%A3%E6%95%A3%E7%BE%A4",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/invite.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主或管理员用户帐号，最大长度32字符
inviteAccids	String	是	被拉入群的accid(JSONArray)，["aaa","bbb"]，一次最多操作200个
msg	String	是	邀请发送的文字，最大长度150字符`,
		Note: "拉人入群",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E6%8B%89%E4%BA%BA%E5%85%A5%E7%BE%A4",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/kick.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主或管理员用户帐号，最大长度32字符
kickAccids	String	是	被踢出群的accid(JSONArray)，["aaa","bbb"]，一次最多操作200个`,
		Note: "踢人出群",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E8%B8%A2%E4%BA%BA%E5%87%BA%E7%BE%A4",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/updateTinfo.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主或管理员用户帐号，最大长度32字符
tname	String	否	群名称，最大长度64字符
intro	String	否	群描述，最大长度512字符
announcement	String	否	群公告，最大长度1024字符
serverCustom	String	否	自定义群扩展属性，第三方可以根据此属性自定义扩展自己的群属性，最大长度1024字符
icon	String	否	群头像，最大长度1024字符
joinmode	String	否	申请入群模式，0-入群不需要申请，1-入群需要申请，2-不允许申请入群。其它返回414
invitemode	String	否	谁可以邀请他人入群，0-管理员(默认),1-所有人。其它返回414
uptinfomode	String	否	谁可以修改群资料，0-管理员(默认),1-所有人。其它返回414
upcustommode	String	否	谁可以更新群自定义属性，0-管理员(默认),1-所有人。其它返回414
beinvitemode	String	否	邀请同意模式，0-邀请需要同意(默认)，1-邀请不需要同意。其它返回414`,
		Note: "修改群信息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E4%BF%AE%E6%94%B9%E7%BE%A4%E4%BF%A1%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/getTinfos.action",
		Param: `tids	String	是	tid列表，如["3083","3084"]`,
		Note: "获取群信息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E7%BE%A4%E4%BF%A1%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/updateTlist.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
accid	String	是	要修改的用户对应的accid
silentType	String	否	1:关闭消息提醒，0:打开消息提醒，其他值无效
nick	String	否	群成员昵称，最大长度32字符
custom	String	否	自定义扩展字段，最大长度32字符`,
		Note: "修改群成员信息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E4%BF%AE%E6%94%B9%E7%BE%A4%E6%88%90%E5%91%98%E4%BF%A1%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/getTlists.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
timetag	String	是	时间戳，单位毫秒，查询的时间起点。
limit	String	是	本次查询的条数上限(最多100条)，小于等于0，或者大于100，会提示参数错误
reverse	String	否	1:按时间正序排列，2:按时间降序排列。其它会提示参数错误。默认是1按时间正序排列`,
		Note: "获取群成员信息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E7%BE%A4%E6%88%90%E5%91%98%E4%BF%A1%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/sendMsg.action",
		Param: `tid	String	是	群tid
fromAccid	String	是	消息发送者accid，必须是群成员
type	String	是	0 表示文本消息,
1 表示图片，
2 表示语音，
3 表示视频，
4 表示地理位置信息，
6 表示文件，
100 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
body	String	是	最大长度5000字符，JSON格式。
具体请参考： 消息格式示例
antispam	String	否	对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。
true或false, 默认false。
只对消息类型为：100 自定义消息类型 的消息生效。
antispamCustom	String	否	在antispam参数为true时生效。
自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：

{"type":1,"data":"custom content"}

字段说明：
1. type: 1：文本，2：图片。
2. data: 文本内容or图片地址。
useYidun	String	否	可选，单条消息是否使用易盾反垃圾，可选值为0。
0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。
若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断。
yidunAntiCheating	String	否	可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
option	String	否	发消息时特殊指定的行为选项,JSON格式，可用于指定消息的漫游，存云端历史，发送方多端同步，消息抄送等特殊行为;option中字段不填时表示默认值 ，option示例:

{"roam":true,"history":false,"sendersync":true,"route":false}


字段说明：
1. roam: 该消息是否需要漫游，默认true（需要app开通漫游消息功能）；
2. history: 该消息是否存云端历史，默认true；
3. sendersync: 该消息是否需要发送方多端同步，默认true；
4. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)；
5. persistent: 是否需要存离线消息，不设置该参数时默认true；
6. push: 该消息是否需要推送，默认true；
7. badge: 该消息是否需要计入到未读计数中，默认true；
8. needPushNick: 推送文案是否需要带上昵称，默认true；
ext	String	否	开发者扩展字段，长度限制1024字符
pushContent	String	否	推送内容，不超过500字符
pushPayload	String	否	推送对应的payload，必须是JSON，不能超过2k字符
isForcePush	String	否	发送消息时，是否强制推送
forcePushContent	String	否	发送消息时，强制推送的内容
forcePushAll	String	否	发送消息时，强推（@操作）列表是否为群里除发送者外的所有有效成员
forcePushList	String	否	发送消息时，强推（@操作）列表，格式为JSONArray，如"accid1","accid2"
env	String	否	所属环境，根据env可以配置不同的抄送地址`,
		Note: "发送普通消息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E5%8F%91%E9%80%81%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/queryHistoryMsg.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
accid	String	是	查询用户对应的accid
begintime	String	是	开始时间，ms
endtime	String	是	截止时间，ms
limit	int	是	本次查询的消息条数上限(最多100条)，小于等于0，或者大于100，会提示参数错误
reverse	int	否	1按时间正序排列，2按时间降序排列，其它返回参数414错误，默认是按降序排列
type	String	否	查询指定的多个消息类型，类型之间用","分割，不设置该参数则查询全部类型消息。 类型支持，1:图片，2:语音，3:视频，4:地理位置，5:通知，6:文件，10:提示，100:自定义`,
		Note: "查询云端历史消息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E6%9F%A5%E8%AF%A2%E4%BA%91%E7%AB%AF%E5%8E%86%E5%8F%B2%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/leave.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
accid	String	是	要退群的用户对应的accid`,
		Note: "主动退群",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E4%B8%BB%E5%8A%A8%E9%80%80%E7%BE%A4",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/changeOwner.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主用户帐号，最大长度32字符
accid	String	是	新群主的用户对应的accid
leave	String	是	1:群主移交群主后离开此群，2:群主移交群主后成为普通成员，其它会提示参数错误`,
		Note: "移交群主",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E7%A7%BB%E4%BA%A4%E7%BE%A4%E4%B8%BB",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/addManager.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主用户帐号，最大长度32字符
managerAccids	String	是	要添加为管理员的accid(JSONArray)，["aaa","bbb"]，一次最多操作10个`,
		Note: "添加管理员",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E6%B7%BB%E5%8A%A0%E7%AE%A1%E7%90%86%E5%91%98",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/removeManager.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主用户帐号，最大长度32字符
managerAccids	String	是	要解除掉管理员的accid(JSONArray)，["aaa","bbb"]，一次最多操作10个`,
		Note: "解除管理员",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E8%A7%A3%E9%99%A4%E7%AE%A1%E7%90%86%E5%91%98",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/mute.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主用户帐号，最大长度32字符
muteType	String	是	0:解除禁言，1:禁言普通成员，3:禁言整个群(包括群主)`,
		Note: "禁言群",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E7%A6%81%E8%A8%80%E7%BE%A4",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/muteTlist.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主或管理员用户帐号，最大长度32字符
muteAccids	String	是	要禁言的accid(JSONArray)，["aaa","bbb"]，一次最多操作10个
mute	String	是	1:禁言，0:解禁`,
		Note: "禁言群成员",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E7%A6%81%E8%A8%80%E7%BE%A4%E6%88%90%E5%91%98",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/sendAttachMsg.action",
		Param: `from	String	是	发送者accid, 用户帐号，最大长度32字符
to	String	是	群tid
attach	String	是	通知具体内容，第三方组装的字符串，是JSON串，最大长度4096字符
pushContent	String	否	推送内容，不超过500字符
pushPayload	String	否	推送对应的payload，必须是JSON，不能超过2k字符
sound	String	否	可以指定为客户端本地的声音文件，长度不要超过30个字符
option	String	否	发消息时特殊指定的行为选项，Json格式，可用于指定消息计数等特殊行为；option中字段不填时表示默认值。
option示例：

{"badge":false,"needPushNick":false,"route":false}

字段说明：
1. badge:该消息是否需要计入到未读计数中，默认true；
2. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认false(注意与sendMsg.action接口有别)；
3. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)；
isForcePush	String	否	发自定义通知时，是否强制推送
forcePushContent	String	否	发自定义通知时，强制推送文案，最长500个字符
forcePushAll	String	否	发自定义通知时，强推列表是否为群里除发送者外的所有有效成员
forcePushList	String	否	发自定义通知时，强推列表，格式为JSONArray，如"accid1","accid2"`,
		Note: "发送自定义系统通知",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E5%8F%91%E9%80%81%E8%87%AA%E5%AE%9A%E4%B9%89%E7%B3%BB%E7%BB%9F%E9%80%9A%E7%9F%A5",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/recallMsg.action",
		Param: `deleteMsgid	String	是	要撤回消息的msgid
timetag	String	是	要撤回消息的创建时间(创建时间为云信服务器生成的消息发送时间戳)
from	String	是	发送者accid, 用户帐号，最大长度32字符
to	String	是	群tid
msg	String	否	可以带上对应的描述
ignoreTime	String	否	1表示忽略撤回时间检测，0表示不忽略，其它为非法参数，默认0，如果需要撤回时间检测，不填即可
pushContent	String	否	推送内容，不超过500字符
pushPayload	String	否	推送对应的payload，必须是JSON，不能超过2k字符`,
		Note: "撤回消息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E6%92%A4%E5%9B%9E%E6%B6%88%E6%81%AF",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/changeLevel.action",
		Param: `tid	String	是	云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主用户帐号，最大长度32字符
tlevel	String	是	群人数级别，[2,200(默认)]`,
		Note: "变更群人数级别",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E5%8F%98%E6%9B%B4%E7%BE%A4%E4%BA%BA%E6%95%B0%E7%BA%A7%E5%88%AB",
	},
	{
		Action: "https://api.netease.im/nimserver/superteam/joinTeams.action",
		Param: `accid	String	是	用户帐号，最大长度32字符`,
		Note: "获取某用户所加入的群信息",
		Doc:  "https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E6%9F%90%E7%94%A8%E6%88%B7%E6%89%80%E5%8A%A0%E5%85%A5%E7%9A%84%E7%BE%A4%E4%BF%A1%E6%81%AF",
	},
	// 聊天室
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
	{
		Action: "",
		Param:  ``,
		Note:   "",
		Doc:    "",
	},
}
