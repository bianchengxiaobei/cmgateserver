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
	"errors"
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
var NotWriteMessage = errors.New("Message != WriteMsg")
var NotInnerMessage = errors.New("Message != InnerMsg")
var NotProtoMessage = errors.New("Message != ProtoMsg")
var CantFindMsgType = errors.New("找不到该消息类型")
func (protocol InnerProtocol) Init() {
	//注册消息
	protocol.pool.Register(10000, reflect.TypeOf(message.M2G_RegisterGate{}))
	//protocol.pool.Register(10001, reflect.TypeOf(message.G2M_LoginToGameServer{}))
	protocol.pool.Register(10002,reflect.TypeOf(message.M2G_LoginSuccessNotifyGate{}))
	//protocol.pool.Register(10003,reflect.TypeOf(message.G2M_RoleRegisterToGateSuccess{}))
	//protocol.pool.Register(10004,reflect.TypeOf(message.G2M_RoleQuitGameServer{}))

	//直接发给客户端消息
	protocol.pool.Register(5000,reflect.TypeOf(message.M2C_EnterLobby{}))
	protocol.pool.Register(5001,reflect.TypeOf(message.C2M_ReqRefreshRoomList{}))
	protocol.pool.Register(5002,reflect.TypeOf(message.M2C_RefreshRoomList{}))
	protocol.pool.Register(5003,reflect.TypeOf(message.C2M_CreateRoom{}))
	protocol.pool.Register(5004,reflect.TypeOf(message.M2C_JoinRoom{}))
	protocol.pool.Register(5005,reflect.TypeOf(message.C2M_ReqJoinRoom{}))
	protocol.pool.Register(5006,reflect.TypeOf(message.C2M_ReqReady{}))
	protocol.pool.Register(5007,reflect.TypeOf(message.M2C_ReadySuccess{}))
	protocol.pool.Register(5008,reflect.TypeOf(message.C2M_StartBattle{}))
	protocol.pool.Register(5009,reflect.TypeOf(message.M2C_StartBattleLoad{}))
	protocol.pool.Register(5010,reflect.TypeOf(message.C2M_LoadFinished{}))
	protocol.pool.Register(5011,reflect.TypeOf(message.M2C_StartBattle{}))
	protocol.pool.Register(5012,reflect.TypeOf(message.M2C_BattleFrame{}))
	protocol.pool.Register(5013,reflect.TypeOf(message.C2M_Command{}))
	protocol.pool.Register(5014,reflect.TypeOf(message.M2C_GamePing{}))
	protocol.pool.Register(5015,reflect.TypeOf(message.M2C_RoomDelete{}))
	protocol.pool.Register(5016,reflect.TypeOf(message.M2C_RoleQuitRoom{}))
}

func (protocol InnerProtocol) Decode(session network.SocketSessionInterface, data []byte) (interface{}, int, error) {
	var (
		err       error
		ioBuffer  *bytes.Buffer
		msgHeader *InnerMessageHeader
		chanMsg   network.WriteMessage
		innerMsg  network.InnerWriteMessage
	)
	msgHeader = new(InnerMessageHeader)
	ioBuffer = bytes.NewBuffer(data)
	if ioBuffer.Len() < InnerMessageHeaderLen {
		return nil, 0, nil
	}
	err = binary.Read(ioBuffer, binary.LittleEndian, msgHeader)
	if err != nil {
		return nil, 0, err
	}
	if ioBuffer.Len() < int(msgHeader.MsgBodyLen) {
		return nil, 0, nil
	}
	bodyLen := int(msgHeader.MsgBodyLen)
	allLen := bodyLen + InnerMessageHeaderLen
	var msgType = protocol.pool.GetMessageType(msgHeader.MessageId)
	if msgType == nil{
		log4g.Infof("找不到MsgId:%d",msgHeader.MessageId)
		return nil, allLen, CantFindMsgType
	}

	msg := reflect.New(msgType).Interface()
	bodyBytes := ioBuffer.Next(bodyLen)
	err = proto.Unmarshal(bodyBytes, msg.(proto.Message))
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
		return NotWriteMessage
	}
	if innerMsg, ok = msg.MsgData.(network.InnerWriteMessage); !ok {
		return NotInnerMessage
	}
	msgHeader = InnerMessageHeader{}
	msgHeader.MessageId = int32(msg.MsgId)

	msgHeader.RoleId = innerMsg.RoleId

	protoMsg, ok = innerMsg.MsgData.(proto.Message)
	if ok == false {
		return NotProtoMessage
	}
	data, err = proto.Marshal(protoMsg)
	if err != nil {
		return err
	}
	msgHeader.MsgBodyLen = int32(len(data))
	ioBuffer = &bytes.Buffer{}
	err = binary.Write(ioBuffer, binary.LittleEndian, msgHeader)
	if err != nil {
		return err
	}
	_,err = ioBuffer.Write(data)
	if err != nil {
		return err
	}
	err = session.WriteBytes(ioBuffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}
