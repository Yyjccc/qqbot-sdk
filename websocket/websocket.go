package websocket

import (
	"github.com/Yyjccc/qqbotsdk/util"
	"runtime"
	"syscall"
)

var (
	// ClientImpl websocket 实现
	ClientImpl WebSocket
	// ResumeSignal 用于强制 resume 连接的信号量
	ResumeSignal syscall.Signal
)

// WebSocket 需要实现的接口
type WebSocket interface {
	// New 创建一个新的ws实例，需要传递 session 对象
	New(session Session) WebSocket
	// Connect 连接到 wss 地址
	Connect() error
	// Identify 鉴权连接
	Identify() error
	// Session 拉取 session 信息，包括 token，shard，seq 等
	Session() *Session
	// Resume 重连
	Resume() error
	// Listening 监听websocket事件
	Listening() error
	// Write 发送数据
	Write(message *WSPayload) error
	// Close 关闭连接
	Close()
}

// Register 注册 websocket 实现
func RegisterWS(ws WebSocket) {
	ClientImpl = ws
}

// RegisterResumeSignal 注册用于通知 client 将连接进行 resume 的信号
func RegisterResumeSignal(signal syscall.Signal) {
	ResumeSignal = signal
}

// PanicBufLen Panic 堆栈大小
var PanicBufLen = 1024

// PanicHandler 处理websocket场景的 panic ，打印堆栈
func PanicHandler(e interface{}, session *Session) {
	buf := make([]byte, PanicBufLen)
	buf = buf[:runtime.Stack(buf, false)]
	util.Errorf("[PANIC]%s\n%v\n%s\n", session, e, buf)
}

// RegisterHandlers 兼容老版本的注册方式
func RegisterHandlers(handlers ...interface{}) Intent {
	return RegisterHandlers(handlers...)
}
