package main

import (
	"context"
	"github.com/Yyjccc/qqbotsdk/config"
	"github.com/Yyjccc/qqbotsdk/entry"
	"github.com/Yyjccc/qqbotsdk/manager"
	"github.com/Yyjccc/qqbotsdk/openapi"
	v2 "github.com/Yyjccc/qqbotsdk/openapi/v2"
	"github.com/Yyjccc/qqbotsdk/send"
	"github.com/Yyjccc/qqbotsdk/util"

	"github.com/Yyjccc/qqbotsdk/websocket"
	"log"
	"time"
)

func main() {
	v2.Setup()
	websocket.Setup()
	token := entry.BotToken(config.APPID, config.APP_TOKEN)
	api := NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
	ctx := context.Background()
	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Printf("%+v, err:%v", ws, err)
	}
	intent := websocket.RegisterHandlers( //DirectMessageHandler(),
		AloneMessage(ctx, api),
	)

	// 启动 session manager 进行 ws 连接的管理，如果接口返回需要启动多个 shard 的连接，这里也会自动启动多个
	e := NewSessionManager().Start(ws, token, &intent)
	if e != nil {
		log.Printf(e.Error())
	}

}

func AloneMessage(ctx context.Context, api openapi.OpenAPI) websocket.MessageHandler {
	return func(event *websocket.WSPayload, data *websocket.WSMessageData) error {
		wrapper := &send.RawMessageWrapper{
			Payload: event,
			Data:    data,
		}
		url := "https://img.nga.178.com/attachments/mon_202402/11/mqQ2t-4ea9K1sT3cSk8-hw.jpg"
		//raw, err := api.ReplyTextMessageByRaw(ctx, data.Content, wrapper)
		rae, err := api.ReplyMediaMessageByRae(ctx, url, send.ImageType, wrapper)
		if err != nil {
			util.Errorf(err.Error())
		}
		util.Infof("发送消息id: %v", rae)
		return err
	}
}

func NewSandboxOpenAPI(token *entry.Token) openapi.OpenAPI {
	return openapi.DefaultImpl.Setup(token, true)
}
func NewSessionManager() manager.SessionManager {
	return manager.DefaultSessionManager
}
