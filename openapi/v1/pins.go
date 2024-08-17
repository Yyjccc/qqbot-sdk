package v1

import (
	"context"
	"github.com/Yyjccc/qqbotsdk/entry"
)

// AddPins 添加精华消息
func (o *openAPI) AddPins(ctx context.Context, channelID string, messageID string) (*entry.PinsMessage, error) {
	resp, err := o.request(ctx).
		SetResult(entry.PinsMessage{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", messageID).
		Put(o.getURL(pinURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*entry.PinsMessage), nil
}

// DeletePins 删除精华消息
func (o *openAPI) DeletePins(ctx context.Context, channelID, messageID string) error {
	_, err := o.request(ctx).
		SetResult(entry.PinsMessage{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", messageID).
		Delete(o.getURL(pinURI))
	return err
}

// GetPins 获取精华消息
func (o *openAPI) GetPins(ctx context.Context, channelID string) (*entry.PinsMessage, error) {
	resp, err := o.request(ctx).
		SetResult(entry.PinsMessage{}).
		SetPathParam("channel_id", channelID).
		Get(o.getURL(pinsURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*entry.PinsMessage), nil
}

// CleanPins 清除全部精华消息
func (o *openAPI) CleanPins(ctx context.Context, channelID string) error {
	_, err := o.request(ctx).
		SetResult(entry.PinsMessage{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", "all").
		Delete(o.getURL(pinURI))
	return err
}
