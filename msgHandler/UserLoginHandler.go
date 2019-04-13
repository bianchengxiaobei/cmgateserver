package msgHandler

import (
	"cmgateserver/bean"
	"cmgateserver/message"
	"github.com/bianchengxiaobei/cmgo/log4g"
	"github.com/bianchengxiaobei/cmgo/network"
	"github.com/bianchengxiaobei/cmgo/tool"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"cmgateserver/face"
	"time"
)

type UserLoginHandler struct {
	GateServer face.IGateServer
	idGen      *tool.Generator
}

func (handler *UserLoginHandler) Action(session network.SocketSessionInterface, msg interface{}) {
	if protoMsg, ok := msg.(*message.C2G_UserLogin); ok {
		var err error
		//去数据库判断是否存在，不存在直接存数据库
		//设置登录状态
		//网关注册玩家通信列表
		dbSession := handler.GateServer.GetDBManager().Get()
		user := bean.User{}
		if dbSession != nil {
			c := dbSession.DB("sanguozhizhan").C("User")
			err = c.Find(bson.M{"username": protoMsg.UserName}).One(&user)
			if err != nil {
				//说明没找到，就直接创建
				user.UserName = protoMsg.UserName
				user.ServerId = protoMsg.GameServerId
				user.Password = protoMsg.Password
				if handler.idGen == nil {
					handler.idGen, err = tool.NewGenerator(int64(handler.GateServer.GetId()))
				}
				user.UserId = handler.idGen.GetId()
				err = c.Insert(&user)
				if err != nil {
					log4g.Error("玩家账号插入数据库出错!")
					return
				}
			}
		} else {
			log4g.Error("DBManager没有初始化!")
			return
		}
		//设置登录状态
		//如果已经登录的话，就替换老的session
		oldSession := handler.GateServer.GetUserSession(user.UserId)
		if oldSession != nil && oldSession.Id() != session.Id() {
			log4g.Infof("玩家UserId[%d]IP[%s]被踢下线，另一个玩家IP[%s]上线", user.UserId, oldSession.RemoteAddr(), session.RemoteAddr())
			//发送给老session顶替下线的消息
			//然后移除网关服务器里面的玩家通信和角色通信
			aRoleId := oldSession.GetAttribute(network.ROLEID)
			if aRoleId != nil {
				if oldRoleId, ok := aRoleId.(int64); ok {
					handler.GateServer.RemoveRoleSession(oldRoleId)
				}
				oldSession.RemoveAttribute(network.ROLEID)
			}
			aUserId := oldSession.GetAttribute(network.USERID)
			if aUserId != nil {
				if oldUserId, ok := aUserId.(int64); ok {
					handler.GateServer.RemoveUserSession(oldUserId)
				}
				oldSession.RemoveAttribute(network.USERID)
			}
			//关闭oldSession
			oldSession.Close(0)
		}
		handler.GateServer.RegisterUserSession(user.ServerId, user.UserId, session)
		session.SetAttribute(network.USERNAME, user.UserName)
		session.SetAttribute(network.SERVERID, user.ServerId)
		//查询数据库是否已经存在角色
		role := bean.Role{}
		c := dbSession.DB("sanguozhizhan").C("Role")
		err = c.Find(bson.M{"userid": user.UserId}).One(&role)
		if err != nil {
			//说明游戏服务器还没有该玩家的角色，自动创建角色发送给客户端
			if handler.idGen == nil {
				handler.idGen, err = tool.NewGenerator(int64(handler.GateServer.GetId()))
			}
			role.RoleId = handler.idGen.GetId()
			role.UserId = user.UserId
			role.NickName = "Player_" + strconv.Itoa(int(role.RoleId))
			role.ServerId = user.ServerId
			role.Level = 1
			role.Exp = 0
			role.Gold = 0
			role.Diam = 0
			role.RankScore = 0
			role.AvatarId = 1
			role.MaxBagNum = 32
			role.Sign = "(ง •_•)ง,加油"
			role.Sex = 0
			now := time.Now()
			role.LoginTime = now
			role.HeroCount = 1
			role.DayGetTask = make([]int32,0)
			role.WinLevel = make([]int32,0)
			role.Achievement = make([]int32,0)
			role.TaskSeed = int32(now.Nanosecond())
			//role.WinLevel = append(role.WinLevel, 1,2)
			for i:=0;i<int(role.MaxBagNum);i++{
				item := bean.Item{}
				item.ItemId = 0
				item.ItemSeed = 0
				item.ItemNum = 0
				role.Items = append(role.Items, item)
			}
			//邮件，新手加入就发邮件
			email := bean.Email{
				EmailId:1,
				EmailTime:time.Now().In(time.FixedZone("CST", 8*60*60)).Unix(),
				Get:false,
			}
			role.Emails = append(role.Emails, email)
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
		}
		msg := new(message.G2C_CharacterInfo)
		if msg.Role == nil {
			msg.Role = new(message.Role)
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
		log4g.Infof("游戏玩家[%s]登录游戏服务器[%d]", protoMsg.UserName, protoMsg.GameServerId)
	}

}
