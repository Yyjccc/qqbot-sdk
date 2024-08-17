package openapi

import (
	"QueziyaBot/sdk/entry"
	"QueziyaBot/sdk/util"
	"QueziyaBot/sdk/websocket"
	"context"

	"time"
)

// RetractMessageOption 撤回消息可选参数
type RetractMessageOption int

const (
	// RetractMessageOptionHidetip 撤回消息隐藏小灰条可选参数
	RetractMessageOptionHidetip RetractMessageOption = 1
)

// OpenAPI openapi 完整实现
type OpenAPI interface {
	Base
	WebsocketAPI
	UserAPI
	MessageAPI
	AudioAPI
	PinsAPI
	MessageReactionAPI
	WebhookAPI
	InteractionAPI
}

// Base 基础能力接口
type Base interface {
	Version() APIVersion
	Setup(token *entry.Token, inSandbox bool) OpenAPI
	// WithTimeout 设置请求接口超时时间
	WithTimeout(duration time.Duration) OpenAPI
	// Transport 透传请求，如果 sdk 没有及时跟进新的接口的变更，可以使用该方法进行透传，openapi 实现时可以按需选择是否实现该接口
	Transport(ctx context.Context, method, url string, body interface{}) ([]byte, error)
	// TraceID 返回上一次请求的 trace id
	TraceID() string
}

// WebsocketAPI websocket 接入地址
type WebsocketAPI interface {
	WS(ctx context.Context, params map[string]string, body string) (*websocket.WebsocketAP, error)
}

// UserAPI 用户相关接口
type UserAPI interface {
	Me(ctx context.Context) (*entry.User, error)
	//MeGuilds(ctx context.Context, pager *entry.GuildPager) ([]*entry.Guild, error)
}

// MessageAPI 消息相关接口
type MessageAPI interface {
	Message(ctx context.Context, channelID string, messageID string) (*entry.Message, error)
	Messages(ctx context.Context, channelID string, pager *util.MessagesPager) ([]*entry.Message, error)
	PostMessage(ctx context.Context, channelID string, msg *entry.MessageToCreate) (*entry.Message, error)
	PatchMessage(ctx context.Context,
		channelID string, messageID string, msg *entry.MessageToCreate) (*entry.Message, error)
	RetractMessage(ctx context.Context, channelID, msgID string, options ...RetractMessageOption) error
	// PostSettingGuide 发送设置引导
	PostSettingGuide(ctx context.Context, channelID string, atUserIDs []string) (*entry.Message, error)
}

// GuildAPI guild 相关接口
//type GuildAPI interface {
//	Guild(ctx context.Context, guildID string) (*entry.Guild, error)
//	GuildMember(ctx context.Context, guildID, userID string) (*entry.Member, error)
//	GuildMembers(ctx context.Context, guildID string, pager *entry.GuildMembersPager) ([]*entry.Member, error)
//	DeleteGuildMember(ctx context.Context, guildID, userID string, opts ...entry.MemberDeleteOption) error
//	// 频道禁言
//	GuildMute(ctx context.Context, guildID string, mute *entry.UpdateGuildMute) error
//}

// ChannelAPI 频道相关接口
//type ChannelAPI interface {
//	// Channel 拉取指定子频道信息
//	Channel(ctx context.Context, channelID string) (*entry.Channel, error)
//	// Channels 拉取子频道列表
//	Channels(ctx context.Context, guildID string) ([]*entry.Channel, error)
//	// PostChannel 创建子频道
//	PostChannel(ctx context.Context, guildID string, value *entry.ChannelValueObject) (*entry.Channel, error)
//	// PatchChannel 修改子频道
//	PatchChannel(ctx context.Context, channelID string, value *entry.ChannelValueObject) (*entry.Channel, error)
//	// DeleteChannel 删除指定子频道
//	DeleteChannel(ctx context.Context, channelID string) error
//	// CreatePrivateChannel 创建私密子频道
//	CreatePrivateChannel(ctx context.Context,
//		guildID string, value *entry.ChannelValueObject, userIds []string) (*entry.Channel, error)
//	ListVoiceChannelMembers(ctx context.Context, channelID string) ([]*entry.Member, error)
//}

