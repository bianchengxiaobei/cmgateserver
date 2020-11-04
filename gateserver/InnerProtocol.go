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
	protocol.pool.Register(10005,reflect.TypeOf(message.M2G_RoleQuitGate{}))
	protocol.pool.Register(10006,reflect.TypeOf(message.M2G_CloseSession{}))
	protocol.pool.Register(10007,reflect.TypeOf(message.M2G_BindZhangHao{}))


	protocol.pool.Register(100001, reflect.TypeOf(message.C2M_GMCommand{}))
	protocol.pool.Register(100002, reflect.TypeOf(message.M2C_GMCommandResult{}))

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
	protocol.pool.Register(5014,reflect.TypeOf(message.M2C2M_GamePing{}))
	protocol.pool.Register(5015,reflect.TypeOf(message.M2C_RoomDelete{}))
	protocol.pool.Register(5016,reflect.TypeOf(message.M2C_RoleQuitRoom{}))
	protocol.pool.Register(5017,reflect.TypeOf(message.C2M_WinBattle{}))
	protocol.pool.Register(5018,reflect.TypeOf(message.C2M_FailedBattle{}))
	protocol.pool.Register(5019,reflect.TypeOf(message.M2C_BattleResult{}))
	protocol.pool.Register(5020,reflect.TypeOf(message.C2M_WatchAds{}))
	protocol.pool.Register(5021,reflect.TypeOf(message.M2C_WatchAdsResult{}))
	protocol.pool.Register(5022,reflect.TypeOf(message.C2M_ChangeNickName{}))
	protocol.pool.Register(5023,reflect.TypeOf(message.C2M_ChangeAvatarIcon{}))
	protocol.pool.Register(5024,reflect.TypeOf(message.M2C_ChangeNickName{}))
	protocol.pool.Register(5025,reflect.TypeOf(message.M2C_ChangeAvatarIcon{}))
	protocol.pool.Register(5026,reflect.TypeOf(message.C2M2C_ChangeEquipItemPos{}))
	protocol.pool.Register(5027,reflect.TypeOf(message.C2M_QuitRoom{}))
	protocol.pool.Register(5028,reflect.TypeOf(message.M2C_ReBattleConnect{}))
	protocol.pool.Register(5029,reflect.TypeOf(message.M2C_ReRoomConnect{}))
	protocol.pool.Register(5030,reflect.TypeOf(message.C2M_BuyHero{}))
	protocol.pool.Register(5031,reflect.TypeOf(message.M2C_BuyHeroResult{}))
	protocol.pool.Register(5032,reflect.TypeOf(message.C2M2C_Chat{}))
	protocol.pool.Register(5033,reflect.TypeOf(message.C2M_LearnSkill{}))
	protocol.pool.Register(5034,reflect.TypeOf(message.M2C_LearnSkillResult{}))
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
	if msgHeader.MessageId == 5028{
		log4g.Infof("[%d]",allLen)
	}
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
