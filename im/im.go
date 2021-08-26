package im

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type IM interface {
	// 创建网易云信IM账号
	ApiUserCreate(param *UserCreateParam) *ImResp
	// 更新网易云信IM token
	ApiUserUpdate(accid string, props string, token string) *ImResp
	// 重置网易云信IM token
	ApiUserRefreshToken(accid string) *ImResp
	// 封禁网易云信IM账号
	ApiUserBlock(accid string, needkick string, kickNotifyExt string) *ImResp
	// 解禁网易云信IM账号
	ApiUserUnblock(accid string) *ImResp
	// 更新用户名片
	ApiUserUpdateUinfo(param *UserUpdateUinfoParam) *ImResp
	// 获取用户名片
	ApiUserGetUinfos(accids string) *ImResp
	// 设置桌面端在线时，移动端是否需要推送
	ApiUserSetDonnop(accid string, donnopOpen string) *ImResp
	// 账号全局禁言
	ApiUserMute(accid string, mute bool) *ImResp
	// 加好友
	ApiFriendAdd(param *FriendAddParam) *ImResp
	// 更新好友相关信息
	ApiFriendUpdate(param *FriendUpdateParam) *ImResp
	// 删除好友
	ApiFriendDelete(accid string, faccid string, isDeleteAlias bool) *ImResp
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
	ApiTeamQuery(tids string, ope int, ignoreInvalid bool) *ImResp
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
	// 创建超级群
	ApiSuperteamCreate(param *SuperteamCreateParam) *ImResp
	// 解散超级群
	ApiSuperteamDismiss(tid string, owner string) *ImResp
	// 拉人入超级群
	ApiSuperteamInvite(param *SuperteamInviteParam) *ImResp
	// 踢人出超级群
	ApiSuperteamKick(tid string, owner string, kickAccids string) *ImResp
	// 修改超级群信息
	ApiSuperteamUpdateTinfo(param *SuperteamUpdateTinfoParam) *ImResp
	// 获取超级群信息
	ApiSuperteamGetTinfos(tids string) *ImResp
	// 修改超级群成员信息
	ApiSuperteamUpdateTlist(param *SuperteamUpdateTlistParam) *ImResp
	// 获取超级群成员信息
	ApiSuperteamGetTlists(param *SuperteamGetTlistsParam) *ImResp
	// 超级群发送普通消息
	ApiSuperteamSendMsg(param *SuperteamSendMsgParam) *ImResp
	// 超级群查询云端历史消息
	ApiSuperteamQueryHistoryMsg(param *SuperteamQueryHistoryMsgParam) *ImResp
	// 主动退超级群
	ApiSuperteamLeave(tid string, accid string) *ImResp
	// 超级群移交群主
	ApiSuperteamChangeOwner(param *SuperteamChangeOwnerParam) *ImResp
	// 超级群添加管理员
	ApiSuperteamAddManager(tid string, owner string, managerAccids string) *ImResp
	// 超级群解除管理员
	ApiSuperteamRemoveManager(tid string, owner string, managerAccids string) *ImResp
	// 超级群禁言群
	ApiSuperteamMute(tid string, owner string, muteType string) *ImResp
	// 超级群禁言群成员
	ApiSuperteamMuteTlist(param *SuperteamMuteTlistParam) *ImResp
	// 超级群发送自定义系统通知
	ApiSuperteamSendAttachMsg(param *SuperteamSendAttachMsgParam) *ImResp
	// 超级群撤回消息
	ApiSuperteamRecallMsg(param *SuperteamRecallMsgParam) *ImResp
	// 超级群变更群人数级别
	ApiSuperteamChangeLevel(tid string, owner string, tlevel string) *ImResp
	// 超级群获取某用户所加入的群信息
	ApiSuperteamJoinTeams(accid string) *ImResp
	// 创建聊天室
	ApiChatroomCreate(param *ChatroomCreateParam) *ImResp
	// 查询聊天室信息
	ApiChatroomGet(roomid int64, needOnlineUserCount string) *ImResp
	// 批量查询聊天室信息
	ApiChatroomGetBatch(roomids string, needOnlineUserCount string) *ImResp
	// 更新聊天室信息
	ApiChatroomUpdate(param *ChatroomUpdateParam) *ImResp
	// 修改聊天室开/关闭状态
	ApiChatroomToggleCloseStat(roomid int64, operator string, valid string) *ImResp
	// 设置聊天室内用户角色
	ApiChatroomSetMemberRole(param *ChatroomSetMemberRoleParam) *ImResp
	// 请求聊天室地址
	ApiChatroomRequestAddr(param *ChatroomRequestAddrParam) *ImResp
	// 发送聊天室消息
	ApiChatroomSendMsg(param *ChatroomSendMsgParam) *ImResp
	// 往聊天室内添加机器人
	ApiChatroomAddRobot(param *ChatroomAddRobotParam) *ImResp
	// 从聊天室内删除机器人
	ApiChatroomRemoveRobot(roomid int64, accids string) *ImResp
	// 清空聊天室机器人
	ApiChatroomCleanRobot(roomid int64, notify bool) *ImResp
	// 设置临时禁言状态
	ApiChatroomTemporaryMute(param *ChatroomTemporaryMuteParam) *ImResp
	// 往聊天室有序队列中新加或更新元素
	ApiChatroomQueueOffer(param *ChatroomQueueOfferParam) *ImResp
	// 从队列中取出元素
	ApiChatroomQueuePoll(roomid int64, key string) *ImResp
	// 排序列出队列中所有元素
	ApiChatroomQueueList(roomid int64) *ImResp
	// 删除清理整个队列
	ApiChatroomQueueDrop(roomid int64) *ImResp
	// 初始化队列
	ApiChatroomQueueInit(roomid int64, sizeLimit int64) *ImResp
	// 将聊天室整体禁言
	ApiChatroomMuteRoom(param *ChatroomMuteRoomParam) *ImResp
	// 查询聊天室统计指标TopN
	ApiStatsChatroomTopn(param *StatsChatroomTopnParam) *ImResp
	// 分页获取成员列表
	ApiChatroomMembersByPage(param *ChatroomMembersByPageParam) *ImResp
	// 批量获取在线成员信息
	ApiChatroomQueryMembers(roomid int64, accids string) *ImResp
	// 变更聊天室内的角色信息
	ApiChatroomUpdateMyRoomRole(param *ChatroomUpdateMyRoomRoleParam) *ImResp
	// 批量更新聊天室队列元素
	ApiChatroomQueueBatchUpdateElements(param *ChatroomQueueBatchUpdateElementsParam) *ImResp
	// 查询用户创建的开启状态聊天室列表
	ApiChatroomQueryUserRoomIds(creator string) *ImResp
	// 关闭指定聊天室进出通知
	ApiChatroomUpdateInOutNotification(roomid int64, close bool) *ImResp
	// 单聊云端历史消息查询
	ApiHistoryQuerySessionMsg(param *HistoryQuerySessionMsgParam) *ImResp
	// 群聊云端历史消息查询
	ApiHistoryQueryTeamMsg(param *HistoryQueryTeamMsgParam) *ImResp
	// 聊天室云端历史消息查询
	ApiHistoryQueryChatroomMsg(param *HistoryQueryChatroomMsgParam) *ImResp
	// 删除聊天室云端历史消息
	ApiChatroomDeleteHistoryMessage(roomid int64, fromAcc string, msgTimetag int64) *ImResp
	// 用户登录登出事件记录查询
	ApiHistoryQueryUserEvents(param *HistoryQueryUserEventsParam) *ImResp
	// 批量查询广播消息
	ApiHistoryQueryBroadcastMsg(broadcastId int64, limit int, kind int64) *ImResp
	// 查询单条广播消息
	ApiHistoryQueryBroadcastMsgById(broadcastId int64) *ImResp
	// 订阅在线状态事件
	ApiEventSubscribeAdd(param *EventSubscribeAddParam) *ImResp
	// 取消在线状态事件订阅
	ApiEventSubscribeDelete(accid string, eventType int, publisherAccids string) *ImResp
	// 取消全部在线状态事件订阅
	ApiEventSubscribeBatchdel(accid string, eventType int) *ImResp
	// 查询在线状态事件订阅关系
	ApiEventSubscribeQuery(accid string, eventType int, publisherAccids string) *ImResp
}

type YunxinIM struct {
	AppKey    string
	AppSecret string

	// 自定义处理日志(错误信息)
	LogHandler func(err error)
	// 自定义随机数
	RandHandler func() string
	// 自定义 Http 客户端
	HttpClient *http.Client
}

func NewIM(appKey, appSecret string) IM {
	return &YunxinIM{
		AppKey:      appKey,
		AppSecret:   appSecret,
		LogHandler:  defaultLogHandler,
		RandHandler: defaultRandHandler,
		HttpClient:  http.DefaultClient,
	}
}

var defaultLogHandler = func(err error) {
	log.Println(err)
}

var defaultRandHandler = func() string {
	return uuid.NewV4().String()
}
