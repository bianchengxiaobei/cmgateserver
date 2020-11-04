package msgHandler

import (
	"cmgateserver/face"
	"cmgateserver/message"
	"github.com/bianchengxiaobei/cmgo/log4g"
	"github.com/bianchengxiaobei/cmgo/network"
	"cmgateserver/bean"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"gopkg.in/mgo.v2"
)

type BindZhangHanHandler struct {
	GateServer face.IGateServer
}

func (handler *BindZhangHanHandler) Action(session network.SocketSessionInterface, msg interface{}) {
	if innerMsg, ok := msg.(network.InnerWriteMessage); ok {
		if protoMsg, ok := innerMsg.MsgData.(*message.M2G_BindZhangHao); ok {
			var c *mgo.Collection
			userId := protoMsg.UserId
			dbSession := handler.GateServer.GetDBManager().Get()
			user := bean.User{}
			if dbSession != nil{
				c = dbSession.DB("sanguozhizhan").C("User")
				err := c.Find(bson.M{"userid": userId}).One(&user)
				if err != nil{
					log4g.Infof("找不到User[%d]",userId)
					return
				}
			}else {
				log4g.Info("dbSession == null")
				return
			}
			if protoMsg.ZhangHaoType == message.ZhangHaoType_EEmail{
				user.MailAddress = protoMsg.Value
			}else if protoMsg.ZhangHaoType == message.ZhangHaoType_Phone{
				num,_ := strconv.Atoi(protoMsg.Value)
				user.Phone = int32(num)
			}else if protoMsg.ZhangHaoType == message.ZhangHaoType_QQ{
				num,_ := strconv.Atoi(protoMsg.Value)
				user.QQ = int32(num)
			}else if protoMsg.ZhangHaoType == message.ZhangHaoType_WeiXin{
				user.WeiXin =protoMsg.Value
			}
			err := c.Update(bson.M{"userid": userId},user)
			if err != nil{
				log4g.Error("更新出错!")
			}
		} else {
			log4g.Error("NO M2G_BindZhangHao！")
		}
	}
}
