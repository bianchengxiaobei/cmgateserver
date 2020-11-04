package msgHandler

import (
	"cmgateserver/face"
	"github.com/bianchengxiaobei/cmgo/network"
	"cmgateserver/message"
	"github.com/bianchengxiaobei/cmgo/tool"
	"time"
	"cmgateserver/bean"
	"github.com/bianchengxiaobei/cmgo/log4g"
)

type CreateCharacterHandler struct {
	GateServer face.IGateServer
	idGen      *tool.Generator
}

func (handler *CreateCharacterHandler) Action(session network.SocketSessionInterface, msg interface{}) {
	if protoMsg, ok := msg.(*message.C2G_CreateCharacter); ok {
		var err error
		userId, _ := session.GetAttribute(network.USERID).(int64)
		//userName, _ := session.GetAttribute(network.USERNAME).(string)
		serverId, _ := session.GetAttribute(network.SERVERID).(int32)
		msg := new(message.G2C_CharacterInfo)
		if msg.Role == nil {
			msg.Role = new(message.Role)
			msg.Role.RoleId = -1
		}
		dbSession := handler.GateServer.GetDBManager().Get()
		if dbSession == nil{
			msg.Role.RoleId = -2
			handler.GateServer.SendMsgToClient(session, 1001, msg)
			return
		}
		c := dbSession.DB("sanguozhizhan").C("Role")
		//说明游戏服务器还没有该玩家的角色，自动创建角色发送给客户端
		role := bean.Role{}
		if handler.idGen == nil {
			handler.idGen, err = tool.NewGenerator(int64(handler.GateServer.GetId()))
		}
		role.RoleId = handler.idGen.GetId()
		role.UserId = userId
		role.NickName = protoMsg.NickName
		role.ServerId = serverId
		role.Level = 1
		role.Exp = 0
		role.Gold = 0
		role.Diam = 0
		role.RankScore = 0
		role.AvatarId = 1
		role.MaxBagNum = 64
		role.MaxEmailNum = 10
		role.Sign = "(ง •_•)ง,加油"
		role.Sex = 0
		now := time.Now()
		role.LoginTime = now
		role.HeroCount = 1
		role.DayGetTask = make([]int32, 0)
		role.WinLevel = make([]int32, 0)
		role.Achievement = make([]int32, 0)
		role.TaskSeed = int32(now.Unix())
		//role.WinLevel = append(role.WinLevel, 1,2)
		for i := 0; i < int(role.MaxBagNum); i++ {
			item := bean.Item{}
			item.ItemId = 0
			item.ItemSeed = 0
			item.ItemNum = 0
			role.Items = append(role.Items, item)
		}
		//邮件，新手加入就发邮件
		for i := 0; i < role.MaxEmailNum; i++ {
			email := bean.Email{}
			role.Emails = append(role.Emails, email)
		}
		email := bean.Email{
			EmailIndex: 0,
			Title:      "新手礼包",
			Content:    "新手礼包大放送",
			EmailTime:  time.Now().In(time.FixedZone("CST", 8*60*60)).Unix(),
			Get:        false,
			Valid:true,
		}
		role.Emails[0] = email
		//成就记录,type应该是从枚举 的最后一个开始添加
		achieve := new(bean.Achievement)
		allAchieveLen := int(bean.ConditionTypeEnd)
		achieve.ConditionType = make([]int32, allAchieveLen)
		achieve.ConditionValue = make([]int32, allAchieveLen)
		for i := 0; i < allAchieveLen; i++ {
			achieve.ConditionType[i] = int32(i)
			if i == int(bean.HighestRankLevel) {
				//说明最高段位为1
				achieve.ConditionValue[i] = 1
			} else {
				achieve.ConditionValue[i] = 0
			}
		}
		role.AchieveRecord = *achieve
		//月卡等
		//card := bean.Card{
		//	CardType:bean.NoneCard,
		//	BuyTime:time.Now(),
		//}
		//role.Card = card
		//Arrower
		soldierData1 := bean.FreeSoldierData{}
		soldierData1.PlayerType = 2
		soldierData1.TouKuiId = 1
		soldierData1.WeapId = 5001
		role.FreeSoldierData[0] = soldierData1
		//Daodun
		soldierData2 := bean.FreeSoldierData{}
		soldierData2.PlayerType = 1
		soldierData2.TouKuiId = 1
		soldierData2.WeapId = 4001
		role.FreeSoldierData[1] = soldierData2
		//Spear
		soldierData3 := bean.FreeSoldierData{}
		soldierData3.PlayerType = 7
		soldierData3.TouKuiId = 1
		soldierData3.WeapId = 8001
		role.FreeSoldierData[2] = soldierData3
		//fashi
		soldierData4 := bean.FreeSoldierData{}
		soldierData4.PlayerType = 10
		soldierData4.TouKuiId = 1
		soldierData4.WeapId = 7001
		role.FreeSoldierData[3] = soldierData4
		err = c.Insert(&role)
		if err != nil {
			log4g.Error("玩家角色账号插入数据库出错!")
			return
		}
		msg.Role.RoleId = role.RoleId
		msg.Role.NickName = role.NickName
		msg.Role.AvatarId = role.AvatarId
		msg.Role.Level = role.Level
		msg.Role.Diam = role.Diam
		msg.Role.Gold = role.Gold
		msg.Role.Exp = role.Exp
		msg.Role.MaxBagNum = role.MaxBagNum
		handler.GateServer.SendMsgToClient(session, 1001, msg)
	}
}
