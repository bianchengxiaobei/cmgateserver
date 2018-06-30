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
			}
			aUserId := oldSession.GetAttribute(network.USERID)
			if aUserId != nil {
				if oldUserId, ok := aUserId.(int64); ok {
					handler.GateServer.RemoveUserSession(oldUserId)
				}
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
			err = c.Insert(&role)
			if err != nil {
				log4g.Error("玩家角色账号插入数据库出错!")
				return
			}
		}
		msg := new(message.G2C_CharacterInfo)
		if msg.Role == nil {
			msg.Role = new(message.G2C_CharacterInfo_Role)
		}
		msg.Role.RoleId = role.RoleId
		msg.Role.NickName = role.NickName
		handler.GateServer.SendMsgToClient(session, 1001, msg)
		log4g.Infof("游戏玩家[%s]登录游戏服务器[%d]", protoMsg.UserName, protoMsg.GameServerId)
	}

}
