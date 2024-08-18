package websocket

type Intent int

// websocket intent 声明

// EventType 事件类型
type EventType string

// 参考：https://bot.q.qq.com/wiki/develop/api/gateway/intents.html
// 忽略频道相关接口实现
const (
	IntentGuilds Intent = 1 << iota

	//Guilds包含事件：
	//- GUILD_CREATE           // 当机器人加入新guild时
	//- GUILD_UPDATE           // 当guild资料发生变更时
	//- GUILD_DELETE           // 当机器人退出guild时
	//- CHANNEL_CREATE         // 当channel被创建时
	//- CHANNEL_UPDATE         // 当channel被更新时
	//- CHANNEL_DELETE         // 当channel被删除时

	IntentGuildMembers Intent = 1 << 2

	//GuildMembers包含事件：
	//- GUILD_MEMBER_ADD       // 当成员加入时
	//- GUILD_MEMBER_UPDATE    // 当成员资料变更时
	//- GUILD_MEMBER_REMOVE    // 当成员被移除时

	IntentGUILD_MESSAGES = 1 << 9

	//GUILD_MESSAGES 包含事件  // 消息事件，仅 *私域* 机器人能够设置此 intents。
	//- MESSAGE_CREATE         // 发送消息事件，代表频道内的全部消息，而不只是 at 机器人的消息。内容与 AT_MESSAGE_CREATE 相同
	//- MESSAGE_DELETE         // 删除（撤回）消息事件

	IntentGUILD_MESSAGE_REACTIONS = 1 << 10

	//GUILD_MESSAGE_REACTIONS包含事件：
	//- MESSAGE_REACTION_ADD    // 为消息添加表情表态
	//- MESSAGE_REACTION_REMOVE // 为消息删除表情表态

	IntentDIRECT_MESSAGE = 1 << 12

	//DIRECT_MESSAGE 包含事件：
	//- DIRECT_MESSAGE_CREATE   // 当收到用户发给机器人的私信消息时
	//- DIRECT_MESSAGE_DELETE   // 删除（撤回）消息事件

	IntentOPEN_FORUMS_EVENT = 1 << 18

	//	OPEN_FORUMS_EVENT 包含事件：
	//论坛事件, 此为公域的论坛事件
	//- OPEN_FORUM_THREAD_CREATE     // 当用户创建主题时
	//- OPEN_FORUM_THREAD_UPDATE     // 当用户更新主题时
	//- OPEN_FORUM_THREAD_DELETE     // 当用户删除主题时
	//- OPEN_FORUM_POST_CREATE       // 当用户创建帖子时
	//- OPEN_FORUM_POST_DELETE       // 当用户删除帖子时
	//- OPEN_FORUM_REPLY_CREATE      // 当用户回复评论时
	//- OPEN_FORUM_REPLY_DELETE      // 当用户删除评论时

	IntentAUDIO_OR_LIVE_CHANNEL_MEMBER = 1 << 19

	//	AUDIO_OR_LIVE_CHANNEL_MEMBER 包含事件：
	//音视频/直播子频道成员进出事件
	//- AUDIO_OR_LIVE_CHANNEL_MEMBER_ENTER  // 当用户进入音视频/直播子频道
	//- AUDIO_OR_LIVE_CHANNEL_MEMBER_EXIT   // 当用户离开音视频/直播子频道

	IntentGROUP_AND_C2C_EVENT = 1 << 25

	//	GROUP_AND_C2C_EVENT 包含事件：
	//- C2C_MESSAGE_CREATE      // 用户单聊发消息给机器人时候
	//- FRIEND_ADD              // 用户添加使用机器人
	//- FRIEND_DEL              // 用户删除机器人
	//- C2C_MSG_REJECT          // 用户在机器人资料卡手动关闭"主动消息"推送
	//- C2C_MSG_RECEIVE         // 用户在机器人资料卡手动开启"主动消息"推送开关
	//- GROUP_AT_MESSAGE_CREATE // 用户在群里@机器人时收到的消息
	//- GROUP_ADD_ROBOT         // 机器人被添加到群聊
	//- GROUP_DEL_ROBOT         // 机器人被移出群聊
	//- GROUP_MSG_REJECT        // 群管理员主动在机器人资料页操作关闭通知
	//- GROUP_MSG_RECEIVE       // 群管理员主动在机器人资料页操作开启通知

	IntentINTERACTION = 1 << 26

	//	INTERACTION 包含事件：
	//- INTERACTION_CREATE     // 互动事件创建时

	IntentMESSAGE_AUDIT = 1 << 27

	//	MESSAGE_AUDIT 包含事件
	//- MESSAGE_AUDIT_PASS     // 消息审核通过
	//- MESSAGE_AUDIT_REJECT   // 消息审核不通过

	IntentFORUMS_EVENT = 1 << 28

	// FORUMS_EVENT 包含事件：
	//论坛事件，仅 *私域* 机器人能够设置此 intents。
	//- FORUM_THREAD_CREATE     // 当用户创建主题时
	//- FORUM_THREAD_UPDATE     // 当用户更新主题时
	//- FORUM_THREAD_DELETE     // 当用户删除主题时
	//- FORUM_POST_CREATE       // 当用户创建帖子时
	//- FORUM_POST_DELETE       // 当用户删除帖子时
	//- FORUM_REPLY_CREATE      // 当用户回复评论时
	//- FORUM_REPLY_DELETE      // 当用户删除评论时
	//- FORUM_PUBLISH_AUDIT_RESULT      // 当用户发表审核通过时

	IntentAUDIO_ACTION = 1 << 29

	//AUDIO_ACTION 包含事件：
	//- AUDIO_START             // 音频开始播放时
	//- AUDIO_FINISH            // 音频播放结束时
	//- AUDIO_ON_MIC            // 上麦时
	//- AUDIO_OFF_MIC           // 下麦时

	IntentPUBLIC_GUILD_MESSAGES = 1 << 30

	//PUBLIC_GUILD_MESSAGES 包含事件：
	// 消息事件，此为公域的消息事件
	//- AT_MESSAGE_CREATE       // 当收到@机器人的消息时
	//- PUBLIC_MESSAGE_DELETE   // 当频道的消息被删除时

	IntentNone Intent = 0
)

