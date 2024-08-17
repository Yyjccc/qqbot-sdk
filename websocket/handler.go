package websocket

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

// 维护所有的handler
// DefaultHandlers 默认的 handler 结构，管理所有支持的 handler 类型
var DefaultHandlers struct {
	Ready               ReadyHandler
	ErrorNotify         ErrorNotifyHandler
	Plain               PlainEventHandler
	Message             MessageEventHandler
	MessageReaction     MessageReactionEventHandler
	ATMessage           ATMessageEventHandler
	DirectMessage       DirectMessageEventHandler
	MessageAudit        MessageAuditEventHandler
	MessageDelete       MessageDeleteEventHandler
	PublicMessageDelete PublicMessageDeleteEventHandler
	DirectMessageDelete DirectMessageDeleteEventHandler
	AloneMessage        AloneMessageHandler
	GroupAtMessage      GroupAtMessageHandler
	Audio               AudioEventHandler
	Interaction         InteractionEventHandler
}

// 单聊消息
type AloneMessageHandler func(event *WSPayload, data *WSAloneMessage) error

// 群聊at消息
type GroupAtMessageHandler func(event *WSPayload, data *WSGroupAtMessage) error

// ReadyHandler 可以处理 ws 的 ready 事件
type ReadyHandler func(event *WSPayload, data *WSReadyData)

// ErrorNotifyHandler 当 ws 连接发生错误的时候，会回调，方便使用方监控相关错误
// 比如 reconnect invalidSession 等错误，错误可以转换为 bot.Err
type ErrorNotifyHandler func(err error)

// PlainEventHandler 透传handler
type PlainEventHandler func(event *WSPayload, message []byte) error

// MessageEventHandler 消息事件 handler
type MessageEventHandler func(event *WSPayload, data *WSMessageData) error

// MessageDeleteEventHandler 消息事件 handler
type MessageDeleteEventHandler func(event *WSPayload, data *WSMessageDeleteData) error

// PublicMessageDeleteEventHandler 消息事件 handler
type PublicMessageDeleteEventHandler func(event *WSPayload, data *WSPublicMessageDeleteData) error

// DirectMessageDeleteEventHandler 消息事件 handler
type DirectMessageDeleteEventHandler func(event *WSPayload, data *WSDirectMessageDeleteData) error

// MessageReactionEventHandler 表情表态事件 handler
type MessageReactionEventHandler func(event *WSPayload, data *WSMessageReactionData) error

// ATMessageEventHandler at 机器人消息事件 handler
type ATMessageEventHandler func(event *WSPayload, data *WSATMessageData) error

// DirectMessageEventHandler 私信消息事件 handler
type DirectMessageEventHandler func(event *WSPayload, data *WSDirectMessageData) error

// AudioEventHandler 音频机器人事件 handler
type AudioEventHandler func(event *WSPayload, data *WSAudioData) error

// MessageAuditEventHandler 消息审核事件 handler
type MessageAuditEventHandler func(event *WSPayload, data *WSMessageAuditData) error

//// ThreadEventHandler 论坛主题事件 handler
//type ThreadEventHandler func(event *WSPayload, data *WSThreadData) error
//
//// PostEventHandler 论坛回帖事件 handler
//type PostEventHandler func(event *WSPayload, data *WSPostData) error
//
//// ReplyEventHandler 论坛帖子回复事件 handler
//type ReplyEventHandler func(event *WSPayload, data *WSReplyData) error
//
//// ForumAuditEventHandler 论坛帖子审核事件 handler
//type ForumAuditEventHandler func(event *WSPayload, data *WSForumAuditData) error

// InteractionEventHandler 互动事件 handler
type InteractionEventHandler func(event *WSPayload, data *WSInteractionData) error

type eventParseFunc func(event *WSPayload, message []byte) error

