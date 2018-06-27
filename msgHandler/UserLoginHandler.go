package msgHandler

import (
"github.com/bianchengxiaobei/cmgo/network"
"cmgateserver/message"
"github.com/bianchengxiaobei/cmgo/log4g"
	"cmgateserver/bean"
	"gopkg.in/mgo.v2/bson"
	"github.com/bianchengxiaobei/cmgo/tool"
)

type UserLoginHandler struct {
	GateServer IGateServer
	idGen	*tool.Generator
}

func (handler *UserLoginHandler) Action(session network.SocketSessionInterface,msg interface{}) {
	if protoMsg,ok := msg.(*message.C2G_UserLogin);ok{
		var err error
		//去数据库判断是否存在，不存在直接存数据库
		//设置登录状态
		//网关注册玩家通信列表
		dbSession := handler.GateServer.GetDBManager().Get()
		user := bean.User{}
		if dbSession != nil{
			c := dbSession.DB("sanguozhizhan").C("User")
			err = c.Find(bson.M{"username":protoMsg.UserName}).One(&user)
			if err != nil{
				//说明没找到，就直接创建
				user.UserName = protoMsg.UserName
				user.ServerId = protoMsg.GameServerId
				user.Password = protoMsg.Password
				if handler.idGen == nil{
					handler.idGen,err = tool.NewGenerator(int64(handler.GateServer.GetId()))
				}
				user.UserId = handler.idGen.GetId()
				err = c.Insert(&user)
				if err != nil{
					log4g.Error("玩家账号插入数据库出错!")
					return
				}
			}
		}else{
			log4g.Error("DBManager没有初始化!")
			return
		}
		//设置登录状态
		handler.GateServer.RegisterUser(user.ServerId,user.UserId,session)
		log4g.Infof("游戏玩家[%s]登录游戏服务器[%d]",protoMsg.UserName,protoMsg.GameServerId)
	}
}