// ChannelPermissionsAPI 子频道权限相关接口
//type ChannelPermissionsAPI interface {
//	// ChannelPermissions 获取指定子频道的权限
//	ChannelPermissions(ctx context.Context, channelID, userID string) (*entry.ChannelPermissions, error)
//	// PutChannelPermissions 修改指定子频道的权限
//	PutChannelPermissions(ctx context.Context, channelID, userID string, p *entry.UpdateChannelPermissions) error
//	// ChannelRolesPermissions  获取指定子频道身份组的权限
//	ChannelRolesPermissions(ctx context.Context, channelID, roleID string) (*entry.ChannelRolesPermissions, error)
//	// PutChannelRolesPermissions 修改指定子频道身份组的权限
//	PutChannelRolesPermissions(ctx context.Context, channelID, roleID string, p *entry.UpdateChannelPermissions) error
//}

// AudioAPI 音频接口
type AudioAPI interface {
	// PostAudio 执行音频播放，暂停等操作
	PostAudio(ctx context.Context, channelID string, value *entry.AudioControl) (*entry.AudioControl, error)
}

// RoleAPI 用户组相关接口
//type RoleAPI interface {
//	Roles(ctx context.Context, guildID string) (*entry.GuildRoles, error)
//	PostRole(ctx context.Context, guildID string, role *entry.Role) (*entry.UpdateResult, error)
//	PatchRole(ctx context.Context, guildID string, roleID entry.RoleID, role *entry.Role) (*entry.UpdateResult, error)
//	DeleteRole(ctx context.Context, guildID string, roleID entry.RoleID) error
//}

// MemberAPI 成员相关接口，添加成员到用户组等
//type MemberAPI interface {
//	MemberAddRole(
//		ctx context.Context,
//		guildID string, roleID entry.RoleID, userID string, value *entry.MemberAddRoleBody,
//	) error
//	MemberDeleteRole(
//		ctx context.Context,
//		guildID string, roleID entry.RoleID, userID string, value *entry.MemberAddRoleBody,
//	) error
//	// MemberMute 频道指定单个成员禁言
//	MemberMute(ctx context.Context, guildID, userID string, mute *entry.UpdateGuildMute) error
//	// MultiMemberMute 频道指定批量成员禁言
//	MultiMemberMute(ctx context.Context, guildID string,
//		mute *entry.UpdateGuildMute) (*entry.UpdateGuildMuteResponse, error)
//}

// DirectMessageAPI 信息相关接口
//type DirectMessageAPI interface {
//	// CreateDirectMessage 创建私信频道
//	CreateDirectMessage(ctx context.Context, dm *entry.DirectMessageToCreate) (*entry.DirectMessage, error)
//	// PostDirectMessage 在私信频道内发消息
//	PostDirectMessage(ctx context.Context, dm *entry.DirectMessage, msg *entry.MessageToCreate) (*entry.Message, error)
//	// RetractDMMessage 撤回私信频道消息
//	RetractDMMessage(ctx context.Context, guildID, msgID string, options ...RetractMessageOption) error
//	// PostDMSettingGuide 发送私信设置引导, jumpGuildID为设置引导要跳转的频道ID
//	PostDMSettingGuide(ctx context.Context, dm *entry.DirectMessage, jumpGuildID string) (*entry.Message, error)
//}

