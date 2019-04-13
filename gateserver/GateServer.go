package gateserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bianchengxiaobei/cmgo/db"
	"github.com/bianchengxiaobei/cmgo/log4g"
	"github.com/bianchengxiaobei/cmgo/network"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"cmgateserver/face"
	"cmgateserver/roleManager"
)

type GateServer struct {
	RootDirPath           string
	gateServerConfigPath  string
	innerServerConfigPath string
	gateBaseConfigPath    string
	Name                  string
	Id                    int
	GateAddr              string
	InnerAddr             string
	UserClientServer      network.ISocket
	InnerConnectServer    network.ISocket
	IsRunning             bool
	//玩家客户端通信列表，存的是网关与客户端的session
	userSessions map[int64]network.SocketSessionInterface
	//内部游戏服务器通信列表，存的是网关与游戏服务器的session
	gameSessions map[int32]network.SocketSessionInterface
	//玩家游戏角色通信列表，存的是玩家角色和客户端的session
	roleSessions map[int64]network.SocketSessionInterface
	lock         sync.Mutex
	DBManager    *db.MongoBDManager
	RoleManager  face.IRoleManager
}
type GateServerConfig struct {
	Name       string
	Id         int
	SocketAddr string
	InnerAddr  string
}

var (
	serverCodec        ServerProtocol             //服务器编解码
	innerCodec         InnerProtocol              //内部服务器编解码
	serverHandler      ServerMessageHandler       //服务器消息处理器
	innerHandler       InnnerServerMessageHandler //内部服务器消息处理器
	NoGameSessionError error                      = errors.New("没有逻辑服游戏Session节点")
	NoRoleSessionError error                      = errors.New("玩家还未登录到游戏服务器")
)

//初始化网关配置
func (server *GateServer) Init(gateBaseConfig string, gateConfig string, innerConfig string) {
	var (
		gateSessionConfig  network.SocketSessionConfig
		innerSessionConfig network.SocketSessionConfig
	)
	server.RootDirPath, _ = os.Getwd()
	server.gateServerConfigPath = filepath.Join(server.RootDirPath, gateConfig)
	server.innerServerConfigPath = filepath.Join(server.RootDirPath, innerConfig)
	server.gateBaseConfigPath = filepath.Join(server.RootDirPath, gateBaseConfig)
	LoadSessionConfig(server.gateServerConfigPath, &gateSessionConfig)
	LoadSessionConfig(server.innerServerConfigPath, &innerSessionConfig)
	server.LoadSessionConfig(server.gateBaseConfigPath)
	//server.UserClientServer = network.NewTcpServer("tcp4", &gateSessionConfig)
	server.InnerConnectServer = network.NewTcpServer("tcp4", &innerSessionConfig)
	server.UserClientServer = network.NewKcpServer(&gateSessionConfig)
	//server.UserClientServer = network.NewP2PKcpServer(&gateSessionConfig)
	//设置编解码
	serverCodec = ServerProtocol{
		pool: &ProtoMessagePool{
			messages: make(map[int32]reflect.Type),
		},
	}
	innerCodec = InnerProtocol{
		pool:&ProtoMessagePool{
			messages:make(map[int32]reflect.Type),
		},
	}
	serverCodec.Init()
	innerCodec.Init()
	server.UserClientServer.SetProtocolCodec(serverCodec)
	server.InnerConnectServer.SetProtocolCodec(innerCodec)
	//设置事件处理器
	serverHandler = ServerMessageHandler{
		server:     server.UserClientServer,
		gateServer: server,
		pool: &HandlerPool{
			handlers: make(map[int32]HandlerBase),
		},
	}
	serverHandler.Init()
	innerHandler = InnnerServerMessageHandler{
		server:     server.InnerConnectServer,
		gateServer: server,
		pool: &HandlerPool{
			handlers: make(map[int32]HandlerBase),
		},
	}
	innerHandler.Init()
	server.UserClientServer.SetMessageHandler(serverHandler)
	server.InnerConnectServer.SetMessageHandler(innerHandler)
	//BD
	server.DBManager = db.NewMongoBD("127.0.0.1", 5)
	server.RoleManager = roleManager.NewRoleManager(server)
}
func (server *GateServer) Run() {
	defer func() {
		if err := recover(); err != nil {
			//log4g.Error("网关服务器监听出错!")
			fmt.Println(err)
			return
		}
	}()
	var (
		err 		error
	)
	if server.IsRunning == false {
		//开始对玩家客户端的监听
		if server.UserClientServer != nil {
			err = server.UserClientServer.Bind(server.GateAddr)
			if err != nil{
				log4g.Error(err.Error())
				return
			}
			log4g.Infof("%s[%s]开始运行!", server.Name, server.GateAddr)
		}
		//开始对内部逻辑服的监听
		if server.InnerConnectServer != nil {
			err = server.InnerConnectServer.Bind(server.InnerAddr)
			if err != nil{
				log4g.Error(err.Error())
				return
			}
			log4g.Infof("%s内部监听开始运行!,端口:[%s]", server.Name, server.InnerAddr)
		}
		server.IsRunning = true
	}
}
func (server *GateServer) Close() {
	if server.IsRunning == true {
		server.InnerConnectServer.Close()
		server.UserClientServer.Close()
		server.IsRunning = false
	}
}
func NewGateServer() *GateServer {
	server := &GateServer{
		IsRunning:    false,
		userSessions: make(map[int64]network.SocketSessionInterface),
		gameSessions: make(map[int32]network.SocketSessionInterface),
		roleSessions: make(map[int64]network.SocketSessionInterface),
	}
	return server
}

