package manager

import (
	"github.com/Yyjccc/qqbotsdk/entry"
	"github.com/Yyjccc/qqbotsdk/util"
	"github.com/Yyjccc/qqbotsdk/websocket"
	"math"
	"time"
)

// defaultSessionManager 默认实现的 session manager 为单机版本
// 如果业务要自行实现分布式的 session 管理，则实现 SessionManger 后替换掉 defaultSessionManager
var DefaultSessionManager SessionManager = New()

// SessionManager 接口，管理session
type SessionManager interface {
	// Start 启动连接，默认使用 apInfo 中的 shards 作为 shard 数量，如果有需要自己指定 shard 数，请修 apInfo 中的信息
	Start(apInfo *websocket.WebsocketAP, token *entry.Token, intents *websocket.Intent) error
}

// CanNotResumeErrSet 不能进行 resume 操作的错误码
var CanNotResumeErrSet = map[int]bool{
	util.CodeConnCloseCantResume: true,
}

// CanNotIdentifyErrSet 不能进行 identify 操作的错误码
var CanNotIdentifyErrSet = map[int]bool{
	util.CodeConnCloseCantIdentify: true,
}

// concurrencyTimeWindowSec 并发时间窗口，单位秒
const concurrencyTimeWindowSec = 2

// CalcInterval 根据并发要求，计算连接启动间隔
func CalcInterval(maxConcurrency uint32) time.Duration {
	if maxConcurrency == 0 {
		maxConcurrency = 1
	}
	f := math.Round(concurrencyTimeWindowSec / float64(maxConcurrency))
	if f == 0 {
		f = 1
	}
	return time.Duration(f) * time.Second
}

// CanNotResume 是否是不能够 resume 的错误
func CanNotResume(err error) bool {
	e := util.Error(err)
	if flag, ok := CanNotResumeErrSet[e.Code()]; ok {
		return flag
	}
	return false
}

// CanNotIdentify 是否是不能够 identify 的错误
func CanNotIdentify(err error) bool {
	e := util.Error(err)
	if flag, ok := CanNotIdentifyErrSet[e.Code()]; ok {
		return flag
	}
	return false
}

// CheckSessionLimit 检查链接数是否达到限制，如果达到限制需要等待重置
func CheckSessionLimit(apInfo *websocket.WebsocketAP) error {
	if apInfo.Shards > apInfo.SessionStartLimit.Remaining {
		return util.ErrSessionLimit
	}
	return nil
}
