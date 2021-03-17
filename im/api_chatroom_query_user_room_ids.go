package im

const _API_CHATROOM_QUERY_USER_ROOM_IDS_URL = "https://api.netease.im/nimserver/chatroom/queryUserRoomIds.action"

type ChatroomQueryUserRoomIdsParam struct {
	// 聊天室创建者accid
	// 是否必须: true
	Creator string `json:"creator"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E6%9F%A5%E8%AF%A2%E7%94%A8%E6%88%B7%E5%88%9B%E5%BB%BA%E7%9A%84%E5%BC%80%E5%90%AF%E7%8A%B6%E6%80%81%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%88%97%E8%A1%A8
// 查询用户创建的开启状态聊天室列表
func (y *YunxinIM) ApiChatroomQueryUserRoomIds(creator string) *ImResp {
	return y.PostFrom(_API_CHATROOM_QUERY_USER_ROOM_IDS_URL, ChatroomQueryUserRoomIdsParam{creator})
}
