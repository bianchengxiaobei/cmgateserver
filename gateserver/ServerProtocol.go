package gateserver

import (
	"bytes"
	"cmgateserver/message"
	"encoding/binary"
	"github.com/bianchengxiaobei/cmgo/log4g"
	"github.com/bianchengxiaobei/cmgo/network"
	"github.com/golang/protobuf/proto"
	"reflect"
	"unsafe"
)

type ServerProtocol struct {
	pool *ProtoMessagePool
}
type MessageHeader struct {
	MessageId  int32
	OrderId    int32
	MsgBodyLen int32
}

var messageHeaderLen = (int)(unsafe.Sizeof(MessageHeader{}))

func (protocol ServerProtocol) Init() {
	//注册消息
	protocol.pool.Register(10000, reflect.TypeOf(message.M2G_RegisterGate{}))
}
func (protocol ServerProtocol) Decode(session network.SocketSessionInterface, data []byte) (interface{}, int, error) {
	var (
		err       error
		ioBuffer  *bytes.Buffer
		msgHeader MessageHeader
		chanMsg   network.WriteMessage
	)
	msgHeader = MessageHeader{}
	ioBuffer = bytes.NewBuffer(data)
	if ioBuffer.Len() < messageHeaderLen {
		return nil, 0, nil
	}
	allLen := int(msgHeader.MsgBodyLen) + messageHeaderLen
	if ioBuffer.Len() < allLen {
		return nil, 0, nil
	}
	err = binary.Read(ioBuffer, binary.LittleEndian, &msgHeader)
	if err != nil {
		return nil, 0, err
	}
	var perOrder = session.GetAttribute(network.PreOrderId)
	if perOrder == nil {
		session.SetAttribute(network.PreOrderId, msgHeader.OrderId+1)
		//if msgHeader.OrderId == 0 {
		//	fmt.Println("用户客户端发送消息序列成功")
		//}
	} else {
		if msgHeader.OrderId == perOrder {
			session.SetAttribute(network.PreOrderId, msgHeader.OrderId+1)
		} else {
			log4g.Error("发送消息序列出错")
			return nil, 0, nil
		}
	}
	var msgType = protocol.pool.GetMessageType(msgHeader.MessageId)
	msg := reflect.New(msgType).Interface()
	err = proto.Unmarshal(ioBuffer.Next(int(msgHeader.MsgBodyLen)), msg.(proto.Message))
	if err != nil {
		log4g.Error(err.Error())
	}
	chanMsg = network.WriteMessage{
		MsgId:   int(msgHeader.MessageId),
		MsgData: msg,
	}
	return chanMsg, allLen, nil
}
func (protocol ServerProtocol) Encode(session network.SocketSessionInterface, writeMsg interface{}) error {
	var (
		err       error
		ioBuffer  *bytes.Buffer
		msgHeader MessageHeader
		ok        bool
		msg       network.WriteMessage
		protoMsg  proto.Message
		data      []byte
	)
	msg, ok = writeMsg.(network.WriteMessage)
	if ok == false {
		panic("Message != WriteMsg")
	}

	msgHeader = MessageHeader{}
	msgHeader.MessageId = int32(msg.MsgId)

	msgHeader.OrderId = 0

	protoMsg, ok = msg.MsgData.(proto.Message)
	if ok == false {
		panic("Msg != ProtoMessage")
	}
	data, err = proto.Marshal(protoMsg)
	if err != nil {
		panic("ProtoMessage Marshal Error")
	}
	msgHeader.MsgBodyLen = int32(len(data))

	ioBuffer = &bytes.Buffer{}
	err = binary.Write(ioBuffer, binary.LittleEndian, msgHeader)
	if err != nil {
		return err
	}
	ioBuffer.Write(data)
	session.WriteBytes(ioBuffer.Bytes())
	return nil
}
