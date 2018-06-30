package gateserver

import (
	"github.com/bianchengxiaobei/cmgo/network"
)

type HandlerBase interface {
	Action(session network.SocketSessionInterface, msg interface{})
}

type HandlerPool struct {
	handlers map[int32]HandlerBase
}

//注册事件处理器
func (pool *HandlerPool) Register(msgId int32, handler HandlerBase) {
	if _, ok := pool.handlers[msgId]; ok {
		return
	}
	pool.handlers[msgId] = handler
}

//获取事件处理器
func (pool *HandlerPool) GetHandler(msgId int32) HandlerBase {
	if _, ok := pool.handlers[msgId]; !ok {
		return nil
	}
	return pool.handlers[msgId]
}
