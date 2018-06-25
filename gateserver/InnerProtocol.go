package gateserver

import (
	"github.com/bianchengxiaobei/cmgo/network"
)
type InnerProtocol struct {

}

func (protocol InnerProtocol) Encode(network.SocketSessionInterface,[]byte)(interface{},int,error){
	return nil,0,nil
}
func (protocol InnerProtocol) Decode(network.SocketConnectInterface,interface{}) error{
	return nil
}