// 事件类型
const (
	//新增事件,原sdk中没有的部分，核心事件
	EventC2cMessageCreate     EventType = "C2C_MESSAGE_CREATE"      // 用户单聊发消息给机器人时候
	EventFriendAdd            EventType = "FRIEND_ADD"              // 用户添加使用机器人
	EventFriendDel            EventType = "FRIEND_DEL"              // 用户删除机器人
	EventC2cMsgReject         EventType = "C2C_MSG_REJECT"          // 用户在机器人资料卡手动关闭"主动消息"推送
	EventC2cMsgReceive        EventType = "C2C_MSG_RECEIVE"         // 用户在机器人资料卡手动开启"主动消息"推送开关
	EventGroupAtMessageCreate EventType = "GROUP_AT_MESSAGE_CREATE" // 用户在群里@机器人时收到的消息
	EventGroupAddRobot        EventType = "GROUP_ADD_ROBOT"         // 机器人被添加到群聊
	EventGroupDelRobot        EventType = "GROUP_DEL_ROBOT"         // 机器人被移出群聊
	EventGroupMsgReject       EventType = "GROUP_MSG_REJECT"        // 群管理员主动在机器人资料页操作关闭通知
	EventGroupMsgReceive      EventType = "GROUP_MSG_RECEIVE"       // 群管理员主动在机器人资料页操作开启通知

	EventGuildCreate   EventType = "GUILD_CREATE"   // 当机器人加入新guild时
	EventGuildUpdate   EventType = "GUILD_UPDATE"   // 当guild资料发生变更时
	EventGuildDelete   EventType = "GUILD_DELETE"   // 当机器人退出guild时
	EventChannelCreate EventType = "CHANNEL_CREATE" // 当channel被创建时
	EventChannelUpdate EventType = "CHANNEL_UPDATE" // 当channel被更新时
	EventChannelDelete EventType = "CHANNEL_DELETE" // 当channel被删除时

	EventGuildMemberAdd    EventType = "GUILD_MEMBER_ADD"    // 当成员加入时
	EventGuildMemberUpdate EventType = "GUILD_MEMBER_UPDATE" // 当成员资料变更时
	EventGuildMemberRemove EventType = "GUILD_MEMBER_REMOVE" // 当成员被移除时

	EventMessageCreate EventType = "MESSAGE_CREATE" // 发送消息事件，代表频道内的全部消息，而不只是 at 机器人的消息。内容与 AT_MESSAGE_CREATE 相同
	EventMessageDelete EventType = "MESSAGE_DELETE" // 删除（撤回）消息事件

	EventMessageReactionAdd    EventType = "MESSAGE_REACTION_ADD"    // 为消息添加表情表态
	EventMessageReactionRemove EventType = "MESSAGE_REACTION_REMOVE" // 为消息删除表情表态

	EventDirectMessageCreate EventType = "DIRECT_MESSAGE_CREATE" // 当收到用户发给机器人的私信消息时
	EventDirectMessageDelete EventType = "DIRECT_MESSAGE_DELETE" // 删除（撤回）消息事件

	EventOpenForumThreadCreate EventType = "OPEN_FORUM_THREAD_CREATE" // 当用户创建主题时
	EventOpenForumThreadUpdate EventType = "OPEN_FORUM_THREAD_UPDATE" // 当用户更新主题时
	EventOpenForumThreadDelete EventType = "OPEN_FORUM_THREAD_DELETE" // 当用户删除主题时
	EventOpenForumPostCreate   EventType = "OPEN_FORUM_POST_CREATE"   // 当用户创建帖子时
	EventOpenForumPostDelete   EventType = "OPEN_FORUM_POST_DELETE"   // 当用户删除帖子时
	EventOpenForumReplyCreate  EventType = "OPEN_FORUM_REPLY_CREATE"  // 当用户回复评论时
	EventOpenForumReplyDelete  EventType = "OPEN_FORUM_REPLY_DELETE"  // 当用户删除评论时

	EventAudioOrLiveChannelMemberEnter EventType = "AUDIO_OR_LIVE_CHANNEL_MEMBER_ENTER" // 当用户进入音视频/直播子频道
	EventAudioOrLiveChannelMemberExit  EventType = "AUDIO_OR_LIVE_CHANNEL_MEMBER_EXIT"  // 当用户离开音视频/直播子频道

	EventInteractionCreate EventType = "INTERACTION_CREATE" // 互动事件创建时

	EventMessageAuditPass   EventType = "MESSAGE_AUDIT_PASS" // 消息审核通过
	EventMessageAuditReject EventType = "MESSAGE_AUDIT_REJE" // 消息审核不通过

	// 论坛事件，仅 *私域* 机器人能够设置
	EventForumThreadCreate EventType = "FORUM_THREAD_CREATE"
	EventForumThreadUpdate EventType = "FORUM_THREAD_UPDATE"
	EventForumThreadDelete EventType = "FORUM_THREAD_DELETE"
	EventForumPostCreate   EventType = "FORUM_POST_CREATE"
	EventForumPostDelete   EventType = "FORUM_POST_DELETE"
	EventForumReplyCreate  EventType = "FORUM_REPLY_CREATE"
	EventForumReplyDelete  EventType = "FORUM_REPLY_DELETE"
	EventForumAuditResult  EventType = "FORUM_PUBLISH_AUDIT_RESULT"

	EventAudioStart  EventType = "AUDIO_START"   // 音频开始播放时
	EventAudioFinish EventType = "AUDIO_FINISH"  // 音频播放结束时
	EventAudioOnMic  EventType = "AUDIO_ON_MIC"  // 上麦时
	EventAudioOffMic EventType = "AUDIO_OFF_MIC" // 下麦时

	EventAtMessageCreate     EventType = "AT_MESSAGE_CREATE"     // 当收到@机器人的消息时
	EventPublicMessageDelete EventType = "PUBLIC_MESSAGE_DELETE" // 当频道的消息被删除时
)