//加载json配置
func LoadSessionConfig(filePath string, sessionConfig *network.SocketSessionConfig) {
	var (
		err  error
		file *os.File
		data []byte
	)
	_, err = os.Stat(filePath)
	if err != nil {
		//不存在，新建
		if file, err = os.Create(filePath); err != nil {
			fmt.Println(err)
		}
		config := network.SocketSessionConfig{
			TcpNoDelay:         true,
			TcpKeepAlive:       true,
			TcpKeepAlivePeriod: 3e9,
			TcpReadBuffSize:    1024,
			TcpWriteBuffSize:   1024,
			ReadChanLen:        1024,
			WriteChanLen:       1024,
			PeriodTime:5e9,
		}
		data, err = json.Marshal(config)
		if _, err = file.Write(data); err != nil {
			fmt.Println(err)
		}
	}
	if data == nil || len(data) == 0 {
		file, err = os.Open(filePath)
		if err != nil {
			panic(err)

		}
		defer file.Close()
		data, err = ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
	}
	err = json.Unmarshal(data, sessionConfig)
	if err != nil {
		panic(err)
	}
}
func (server *GateServer) LoadSessionConfig(filePath string) {
	var (
		err    error
		file   *os.File
		data   []byte
		config *GateServerConfig
	)
	defer func() {
		file.Close()
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	_, err = os.Stat(filePath)
	if err != nil {
		//不存在，新建
		if file, err = os.Create(filePath); err != nil {
			fmt.Println(err)
		}
		config = &GateServerConfig{
			Name:       "网关服务器",
			Id:         1,
			SocketAddr: ":8000",
			InnerAddr:  ":8001",
		}
		data, err = json.Marshal(config)
		if _, err = file.Write(data); err != nil {
			fmt.Println(err)
		}
	}
	if data == nil || len(data) == 0 {
		file, err = os.Open(filePath)
		if err != nil {
			panic(err)
		}
		data, err = ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
	}
	if config == nil {
		config = new(GateServerConfig)
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
	server.Id = config.Id
	server.Name = config.Name
	server.GateAddr = config.SocketAddr
	server.InnerAddr = config.InnerAddr
}

//网关注册内部游戏服务器
func (server *GateServer) RegisterInnerGameServer(serverId int32, session network.SocketSessionInterface) {
	server.lock.Lock()
	defer server.lock.Unlock()
	session.SetAttribute(network.SERVERID, serverId)
	_, ok := server.gameSessions[serverId]
	if ok {
		log4g.Errorf("GameSession[%d]已经存在!", serverId)
	} else {
		server.gameSessions[serverId] = session
		log4g.Infof("网关注册游戏服务器id:%d", serverId)
	}
}

//网关移除内部游戏服务器
func (server *GateServer) RemoveInnerGameServer(serverId int32, session network.SocketSessionInterface) {
	server.lock.Lock()
	defer server.lock.Unlock()
	_, ok := server.gameSessions[serverId]
	if !ok {
		log4g.Errorf("GameSession[%d]不存在，移除失败！", serverId)
	} else {
		delete(server.gameSessions, serverId)
		log4g.Infof("网关移除游戏服务器id:%d", serverId)
	}
}
func (server *GateServer) GetId() int {
	return server.Id
}
func (server *GateServer) GetDBManager() *db.MongoBDManager {
	return server.DBManager
}
func(server *GateServer)GetRoleManager() face.IRoleManager{
	return server.RoleManager
}
//注册玩家通信
func (server *GateServer) RegisterUserSession(serverId int32, userId int64, session network.SocketSessionInterface) {
	server.lock.Lock()
	defer server.lock.Unlock()
	server.userSessions[userId] = session
	session.SetAttribute(network.USERID, userId)
	session.SetAttribute(network.SERVERID, serverId)
}

//移除玩家通信
func (server *GateServer) RemoveUserSession(userId int64) {
	server.lock.Lock()
	defer server.lock.Unlock()
	delete(server.userSessions, userId)
}

//取得玩家通信
func (server *GateServer) GetUserSession(userId int64) (session network.SocketSessionInterface) {
	server.lock.Lock()
	defer server.lock.Unlock()
	return server.userSessions[userId]
}

//注册玩家角色通信
func (server *GateServer) RegisterRoleSession(roleId int64, session network.SocketSessionInterface) {
	server.lock.Lock()
	defer server.lock.Unlock()
	server.roleSessions[roleId] = session
	session.SetAttribute(network.ROLEID, roleId)
}

//移除玩家角色通信
func (server *GateServer) RemoveRoleSession(roleId int64) {
	server.lock.Lock()
	defer server.lock.Unlock()
	delete(server.roleSessions, roleId)
}

//取得玩家角色通信
func (server *GateServer) GetRoleSession(roleId int64) (session network.SocketSessionInterface) {
	server.lock.Lock()
	defer server.lock.Unlock()
	return server.roleSessions[roleId]
}

//网关发送消息到游戏逻辑服
func (server *GateServer) SendMsgToGameServer(serverId int32, msgId int, msg proto.Message) error {
	session := server.gameSessions[serverId]
	if session == nil {
		//说明不存在游戏逻辑服节点，还没有注册
		log4g.Infof("未注册游戏逻辑服节点Id[%d]", serverId)
		return NoGameSessionError
	} else {
		innerMsg := network.InnerWriteMessage{
			RoleId:  0,
			MsgData: msg,
		}
		if err := session.WriteMsg(msgId, innerMsg); err != nil {
			return err
		}
	}
	return nil
}
func (server *GateServer) SendMsgToGameServerByRoleId(roleId int64, msgId int, msg proto.Message) error {
	onlineRole := server.RoleManager.GetOnlineRole(roleId)
	if onlineRole != nil {
		session := server.gameSessions[onlineRole.GetServerId()]
		if session == nil {
			//说明不存在游戏逻辑服节点，还没有注册
			log4g.Infof("未注册游戏逻辑服节点Id[%d]", onlineRole.GetServerId())
			return NoGameSessionError
		} else {
			innerMsg := network.InnerWriteMessage{
				RoleId:  roleId,
				MsgData: msg,
			}
			if err := session.WriteMsg(msgId, innerMsg); err != nil {
				return err
			}
		}
		return nil
	}
	return NoRoleSessionError
}

//网关转发消息到玩家客户端
func (server *GateServer) SendMsgToClientByRoleId(roleId int64, msgId int, msg interface{}) error {
	session := server.GetRoleSession(roleId)
	if session == nil {
		log4g.Infof("玩家[%d]还未登录到游戏服务器!消息Id[%d]", roleId,msgId)
		return nil
	} else {
		if err := session.WriteMsg(msgId, msg); err != nil {
			return err
		}
	}
	return nil
}

//网关发送消息给玩家客户端
func (server *GateServer) SendMsgToClient(session network.SocketSessionInterface, msgId int, msg interface{}) error {
	err := session.WriteMsg(msgId, msg)
	return err
}
