package gateserver

import "reflect"

type ProtoMessagePool struct {
	messages map[int32]reflect.Type
}

//注册proto消息
func (pool *ProtoMessagePool) Register(msgId int32, msgType reflect.Type) {
	if _, ok := pool.messages[msgId]; ok {
		return
	}
	pool.messages[msgId] = msgType
}

//获取proto消息处理器
func (pool *ProtoMessagePool) GetMessageType(msgId int32) reflect.Type {
	if _, ok := pool.messages[msgId]; !ok {
		return nil
	}
	return pool.messages[msgId]
}