// intentEventMap 不同 intent 对应的事件定义
var intentEventMap = map[Intent][]EventType{
	IntentGROUP_AND_C2C_EVENT: {
		EventC2cMessageCreate, EventC2cMsgReceive, EventC2cMsgReject,
		EventGroupAtMessageCreate, EventGroupMsgReceive, EventGroupMsgReject, EventGroupAtMessageCreate,
		EventFriendAdd, EventFriendDel, EventGroupAddRobot, EventGroupDelRobot,
	},
	IntentGuilds: {
		EventGuildCreate, EventGuildUpdate, EventGuildDelete,
		EventChannelCreate, EventChannelUpdate, EventChannelDelete,
	},
	IntentGuildMembers:            {EventGuildMemberAdd, EventGuildMemberUpdate, EventGuildMemberRemove},
	IntentGUILD_MESSAGES:          {EventMessageCreate, EventMessageDelete},
	IntentGUILD_MESSAGE_REACTIONS: {EventMessageReactionAdd, EventMessageReactionRemove},
	IntentDIRECT_MESSAGE:          {EventDirectMessageCreate, EventDirectMessageDelete},
	IntentOPEN_FORUMS_EVENT: {
		EventOpenForumReplyCreate, EventOpenForumReplyDelete, EventOpenForumPostDelete, EventOpenForumPostCreate,
		EventOpenForumThreadCreate, EventOpenForumThreadDelete, EventOpenForumThreadUpdate,
	},
	IntentAUDIO_OR_LIVE_CHANNEL_MEMBER: {EventAudioOrLiveChannelMemberEnter, EventAudioOrLiveChannelMemberExit},
	IntentINTERACTION:                  {EventInteractionCreate},
	IntentMESSAGE_AUDIT:                {EventMessageAuditPass, EventMessageAuditReject},
	IntentFORUMS_EVENT: {
		EventForumThreadCreate, EventForumThreadUpdate, EventForumThreadDelete, EventForumPostCreate,
		EventForumPostDelete, EventForumReplyCreate, EventForumReplyDelete, EventForumAuditResult,
	},
	IntentAUDIO_ACTION:          {EventAudioStart, EventAudioFinish, EventAudioOnMic, EventAudioOffMic},
	IntentPUBLIC_GUILD_MESSAGES: {EventAtMessageCreate, EventPublicMessageDelete},
}

