package qqbot

import (
	"context"
	"github.com/Yyjccc/qqbotsdk/util"
	"github.com/Yyjccc/qqbotsdk/websocket"
	"log"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	token := BotToken(0, "")
	api := NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
	ctx := context.Background()
	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Printf("%+v, err:%v", ws, err)
	}
	intent := websocket.RegisterHandlers( //DirectMessageHandler(),
		AloneMessage(),
	)

	// 启动 session manager 进行 ws 连接的管理，如果接口返回需要启动多个 shard 的连接，这里也会自动启动多个
	e := NewSessionManager().Start(ws, token, &intent)
	if e != nil {
		log.Printf(e.Error())
	}

}

func AloneMessage() websocket.AloneMessageHandler {
	return func(event *websocket.WSPayload, data *websocket.WSMessageData) error {
		util.Infof("收到消息：" + data.Content)
		return nil
	}
}
