package im

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

type IM interface {
	// 创建网易云信IM账号
	ApiUserCreate(param *UserCreateParam) *ImResp
	// 更新网易云信IM token
	ApiUserUpdate(accid string, props string, token string) *ImResp
	// 重置网易云信IM token
	ApiUserRefreshToken(accid string) *ImResp
	// 封禁网易云信IM账号
	ApiUserBlock(accid string, needkick string, kicknotifyext string) *ImResp
	// 解禁网易云信IM账号
	ApiUserUnblock(accid string) *ImResp
	// 更新用户名片
	ApiUserUpdateUinfo(param *UserUpdateUinfoParam) *ImResp
	// 获取用户名片
	ApiUserGetUinfos(accids string) *ImResp
	// 设置桌面端在线时，移动端是否需要推送
	ApiUserSetDonnop(accid string, donnopopen string) *ImResp
	// 账号全局禁言
	ApiUserMute(accid string, mute bool) *ImResp
	// 加好友
	ApiFriendAdd(param *FriendAddParam) *ImResp
	// 更新好友相关信息
	ApiFriendUpdate(param *FriendUpdateParam) *ImResp
	// 删除好友
	ApiFriendDelete(accid string, faccid string, isdeletealias bool) *ImResp
	// 获取好友关系
	ApiFriendGet(accid string, updatetime int64, createtime int64) *ImResp
	// 设置黑名单/静音
	ApiUserSetSpecialRelation(param *UserSetSpecialRelationParam) *ImResp
	// 查看指定用户的黑名单和静音列表
	ApiUserListBlackAndMuteList(accid string) *ImResp
	// 给用户或者高级群发送普通消息，包括文本，图片，语音，视频和地理位置
	ApiMsgSendMsg(param *MsgSendMsgParam) *ImResp
	// 批量发送点对点普通消息
	ApiMsgSendBatchMsg(param *MsgSendBatchMsgParam) *ImResp
	// 发送自定义系统通知
	ApiMsgSendAttachMsg(param *MsgSendAttachMsgParam) *ImResp
	// 批量发送点对点自定义系统通知
	ApiMsgSendBatchAttachMsg(param *MsgSendBatchAttachMsgParam) *ImResp
	// 文件上传
	ApiMsgUpload(param *MsgUploadParam) *ImResp
	// 文件上传（multipart方式）
	ApiMsgFileUpload(param *MsgFileUploadParam) *ImResp
	// 上传NOS文件清理任务
	ApiJobNosDel(param *JobNosDelParam) *ImResp
	// 消息撤回
	ApiMsgRecall(param *MsgRecallParam) *ImResp
	// 发送广播消息
	ApiMsgBroadcastMsg(param *MsgBroadcastMsgParam) *ImResp
	// 单向撤回消息
	ApiMsgDelMsgOneWay(param *MsgDelMsgOneWayParam) *ImResp
	// 删除会话漫游
	ApiMsgDelRoamSession(kind int, from string, to string) *ImResp
	// 创建群
	ApiTeamCreate(param *TeamCreateParam) *ImResp
	// 拉人入群
	ApiTeamAdd(param *TeamAddParam) *ImResp
	// 踢人出群
	ApiTeamKick(param *TeamKickParam) *ImResp
	// 解散群
	ApiTeamRemove(tid string, owner string) *ImResp
	// 编辑群资料
	ApiTeamUpdate(param *TeamUpdateParam) *ImResp
	// 群信息与成员列表查询
	ApiTeamQuery(tids string, ope int, ignoreinvalid bool) *ImResp
	// 获取群组详细信息
	ApiTeamQueryDetail(tid int64) *ImResp
	// 获取群组已读消息的已读详情信息
	ApiTeamGetMarkReadInfo(param *TeamGetMarkReadInfoParam) *ImResp
	// 移交群主
	ApiTeamChangeOwner(param *TeamChangeOwnerParam) *ImResp
	// 任命管理员
	ApiTeamAddManager(tid string, owner string, members string) *ImResp
	// 移除管理员
	ApiTeamRemoveManager(tid string, owner string, members string) *ImResp
	// 获取某用户所加入的群信息
	ApiTeamJoinTeams(accid string) *ImResp
	// 修改群昵称
	ApiTeamUpdateTeamNick(param *TeamUpdateTeamNickParam) *ImResp
	// 修改消息提醒开关
	ApiTeamMuteTeam(tid string, accid string, ope int) *ImResp
	// 禁言群成员
	ApiTeamMuteTlist(param *TeamMuteTlistParam) *ImResp
	// 主动退群
	ApiTeamLeave(tid string, accid string) *ImResp
	// 将群组整体禁言
	ApiTeamMuteTlistAll(param *TeamMuteTlistAllParam) *ImResp
	// 获取群组禁言列表
	ApiTeamListTeamMute(tid string, owner string) *ImResp
	// 创建群
	ApiSuperteamCreate(param *SuperteamCreateParam) *ImResp
	// 解散群
	ApiSuperteamDismiss(tid string, owner string) *ImResp
	// 拉人入群
	ApiSuperteamInvite(param *SuperteamInviteParam) *ImResp
	// 踢人出群
	ApiSuperteamKick(tid string, owner string, kickaccids string) *ImResp
	// 修改群信息
	ApiSuperteamUpdateTinfo(param *SuperteamUpdateTinfoParam) *ImResp
	// 获取群信息
	ApiSuperteamGetTinfos(tids string) *ImResp
	// 修改群成员信息
	ApiSuperteamUpdateTlist(param *SuperteamUpdateTlistParam) *ImResp
	// 获取群成员信息
	ApiSuperteamGetTlists(param *SuperteamGetTlistsParam) *ImResp
	// 发送普通消息
	ApiSuperteamSendMsg(param *SuperteamSendMsgParam) *ImResp
	// 查询云端历史消息
	ApiSuperteamQueryHistoryMsg(param *SuperteamQueryHistoryMsgParam) *ImResp
	// 主动退群
	ApiSuperteamLeave(tid string, accid string) *ImResp
	// 移交群主
	ApiSuperteamChangeOwner(param *SuperteamChangeOwnerParam) *ImResp
	// 添加管理员
	ApiSuperteamAddManager(tid string, owner string, manageraccids string) *ImResp
	// 解除管理员
	ApiSuperteamRemoveManager(tid string, owner string, manageraccids string) *ImResp
	// 禁言群
	ApiSuperteamMute(tid string, owner string, mutetype string) *ImResp
	// 禁言群成员
	ApiSuperteamMuteTlist(param *SuperteamMuteTlistParam) *ImResp
	// 发送自定义系统通知
	ApiSuperteamSendAttachMsg(param *SuperteamSendAttachMsgParam) *ImResp
	// 撤回消息
	ApiSuperteamRecallMsg(param *SuperteamRecallMsgParam) *ImResp
	// 变更群人数级别
	ApiSuperteamChangeLevel(tid string, owner string, tlevel string) *ImResp
	// 获取某用户所加入的群信息
	ApiSuperteamJoinTeams(accid string) *ImResp
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
}