// AnnouncesAPI 公告相关接口
//type AnnouncesAPI interface {
//	// CreateChannelAnnounces 创建子频道公告
//	CreateChannelAnnounces(
//		ctx context.Context,
//		channelID string, announce *entry.ChannelAnnouncesToCreate,
//	) (*entry.Announces, error)
//	// DeleteChannelAnnounces 删除子频道公告,会校验 messageID 是否匹配
//	DeleteChannelAnnounces(ctx context.Context, channelID, messageID string) error
//	// CleanChannelAnnounces 删除子频道公告,不校验 messageID
//	CleanChannelAnnounces(ctx context.Context, channelID string) error
//	// CreateGuildAnnounces 创建频道全局公告
//	CreateGuildAnnounces(
//		ctx context.Context, guildID string,
//		announce *entry.GuildAnnouncesToCreate,
//	) (*entry.Announces, error)
//	// DeleteGuildAnnounces 删除频道全局公告
//	DeleteGuildAnnounces(ctx context.Context, guildID, messageID string) error
//	// CleanGuildAnnounces 删除频道全局公告,不校验 messageID
//	CleanGuildAnnounces(ctx context.Context, guildID string) error
//}

// ScheduleAPI 日程相关接口
//type ScheduleAPI interface {
//	// ListSchedules 查询某个子频道下，since开始的当天的日程列表。若since为0，默认返回当天的日程列表
//	ListSchedules(ctx context.Context, channelID string, since uint64) ([]*entry.Schedule, error)
//	// GetSchedule 获取单个日程信息
//	GetSchedule(ctx context.Context, channelID, scheduleID string) (*entry.Schedule, error)
//	// CreateSchedule 创建日程
//	CreateSchedule(ctx context.Context, channelID string, schedule *entry.Schedule) (*entry.Schedule, error)
//	// ModifySchedule 修改日程
//	ModifySchedule(ctx context.Context, channelID, scheduleID string, schedule *entry.Schedule) (*entry.Schedule, error)
//	// DeleteSchedule 删除日程
//	DeleteSchedule(ctx context.Context, channelID, scheduleID string) error
//}

// APIPermissionsAPI api 权限相关接口
//type APIPermissionsAPI interface {
//	// GetAPIPermissions 获取频道可用权限列表
//	GetAPIPermissions(ctx context.Context, guildID string) (*entry.APIPermissions, error)
//	// RequireAPIPermissions 创建频道 API 接口权限授权链接
//	RequireAPIPermissions(ctx context.Context,
//		guildID string, demand *entry.APIPermissionDemanentryCreate) (*entry.APIPermissionDemand, error)
//}

// PinsAPI 精华消息接口
type PinsAPI interface {
	// AddPins 添加精华消息
	AddPins(ctx context.Context, channelID string, messageID string) (*entry.PinsMessage, error)
	// DeletePins 删除精华消息
	DeletePins(ctx context.Context, channelID, messageID string) error
	// CleanPins 清除全部精华消息
	CleanPins(ctx context.Context, channelID string) error
	// GetPins 获取精华消息
	GetPins(ctx context.Context, channelID string) (*entry.PinsMessage, error)
}

// MessageReactionAPI 消息表情表态接口
type MessageReactionAPI interface {
	// CreateMessageReaction 对消息发表表情表态
	CreateMessageReaction(ctx context.Context, channelID, messageID string, emoji entry.Emoji) error
	// DeleteOwnMessageReaction 删除自己的消息表情表态
	DeleteOwnMessageReaction(ctx context.Context, channelID, messageID string, emoji entry.Emoji) error
	// GetMessageReactionUsers 获取消息表情表态用户列表
	GetMessageReactionUsers(ctx context.Context, channelID, messageID string, emoji entry.Emoji,
		pager *util.MessageReactionPager) (*entry.MessageReactionUsers, error)
}

// InteractionAPI 互动接口
type InteractionAPI interface {
	// PutInteraction 更新互动信息
	PutInteraction(ctx context.Context, interactionID string, body string) error
}

// WebhookAPI http 事件网关相关接口
type WebhookAPI interface {
	CreateSession(ctx context.Context, identity HTTPIdentity) (*HTTPReady, error)
	CheckSessions(ctx context.Context) ([]*HTTPSession, error)
	SessionList(ctx context.Context) ([]*HTTPSession, error)
	RemoveSession(ctx context.Context, sessionID string) error
}

// MessageSettingAPI 频道消息设置接口
//type MessageSettingAPI interface {
//	GetMessageSetting(ctx context.Context, guildID string) (*entry.MessageSetting, error)
//}
