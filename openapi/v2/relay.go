package v2

import (
	"context"
	"github.com/Yyjccc/qqbotsdk/send"
	"github.com/go-resty/resty/v2"
)

// 回复文本消息
func (o *openAPI) ReplyTextMessageByRaw(ctx context.Context, reply string, wrapper *send.RawMessageWrapper) (*send.ReplyStatus, error) {
	payload := wrapper.Payload
	data := wrapper.Data
	var resp *resty.Response
	var err error = nil
	message := send.TextReplyMessage{
		Content: reply,
		ReplyBase: send.ReplyBase{
			MsgType: send.TextType,
			EventId: payload.Type,
			MsgId:   data.ID,
			MsgSeq:  wrapper.Seq,
		},
	}
	if wrapper.Data.GroupOpenID != "" {
		//群聊消息
		resp, err = o.request(ctx).SetResult(send.ReplyStatus{}).
			SetBody(message).
			SetPathParam("group_openid", data.GroupOpenID).
			Post(o.getURL(replyTextGroupURI))
	} else {
		//单聊消息
		resp, err = o.request(ctx).SetResult(send.ReplyStatus{}).
			SetBody(message).
			SetPathParam("openid", data.Author.UserOpenId).
			Post(o.getURL(replyTextAloneURI))
	}
	wrapper.Seq++
	res := resp.Result()
	if res == nil {
		return nil, err
	}
	result := res.(*send.ReplyStatus)
	return result, err

}

// 回复富媒体消息
func (o *openAPI) ReplyMediaMessageByRae(ctx context.Context, url string, mediaType send.MediaType, wrapper *send.RawMessageWrapper) (*send.MediaReplyStatus, error) {
	payload := wrapper.Payload
	data := wrapper.Data
	var resp *resty.Response
	info, err := o.UploadMediaInfo(ctx, url, mediaType, wrapper)
	//上传图片
	if err != nil {
		return nil, err
	}
	message := send.MediaMessage{
		ReplyBase: send.ReplyBase{
			MsgType: send.MediaMsgType,
			EventId: payload.Type,
			MsgId:   data.ID,
			MsgSeq:  wrapper.Seq,
		},
		Media: send.MediaInfo{
			FileInfo: info.FileInfo,
		},
	}
	if wrapper.Data.GroupOpenID != "" {
		//群聊消息
		resp, err = o.request(ctx).SetResult(send.MediaReplyStatus{}).
			SetBody(message).
			SetPathParam("group_openid", data.GroupOpenID).
			Post(o.getURL(replyTextGroupURI))
	} else {
		//单聊消息
		resp, err = o.request(ctx).SetResult(send.MediaReplyStatus{}).
			SetBody(message).
			SetPathParam("openid", data.Author.UserOpenId).
			Post(o.getURL(replyTextAloneURI))
	}
	wrapper.Seq++
	res := resp.Result()
	if res == nil {
		return nil, err
	}
	result := res.(*send.MediaReplyStatus)
	return result, err
}

func (o *openAPI) UploadMediaInfo(ctx context.Context, url string, mediaType send.MediaType, wrapper *send.RawMessageWrapper) (*send.MediaReplyStatus, error) {
	data := wrapper.Data
	var resp *resty.Response
	var err error = nil
	message := send.Media{
		FileType:   mediaType,
		Url:        url,
		SrvSendMsg: false,
	}
	if wrapper.Data.GroupOpenID != "" {
		//群聊消息
		resp, err = o.request(ctx).SetResult(send.MediaReplyStatus{}).
			SetBody(message).
			SetPathParam("group_openid", data.GroupOpenID).
			Post(o.getURL(uploadMediaGroupURI))
	} else {
		//单聊消息
		resp, err = o.request(ctx).SetResult(send.MediaReplyStatus{}).
			SetBody(message).
			SetPathParam("openid", data.Author.UserOpenId).
			Post(o.getURL(uploadMediaAloneURI))
	}
	res := resp.Result()
	if res == nil {
		return nil, err
	}
	result := res.(*send.MediaReplyStatus)
	return result, err
}
