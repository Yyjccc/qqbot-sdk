package websocket

import "github.com/Yyjccc/qqbotsdk/entry"

// WSPayload websocket 消息结构
type WSPayload struct {
	WSPayloadBase
	Data       interface{} `json:"d,omitempty"`
	RawMessage []byte      `json:"-"` // 原始的 message 数据
}

// WSPayloadBase 基础消息结构，排除了 data
type WSPayloadBase struct {
	OPCode OPCode    `json:"op"`
	Seq    uint32    `json:"s,omitempty"`
	Type   EventType `json:"t,omitempty"`
}

// 以下为发送到 websocket 的 data

// WSIdentityData 鉴权数据
type WSIdentityData struct {
	Token      string   `json:"token"`
	Intents    Intent   `json:"intents"`
	Shard      []uint32 `json:"shard"` // array of two integers (shard_id, num_shards)
	Properties struct {
		Os      string `json:"$os,omitempty"`
		Browser string `json:"$browser,omitempty"`
		Device  string `json:"$device,omitempty"`
	} `json:"properties,omitempty"`
}

// WSResumeData 重连数据
type WSResumeData struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       uint32 `json:"seq"`
}

// 以下为会收到的事件data

// WSHelloData hello 返回
type WSHelloData struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// WSReadyData ready，鉴权后返回
type WSReadyData struct {
	Version   int    `json:"version"`
	SessionID string `json:"session_id"`
	User      struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Bot      bool   `json:"bot"`
	} `json:"user"`
	Shard []uint32 `json:"shard"`
}

//机器人关闭对频道的支持

//// WSGuildData 频道 payload
//type WSGuildData entry.Guild
//
//// WSGuildMemberData 频道成员 payload
//type WSGuildMemberData Member
//
//// WSChannelData 子频道 payload
//type WSChannelData Channel

// 单聊消息 payload
type WSAloneMessage entry.Message

// 群聊at消息
type WSGroupAtMessage entry.Message

// WSMessageData 消息 payload //群聊或者私聊
type WSMessageData entry.Message

// WSATMessageData only at 机器人的消息 payload
type WSATMessageData entry.Message

// WSDirectMessageData 私信消息 payload
type WSDirectMessageData entry.Message

// WSMessageDeleteData 消息 payload
type WSMessageDeleteData entry.MessageDelete

// WSPublicMessageDeleteData 公域机器人的消息删除 payload
type WSPublicMessageDeleteData entry.MessageDelete

// WSDirectMessageDeleteData 私信消息 payload
type WSDirectMessageDeleteData entry.MessageDelete

// WSAudioData 音频机器人的音频流事件
type WSAudioData entry.AudioAction

// WSMessageReactionData 表情表态事件
type WSMessageReactionData entry.MessageReaction

// WSMessageAuditData 消息审核事件
type WSMessageAuditData entry.MessageAudit

//// WSThreadData 主题事件
//type WSThreadData Thread
//
//// WSPostData 帖子事件
//type WSPostData Post
//
//// WSReplyData 帖子回复事件
//type WSReplyData Reply
//
//// WSForumAuditData 帖子审核事件
//type WSForumAuditData ForumAuditResult

// WSInteractionData 互动事件
type WSInteractionData entry.Interaction
