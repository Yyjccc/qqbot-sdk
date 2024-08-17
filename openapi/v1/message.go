package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Yyjccc/qqbotsdk/entry"
	"github.com/Yyjccc/qqbotsdk/openapi"
	"github.com/Yyjccc/qqbotsdk/util"
	"github.com/tidwall/gjson"
)

// Message 拉取单条消息
func (o *openAPI) Message(ctx context.Context, channelID string, messageID string) (*entry.Message, error) {
	resp, err := o.request(ctx).
		SetResult(entry.Message{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", messageID).
		Get(o.getURL(messageURI))
	if err != nil {
		return nil, err
	}

	// 兼容处理
	result := resp.Result().(*entry.Message)
	if result.ID == "" {
		body := gjson.Get(resp.String(), "message")
		if err := json.Unmarshal([]byte(body.String()), result); err != nil {
			return nil, err
		}
	}
	return result, nil
}

// Messages 拉取消息列表
func (o *openAPI) Messages(ctx context.Context, channelID string, pager *util.MessagesPager) ([]*entry.Message, error) {
	if pager == nil {
		return nil, util.ErrPagerIsNil
	}
	resp, err := o.request(ctx).
		SetPathParam("channel_id", channelID).
		SetQueryParams(pager.QueryParams()).
		Get(o.getURL(messagesURI))
	if err != nil {
		return nil, err
	}

	messages := make([]*entry.Message, 0)
	if err := json.Unmarshal(resp.Body(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

// PostMessage 发消息
func (o *openAPI) PostMessage(ctx context.Context, channelID string, msg *entry.MessageToCreate) (*entry.Message, error) {
	resp, err := o.request(ctx).
		SetResult(entry.Message{}).
		SetPathParam("channel_id", channelID).
		SetBody(msg).
		Post(o.getURL(messagesURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*entry.Message), nil
}

// PatchMessage 编辑消息
func (o *openAPI) PatchMessage(ctx context.Context,
	channelID string, messageID string, msg *entry.MessageToCreate) (*entry.Message, error) {
	resp, err := o.request(ctx).
		SetResult(entry.Message{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", messageID).
		SetBody(msg).
		Patch(o.getURL(messageURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*entry.Message), nil
}

// RetractMessage 撤回消息
func (o *openAPI) RetractMessage(ctx context.Context,
	channelID, msgID string, options ...openapi.RetractMessageOption) error {
	request := o.request(ctx).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", string(msgID))
	for _, option := range options {
		if option == openapi.RetractMessageOptionHidetip {
			request = request.SetQueryParam("hidetip", "true")
		}
	}
	_, err := request.Delete(o.getURL(messageURI))
	return err
}

// PostSettingGuide 发送设置引导消息, atUserID为要at的用户
func (o *openAPI) PostSettingGuide(ctx context.Context,
	channelID string, atUserIDs []string) (*entry.Message, error) {
	var content string
	for _, userID := range atUserIDs {
		content += fmt.Sprintf("<@%s>", userID)
	}
	msg := &entry.SettingGuideToCreate{
		Content: content,
	}
	resp, err := o.request(ctx).
		SetResult(entry.Message{}).
		SetPathParam("channel_id", channelID).
		SetBody(msg).
		Post(o.getURL(settingGuideURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*entry.Message), nil
}
