package send

import (
	"github.com/Yyjccc/qqbotsdk/websocket"
)

type ReplyType string
type ContentType int

const (
	TextType     ContentType = 0
	MarkDownType ContentType = 2
	Ark          ContentType = 3
	EmbedType    ContentType = 4
	MediaMsgType ContentType = 7
)

// 回复后消息发送状态
type ReplyStatus struct {
	Id        string `json:"id"`
	Timestamp string `json:"timestamp"`
}

// 发送回复消息
type ReplyBase struct {
	MsgType ContentType         `json:"msg_type"`
	EventId websocket.EventType `json:"event_Id"`
	MsgId   string              `json:"msg_id"`
	MsgSeq  uint32              `json:"msg_seq"`
}

type TextReplyMessage struct {
	ReplyBase
	Content string `json:"content"`
	//Markdown entry.Markdown        `json:"markdown"`
	//Keyboard entry.MessageKeyboard `json:"keyboard"`

}

// 接收消息原始数据封装
type RawMessageWrapper struct {
	Payload *websocket.WSPayload
	Data    *websocket.WSMessageData
	Seq     uint32
}

type MediaType int

const (
	ImageType MediaType = 1 // "png/jpg"
	VideoType MediaType = 2 // "mp4"
	AudioType MediaType = 3 //"silk"
)

// 富媒体 消息
type MediaMessage struct {
	ReplyBase
	Media MediaInfo `json:"media"`
}

// 富媒体(图片、视频、音频) 文件不能直接上传，只能给出url
type Media struct {
	FileType   MediaType `json:"file_type"`
	Url        string    `json:"url"`
	SrvSendMsg bool      `json:"srv_send_msg"` //设置 true 会直接发送消息到目标端，且会占用主动消息频次,回复默认设置false
	//FileData   interface{} //暂未支持！
}

type MediaInfo struct {
	FileInfo string `json:"file_info"`
}

// 富媒体回复发送结果
type MediaReplyStatus struct {
	FileUuid string `json:"file_uuid"`
	FileInfo string `json:"file_info"` //文件信息，用于发消息接口的 media 字段使用
	Ttl      int    `json:"ttl"`       //有效期，表示剩余多少秒到期，到期后 file_info 失效，当等于 0 时，表示可长期使用
	Id       string `json:"id"`        //发送消息的唯一ID，当srv_send_msg设置为true时返回
}