var EventParseFuncMap = map[OPCode]map[EventType]eventParseFunc{
	WSDispatchEvent: {
		//暂且合并
		EventC2cMessageCreate:     messageHandler,
		EventGroupAtMessageCreate: messageHandler,

		EventMessageCreate: messageHandler,
		EventMessageDelete: messageDeleteHandler,

		EventMessageReactionAdd:    messageReactionHandler,
		EventMessageReactionRemove: messageReactionHandler,

		EventAtMessageCreate:     atMessageHandler,
		EventPublicMessageDelete: publicMessageDeleteHandler,

		EventDirectMessageCreate: directMessageHandler,
		EventDirectMessageDelete: directMessageDeleteHandler,

		EventAudioStart:  audioHandler,
		EventAudioFinish: audioHandler,
		EventAudioOnMic:  audioHandler,
		EventAudioOffMic: audioHandler,

		EventMessageAuditPass:   messageAuditHandler,
		EventMessageAuditReject: messageAuditHandler,

		EventInteractionCreate: interactionHandler,
	},
}

// ParseAndHandle 处理回调事件
func ParseAndHandle(payload *WSPayload) error {
	// 指定类型的 handler
	if h, ok := EventParseFuncMap[payload.OPCode][payload.Type]; ok {
		return h(payload, payload.RawMessage)
	}
	// 透传handler，如果未注册具体类型的 handler，会统一投递到这个 handler
	if DefaultHandlers.Plain != nil {
		return DefaultHandlers.Plain(payload, payload.RawMessage)
	}
	return nil
}

// ParseData 解析数据
func ParseData(message []byte, target interface{}) error {
	data := gjson.Get(string(message), "d")
	return json.Unmarshal([]byte(data.String()), target)
}

func messageHandler(payload *WSPayload, message []byte) error {
	data := &WSMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Message != nil {
		return DefaultHandlers.Message(payload, data)
	}
	return nil
}

func aloneMessageHandler(payload *WSPayload, message []byte) error {
	data := &WSAloneMessage{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.AloneMessage != nil {
		return DefaultHandlers.AloneMessage(payload, data)
	}
	return nil
}

func groupAtMessage(payload *WSPayload, message []byte) error {
	data := &WSGroupAtMessage{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupAtMessage != nil {
		return DefaultHandlers.GroupAtMessage(payload, data)
	}
	return nil
}

func messageDeleteHandler(payload *WSPayload, message []byte) error {
	data := &WSMessageDeleteData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.MessageDelete != nil {
		return DefaultHandlers.MessageDelete(payload, data)
	}
	return nil
}

func messageReactionHandler(payload *WSPayload, message []byte) error {
	data := &WSMessageReactionData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.MessageReaction != nil {
		return DefaultHandlers.MessageReaction(payload, data)
	}
	return nil
}

func atMessageHandler(payload *WSPayload, message []byte) error {
	data := &WSATMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.ATMessage != nil {
		return DefaultHandlers.ATMessage(payload, data)
	}
	return nil
}

func publicMessageDeleteHandler(payload *WSPayload, message []byte) error {
	data := &WSPublicMessageDeleteData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.PublicMessageDelete != nil {
		return DefaultHandlers.PublicMessageDelete(payload, data)
	}
	return nil
}

func directMessageHandler(payload *WSPayload, message []byte) error {
	data := &WSDirectMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.DirectMessage != nil {
		return DefaultHandlers.DirectMessage(payload, data)
	}
	return nil
}

func directMessageDeleteHandler(payload *WSPayload, message []byte) error {
	data := &WSDirectMessageDeleteData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.DirectMessageDelete != nil {
		return DefaultHandlers.DirectMessageDelete(payload, data)
	}
	return nil
}

func audioHandler(payload *WSPayload, message []byte) error {
	data := &WSAudioData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Audio != nil {
		return DefaultHandlers.Audio(payload, data)
	}
	return nil
}

func messageAuditHandler(payload *WSPayload, message []byte) error {
	data := &WSMessageAuditData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.MessageAudit != nil {
		return DefaultHandlers.MessageAudit(payload, data)
	}
	return nil
}

func interactionHandler(payload *WSPayload, message []byte) error {
	data := &WSInteractionData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Interaction != nil {
		return DefaultHandlers.Interaction(payload, data)
	}
	return nil
}
