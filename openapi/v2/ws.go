package v2

import (
	"context"
	"github.com/Yyjccc/qqbotsdk/websocket"
)

// WS 获取带分片 WSS 接入点
func (o *openAPI) WS(ctx context.Context, _ map[string]string, _ string) (*websocket.WebsocketAP, error) {
	resp, err := o.request(ctx).
		SetResult(websocket.WebsocketAP{}).
		Get(o.getURL(gatewayBotURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*websocket.WebsocketAP), nil
}
