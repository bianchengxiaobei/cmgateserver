package gateserver

import (
	"github.com/bianchengxiaobei/cmgo/network"
	"bytes"
	"encoding/binary"
	"unsafe"
	"fmt"
	"reflect"
	"github.com/golang/protobuf/proto"
	"cmgateserver/message"
)

const (
	PreOrderId = "pre_orderId"
)
type ServerProtocol struct {
	pool 	*ProtoMessagePool
}
type MessageHeader struct {
	messageId  int
	orderId    int
	msgBodyLen int
}
var messageHeaderLen = (int)(unsafe.Sizeof(MessageHeader{}))
func (protocol ServerProtocol) Init(){
	//注册消息
	protocol.pool.Register(1,reflect.TypeOf(message.Person{}))
}
func (protocol ServerProtocol) Decode(session network.SocketSessionInterface, data []byte)(interface{},int,error){
	 var (
	 	err 		error
	 	ioBuffer 	*bytes.Buffer
	 	msgHeader	MessageHeader
	 	chanMsg		network.WriteMessage
	 )
	 msgHeader = MessageHeader{}
	 ioBuffer = bytes.NewBuffer(data)
	 if ioBuffer.Len() < messageHeaderLen{
	 	return nil,0,nil
	 }
	 allLen := msgHeader.msgBodyLen + messageHeaderLen
	if ioBuffer.Len() < allLen{
		return nil,0,nil
	}
	 err = binary.Read(ioBuffer,binary.LittleEndian,&msgHeader)
	 if err != nil{
	 	return nil,0,err
	 }
	 var perOrder = (session.GetAttribute(PreOrderId)).(int)
	 if msgHeader.orderId == perOrder{
	 	if msgHeader.orderId == 0{
	 		fmt.Println("用户客户端发送消息序列成功")
		}
		session.SetAttribute(PreOrderId,msgHeader.orderId+1)
	 }else {
	 	fmt.Println("发送消息序列出错")
	 }
	 var msgType = protocol.pool.GetMessageType(msgHeader.messageId)
	 msg := reflect.New(msgType.Elem()).Interface()
	 proto.Unmarshal(ioBuffer.Next(msgHeader.msgBodyLen),msg.(proto.Message))
	 chanMsg = network.WriteMessage{
		MsgId:msgHeader.messageId,
		MsgData:msg,
	}
	 return chanMsg,allLen,nil
}
func (protocol ServerProtocol) Encode(session network.SocketSessionInterface,writeMsg interface{}) error{
	var (
		err 			error
		ioBuffer 		*bytes.Buffer
		msgHeader	MessageHeader
		ok 			bool
		msg			network.WriteMessage
		protoMsg 	proto.Message
		data 		[]byte
	)
	msg,ok = writeMsg.(network.WriteMessage)
	if ok == false{
		panic("Message != WriteMsg")
	}

	msgHeader = MessageHeader{}
	msgHeader.messageId = msg.MsgId

	msgHeader.orderId = 0

	protoMsg,ok = msg.MsgData.(proto.Message)
	if ok == false{
		panic("Msg != ProtoMessage")
	}
	data,err = proto.Marshal(protoMsg)
	if err != nil{
		panic("ProtoMessage Marshal Error")
	}
	msgHeader.msgBodyLen = len(data)

	ioBuffer = &bytes.Buffer{}
	err = binary.Write(ioBuffer,binary.LittleEndian,msgHeader)
	if err != nil{
		return err
	}
	ioBuffer.Write(data)
	session.WriteBytes(ioBuffer.Bytes())
	return nil
}
