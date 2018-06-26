package gateserver

import (
	"github.com/bianchengxiaobei/cmgo/network"
	"github.com/bianchengxiaobei/cmgo/log4g"
)
type InnerProtocol struct {

}

func (protocol InnerProtocol) Encode(network.SocketSessionInterface,[]byte)(interface{},int,error){
	return nil,0,nil
}
func (protocol InnerProtocol) Decode(network.SocketConnectInterface,interface{}) error{
	log4g.Info("fefefe222fef")
	return nil
}