var (
	EventHandlers  []interface{}
	eventIntentMap = transposeIntentEventMap(intentEventMap)
)

func Register() Intent {
	var i Intent = 0
	//计算 Intent值
	//默认单聊和群聊绑定在一起
	for _, h := range EventHandlers {
		switch handle := h.(type) {
		case MessageHandler:
			DefaultHandlers.AloneMessage = handle
			DefaultHandlers.GroupAtMessage = handle
			i = i | EventToIntent(EventC2cMessageCreate) | EventToIntent(EventGroupAtMessageCreate)
		//case AloneMessageHandler:
		//	DefaultHandlers.AloneMessage = handle
		//	i = i | EventToIntent(EventC2cMessageCreate)
		//case GroupAtMessageHandler:
		//	DefaultHandlers.GroupAtMessage = handle
		//	i = i | EventToIntent(EventGroupAtMessageCreate)
		case ReadyHandler:
			DefaultHandlers.Ready = handle
		case ErrorNotifyHandler:
			DefaultHandlers.ErrorNotify = handle
		case PlainEventHandler:
			DefaultHandlers.Plain = handle
		case AudioEventHandler:
			DefaultHandlers.Audio = handle
			i = i | EventToIntent(
				EventAudioStart, EventAudioFinish,
				EventAudioOnMic, EventAudioOffMic,
			)
		case InteractionEventHandler:
			DefaultHandlers.Interaction = handle
			i = i | EventToIntent(EventInteractionCreate)
		case MessageEventHandler:
			DefaultHandlers.Message = handle
			i = i | EventToIntent(EventMessageCreate)
		case ATMessageEventHandler:
			DefaultHandlers.ATMessage = handle
			i = i | EventToIntent(EventAtMessageCreate)
		case DirectMessageEventHandler:
			DefaultHandlers.DirectMessage = handle
			i = i | EventToIntent(EventDirectMessageCreate)
		case MessageDeleteEventHandler:
			DefaultHandlers.MessageDelete = handle
			i = i | EventToIntent(EventMessageDelete)
		case PublicMessageDeleteEventHandler:
			DefaultHandlers.PublicMessageDelete = handle
			i = i | EventToIntent(EventPublicMessageDelete)
		case DirectMessageDeleteEventHandler:
			DefaultHandlers.DirectMessageDelete = handle
			i = i | EventToIntent(EventDirectMessageDelete)
		case MessageReactionEventHandler:
			DefaultHandlers.MessageReaction = handle
			i = i | EventToIntent(EventMessageReactionAdd, EventMessageReactionRemove)
		case MessageAuditEventHandler:
			DefaultHandlers.MessageAudit = handle
			i = i | EventToIntent(EventMessageAuditPass, EventMessageAuditReject)
		default:
			panic("不支持的event handler!")
		}
	}
	return i
}

func RegisterHandler(handlers ...interface{}) {
	for _, h := range handlers {
		EventHandlers = append(EventHandlers, h)
	}
}

// transposeIntentEventMap 转置 intent 与 event 的关系，用于根据 event 找到 intent
func transposeIntentEventMap(input map[Intent][]EventType) map[EventType]Intent {
	result := make(map[EventType]Intent)
	for i, eventTypes := range input {
		for _, s := range eventTypes {
			result[s] = i
		}
	}
	return result
}

// EventToIntent 事件转换对应的Intent
func EventToIntent(events ...EventType) Intent {
	var i Intent
	for _, event := range events {
		i = i | eventIntentMap[event]
	}
	return i
}
