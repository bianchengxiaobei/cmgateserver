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

type ServerProtocol struct {
	pool *ProtoMessagePool
}
type MessageHeader struct {
	MsgBodyLen int32
	MessageId  int32
	//OrderId    int32
}

var messageHeaderLen = (int)(unsafe.Sizeof(MessageHeader{}))

func (protocol ServerProtocol) Init() {
	//注册消息
	protocol.pool.Register(1000, reflect.TypeOf(message.C2G_UserLogin{}))
	protocol.pool.Register(1001, reflect.TypeOf(message.G2C_CharacterInfo{}))
	protocol.pool.Register(1002, reflect.TypeOf(message.C2G_SelectCharacter{}))
	protocol.pool.Register(1003, reflect.TypeOf(message.G2C_QuitGame{}))
	protocol.pool.Register(1004, reflect.TypeOf(message.G2C_NeedCloseServer{}))
	protocol.pool.Register(1005,reflect.TypeOf(message.C2G_CreateCharacter{}))


	protocol.pool.Register(100001, reflect.TypeOf(message.C2M_GMCommand{}))
	protocol.pool.Register(100002, reflect.TypeOf(message.M2C_GMCommandResult{}))

	protocol.pool.Register(5001,reflect.TypeOf(message.C2M_ReqRefreshRoomList{}))
	protocol.pool.Register(5003,reflect.TypeOf(message.C2M_CreateRoom{}))
	protocol.pool.Register(5005,reflect.TypeOf(message.C2M_ReqJoinRoom{}))
	protocol.pool.Register(5006,reflect.TypeOf(message.C2M_ReqReady{}))
	protocol.pool.Register(5008,reflect.TypeOf(message.C2M_StartBattle{}))
	protocol.pool.Register(5010,reflect.TypeOf(message.C2M_LoadFinished{}))
	protocol.pool.Register(5013,reflect.TypeOf(message.C2M_Command{}))
	protocol.pool.Register(5014,reflect.TypeOf(message.M2C2M_GamePing{}))
	protocol.pool.Register(5017,reflect.TypeOf(message.C2M_WinBattle{}))
	protocol.pool.Register(5018,reflect.TypeOf(message.C2M_FailedBattle{}))
	protocol.pool.Register(5020,reflect.TypeOf(message.C2M_WatchAds{}))
	protocol.pool.Register(5022,reflect.TypeOf(message.C2M_ChangeNickName{}))
	protocol.pool.Register(5023,reflect.TypeOf(message.C2M_ChangeAvatarIcon{}))
	protocol.pool.Register(5026,reflect.TypeOf(message.C2M2C_ChangeEquipItemPos{}))
	protocol.pool.Register(5027,reflect.TypeOf(message.C2M_QuitRoom{}))
	protocol.pool.Register(5030,reflect.TypeOf(message.C2M_BuyHero{}))
	protocol.pool.Register(5032,reflect.TypeOf(message.C2M2C_Chat{}))
	protocol.pool.Register(5033,reflect.TypeOf(message.C2M_LearnSkill{}))
	protocol.pool.Register(5035,reflect.TypeOf(message.C2M2C_ChangeSkill{}))
	protocol.pool.Register(5036,reflect.TypeOf(message.C2M2C_GetAchievement{}))
	protocol.pool.Register(5037,reflect.TypeOf(message.C2M2C_GetTask{}))
	protocol.pool.Register(5038,reflect.TypeOf(message.C2M2C_GetSign{}))
	protocol.pool.Register(5039,reflect.TypeOf(message.C2M2C_ChangeFreeSoldierData{}))
	protocol.pool.Register(5040,reflect.TypeOf(message.C2M_ChangeSex{}))
	protocol.pool.Register(5041,reflect.TypeOf(message.M2C_ChangeSexResult{}))
	protocol.pool.Register(5042,reflect.TypeOf(message.C2M_ChangeSign{}))
	protocol.pool.Register(5043,reflect.TypeOf(message.M2C_ChangeSignResult{}))
	protocol.pool.Register(5044,reflect.TypeOf(message.C2M_GetBoxAwardItem{}))
	protocol.pool.Register(5045,reflect.TypeOf(message.M2C_GetBoxAwardResult{}))
	protocol.pool.Register(5046,reflect.TypeOf(message.C2M_SellItem{}))
	protocol.pool.Register(5047,reflect.TypeOf(message.M2C_SellItemResult{}))
	protocol.pool.Register(5048,reflect.TypeOf(message.C2M_GetEmailAward{}))
	protocol.pool.Register(5049,reflect.TypeOf(message.M2C_GetEmailResult{}))
	protocol.pool.Register(5050,reflect.TypeOf(message.C2M2C_DeleteEmail{}))
	protocol.pool.Register(5051,reflect.TypeOf(message.M2C_AddEmail{}))
	protocol.pool.Register(5052,reflect.TypeOf(message.C2M_UseItem{}))
	protocol.pool.Register(5053,reflect.TypeOf(message.M2C_UseItemResult{}))
	protocol.pool.Register(5054,reflect.TypeOf(message.C2M2C_ReqPauseBattle{}))
	protocol.pool.Register(5055,reflect.TypeOf(message.C2M_AgreePauseBattle{}))
	protocol.pool.Register(5056,reflect.TypeOf(message.M2C_StartPause{}))
	protocol.pool.Register(5057,reflect.TypeOf(message.C2M_ForceEnterBattle{}))
	protocol.pool.Register(5058,reflect.TypeOf(message.C2M_ReqRankList{}))
	protocol.pool.Register(5059,reflect.TypeOf(message.M2C_RankListResult{}))
	protocol.pool.Register(5060,reflect.TypeOf(message.M2C_RollInfo{}))
	protocol.pool.Register(5061,reflect.TypeOf(message.C2M2C_BangDingZhang{}))
	protocol.pool.Register(5062,reflect.TypeOf(message.C2M_ReqAutoMatch{}))
	protocol.pool.Register(5063,reflect.TypeOf(message.M2C_MatchTeamInfo{}))
	protocol.pool.Register(5064,reflect.TypeOf(message.M2C_RemoveMatchTeam{}))
	protocol.pool.Register(5065,reflect.TypeOf(message.M2C_RemoveMatchPlayerFromMatchTeam{}))
	protocol.pool.Register(5066,reflect.TypeOf(message.C2M_QuitMatchTeam{}))
	protocol.pool.Register(5067,reflect.TypeOf(message.M2C_EnterMatchBattleRoom{}))
	protocol.pool.Register(5068,reflect.TypeOf(message.C2M_TeamStartMatch{}))
	protocol.pool.Register(5069,reflect.TypeOf(message.C2M2C_MatchRoomPrepare{}))
	protocol.pool.Register(5070,reflect.TypeOf(message.M2C_StartPaiWeiGame{}))
	protocol.pool.Register(5071,reflect.TypeOf(message.M2C_MatchBattleRoomEnterPrepare{}))
	protocol.pool.Register(5072,reflect.TypeOf(message.C2M_PaiWeiLoadFinished{}))
	protocol.pool.Register(5073,reflect.TypeOf(message.C2M_CancelStartMatch{}))
	protocol.pool.Register(5074,reflect.TypeOf(message.M2C_RemoveMatchBattleRoom{}))
	protocol.pool.Register(5075,reflect.TypeOf(message.M2C_RePaiWeiBattleConnect{}))
	protocol.pool.Register(5076,reflect.TypeOf(message.C2M_UpdateAchievementData{}))
	protocol.pool.Register(5077,reflect.TypeOf(message.C2M_BuyShopEquip{}))
	protocol.pool.Register(5078,reflect.TypeOf(message.M2C_BuyShopEquipResult{}))
	protocol.pool.Register(5079,reflect.TypeOf(message.M2C_BuyCardResult{}))
	protocol.pool.Register(5080,reflect.TypeOf(message.C2M_BuyShopDiam{}))
	protocol.pool.Register(5081,reflect.TypeOf(message.M2C_BuyDiamResult{}))
	protocol.pool.Register(5082,reflect.TypeOf(message.C2M_CheckBuyShopDiam{}))
	protocol.pool.Register(5083,reflect.TypeOf(message.C2M_BuyShopCard{}))
	protocol.pool.Register(5084,reflect.TypeOf(message.C2M_CheckBuyShopCard{}))
	protocol.pool.Register(5085,reflect.TypeOf(message.C2M_CompleteGuide{}))
	protocol.pool.Register(5086,reflect.TypeOf(message.C2M_BuyShopBox{}))
	protocol.pool.Register(5087,reflect.TypeOf(message.M2C_BuyShopBoxResult{}))
	protocol.pool.Register(5088,reflect.TypeOf(message.C2M_BuyHeroCard{}))
	protocol.pool.Register(5089,reflect.TypeOf(message.M2C_BuyHeroCardResult{}))
	protocol.pool.Register(5090,reflect.TypeOf(message.C2M_UpgradeHeroLevel{}))
	protocol.pool.Register(5091,reflect.TypeOf(message.M2C_UpgradeHeroLevelResult{}))
	protocol.pool.Register(5092,reflect.TypeOf(message.C2M_GetAllNoReadEmail{}))
	protocol.pool.Register(5093,reflect.TypeOf(message.M2C_GetAllNoReadEmailResult{}))
	protocol.pool.Register(5094,reflect.TypeOf(message.C2M2C_DeleteAllReadEmail{}))
	protocol.pool.Register(5095,reflect.TypeOf(message.M2C_CardAward{}))
	protocol.pool.Register(5096,reflect.TypeOf(message.C2M2C_DeleteBagItem{}))
	protocol.pool.Register(5097,reflect.TypeOf(message.C2M_GetGift{}))
	protocol.pool.Register(5098,reflect.TypeOf(message.M2C_GetGiftResult{}))
	protocol.pool.Register(5099,reflect.TypeOf(message.C2M2C_ChangeRoomCityId{}))
	protocol.pool.Register(5100,reflect.TypeOf(message.C2M2C_ReStartPauseBattle{}))
	protocol.pool.Register(5101,reflect.TypeOf(message.C2M_ReqBattleBugData{}))
	protocol.pool.Register(5102,reflect.TypeOf(message.M2C_ReBattleBugData{}))
	protocol.pool.Register(5103,reflect.TypeOf(message.C2M_EnterBattleState{}))
	protocol.pool.Register(5104,reflect.TypeOf(message.C2M_InviteRoom{}))
	protocol.pool.Register(5105,reflect.TypeOf(message.M2C_InviteRoomResult{}))
	protocol.pool.Register(5106,reflect.TypeOf(message.C2M_OnlinePlayer{}))
	protocol.pool.Register(5107,reflect.TypeOf(message.M2C_OnlinePlayerResult{}))
	protocol.pool.Register(5108,reflect.TypeOf(message.M2C_BattleFinished{}))
	protocol.pool.Register(5109,reflect.TypeOf(message.C2M_ChangePassword{}))
	protocol.pool.Register(5110,reflect.TypeOf(message.M2C_ChangePasswordResult{}))
	protocol.pool.Register(5111,reflect.TypeOf(message.C2M_CheckOnlineGetDiam{}))
	protocol.pool.Register(5112,reflect.TypeOf(message.M2C_CheckOnlineGetDiam{}))
}
func (protocol ServerProtocol) Decode(session network.SocketSessionInterface, data []byte) (interface{}, int, error) {
	var (
		err       error
		ioBuffer  *bytes.Buffer
		msgHeader MessageHeader
		chanMsg   network.WriteMessage
	)
	//log4g.Infof("fwef:%d",len(data))
	//if len(data) < 256{
	//	return nil, 0, nil
	//}
	msgHeader = MessageHeader{}
	ioBuffer = bytes.NewBuffer(data)
	if ioBuffer.Len() < messageHeaderLen {
		return nil, 0, nil
	}
	err = binary.Read(ioBuffer, binary.LittleEndian, &msgHeader)
	if err != nil {
		return nil, 0, err
	}
	if ioBuffer.Len() < int(msgHeader.MsgBodyLen) {
		return nil, 0, nil
	}
	bodyLen := int(msgHeader.MsgBodyLen)
	allLen := bodyLen + messageHeaderLen
	//var perOrder = session.GetAttribute(network.PREORDERID)
	//if perOrder == nil {
	//	session.SetAttribute(network.PREORDERID, msgHeader.OrderId+1)
	//	//if msgHeader.OrderId == 0 {
	//	//	fmt.Println("用户客户端发送消息序列成功")
	//	//}
	//	//log4g.Infof("fe242f%d,%d,%d",msgHeader.MessageId,msgHeader.MsgBodyLen,msgHeader.OrderId)
	//} else {
	//	if msgHeader.OrderId == perOrder {
	//		//log4g.Infof("fef:%d",msgHeader.MessageId)
	//		session.SetAttribute(network.PREORDERID, msgHeader.OrderId+1)
	//	} else {
	//		log4g.Infof("发送消息[%d]序列出错[%d]",msgHeader.MessageId, msgHeader.OrderId)
	//		return nil, 0, errors.New("发送消息序列出错")
	//	}
	//}
	var msgType = protocol.pool.GetMessageType(msgHeader.MessageId)
	if msgType == nil{
		return nil, 0, errors.New("不存在该消息类型")
	}
	msg := reflect.New(msgType).Interface()
	bodyBytes := ioBuffer.Next(bodyLen)
	err = proto.Unmarshal(bodyBytes, msg.(proto.Message))
	if err != nil {
		log4g.Error(err.Error())
	}
	chanMsg = network.WriteMessage{
		MsgId:   int(msgHeader.MessageId),
		MsgData: msg,
	}
	return chanMsg, allLen, err
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
	defer func() {

	}()
	msg, ok = writeMsg.(network.WriteMessage)
	if ok == false {
		return NotWriteMessage
	}

	msgHeader = MessageHeader{}
	msgHeader.MessageId = int32(msg.MsgId)
	//msgHeader.OrderId = 0

	protoMsg, ok = msg.MsgData.(proto.Message)
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
