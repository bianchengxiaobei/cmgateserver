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
	"fmt"
)

type InnerProtocol struct {
	pool *ProtoMessagePool
}
type InnerMessageHeader struct {
	MsgBodyLen int32
	MessageId  int32
	RoleId     int64
}

var InnerMessageHeaderLen = (int)(unsafe.Sizeof(InnerMessageHeader{}))

func (protocol InnerProtocol) Init() {
	//注册消息
	protocol.pool.Register(10000, reflect.TypeOf(message.M2G_RegisterGate{}))
	protocol.pool.Register(10001, reflect.TypeOf(message.G2M_LoginToGameServer{}))
	protocol.pool.Register(10002,reflect.TypeOf(message.M2G_LoginSuccessNotifyGate{}))

	//直接发给客户端消息
	protocol.pool.Register(5000,reflect.TypeOf(message.M2C_EnterLobby{}))
}

func (protocol InnerProtocol) Decode(session network.SocketSessionInterface, data []byte) (interface{}, int, error) {
	var (
		err       error
		ioBuffer  *bytes.Buffer
		msgHeader InnerMessageHeader
		chanMsg   network.WriteMessage
		innerMsg  network.InnerWriteMessage
	)
	msgHeader = InnerMessageHeader{}
	ioBuffer = bytes.NewBuffer(data)
	if ioBuffer.Len() < InnerMessageHeaderLen {
		return nil, 0, nil
	}
	err = binary.Read(ioBuffer, binary.LittleEndian, &msgHeader)
	if err != nil {
		return nil, 0, err
	}
	if ioBuffer.Len() < int(msgHeader.MsgBodyLen) {
		return nil, 0, nil
	}
	allLen := int(msgHeader.MsgBodyLen) + InnerMessageHeaderLen

	var msgType = protocol.pool.GetMessageType(msgHeader.MessageId)
	if msgType == nil{
		fmt.Printf("MsgId:%d",msgHeader.MessageId)
	}

	msg := reflect.New(msgType).Interface()
	err = proto.Unmarshal(ioBuffer.Bytes(), msg.(proto.Message))
	if err != nil {
		log4g.Error(err.Error())
	}
	innerMsg = network.InnerWriteMessage{
		RoleId:  msgHeader.RoleId,
		MsgData: msg,
	}
	chanMsg = network.WriteMessage{
		MsgId:   int(msgHeader.MessageId),
		MsgData: innerMsg,
	}
	return chanMsg, allLen, nil
}
func (protocol InnerProtocol) Encode(session network.SocketSessionInterface, writeMsg interface{}) error {
	var (
		err       error
		ioBuffer  *bytes.Buffer
		msgHeader InnerMessageHeader
		ok        bool
		msg       network.WriteMessage
		innerMsg  network.InnerWriteMessage
		protoMsg  proto.Message
		data      []byte
	)
	msg, ok = writeMsg.(network.WriteMessage)
	if ok == false {
		panic("Message != WriteMsg")
	}
	if innerMsg, ok = msg.MsgData.(network.InnerWriteMessage); !ok {
		panic("Message != InnerMsg")
		return nil
	}
	msgHeader = InnerMessageHeader{}
	msgHeader.MessageId = int32(msg.MsgId)

	msgHeader.RoleId = innerMsg.RoleId

	protoMsg, ok = innerMsg.MsgData.(proto.Message)
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
	err = session.WriteBytes(ioBuffer.Bytes())
	if err != nil {
		log4g.Error(err.Error())
	}
	return nil
}
