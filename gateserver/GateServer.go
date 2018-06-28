package gateserver

import (
	"encoding/json"
	"fmt"
	"github.com/bianchengxiaobei/cmgo/network"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"github.com/bianchengxiaobei/cmgo/log4g"
	"sync"
	"cmgateserver/db"
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
	UserClientServer      *network.TcpServer
	InnerConnectServer    *network.TcpServer
	IsRunning             bool
	//玩家客户端通信列表，存的是网关与客户端的session
	userSessions map[int64]network.SocketSessionInterface
	//内部游戏服务器通信列表，存的是网关与游戏服务器的session
	gameSessions map[int32]network.SocketSessionInterface
	//玩家游戏角色通信列表，存的是玩家角色和客户端的session
	roleSessions map[int64]network.SocketSessionInterface
	lock sync.Mutex
	DBManger		*db.MongoBDManager
}
type GateServerConfig struct {
	Name       string
	Id         int
	SocketAddr string
	InnerAddr  string
}

var (
	serverCodec   ServerProtocol             //服务器编解码
	innerCodec    InnerProtocol              //内部服务器编解码
	serverHandler ServerMessageHandler       //服务器消息处理器
	innerHandler  InnnerServerMessageHandler //内部服务器消息处理器
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
	server.UserClientServer = network.NewTcpServer("tcp4", &gateSessionConfig)
	server.InnerConnectServer = network.NewTcpServer("tcp4", &innerSessionConfig)
	//设置编解码
	serverCodec = ServerProtocol{
		pool: &ProtoMessagePool{
			messages: make(map[int32]reflect.Type),
		},
	}
	serverCodec.Init()
	server.UserClientServer.SetProtocolCodec(serverCodec)
	server.InnerConnectServer.SetProtocolCodec(serverCodec)
	//设置事件处理器
	serverHandler = ServerMessageHandler{
		server:     server.UserClientServer,
		gateServer: server,
		pool:&HandlerPool{
			handlers: make(map[int32]HandlerBase),
		},
	}
	serverHandler.Init()
	innerHandler = InnnerServerMessageHandler{
		server:     server.InnerConnectServer,
		gateServer: server,
		pool:&HandlerPool{
			handlers:make(map[int32]HandlerBase),
		},
	}
	innerHandler.Init()
	server.UserClientServer.SetMessageHandler(serverHandler)
	server.InnerConnectServer.SetMessageHandler(innerHandler)
	//BD
	server.DBManger = db.NewMongoBD("127.0.0.1",5)
}
func (server *GateServer) Run() {
	defer func() {
		if err := recover(); err != nil {
			//log4g.Error("网关服务器监听出错!")
			fmt.Println(err)
			return
		}
	}()
	if server.IsRunning == false {
		//开始对玩家客户端的监听
		if server.UserClientServer != nil {
			server.UserClientServer.Bind(server.GateAddr)
			log4g.Infof("%s[%s]开始运行!",server.Name,server.GateAddr)
		}
		//开始对内部逻辑服的监听
		if server.InnerConnectServer != nil {
			server.InnerConnectServer.Bind(server.InnerAddr)
			log4g.Infof("%s内部监听开始运行!,端口:[%s]",server.Name,server.InnerAddr)
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
		roleSessions:make(map[int64]network.SocketSessionInterface),
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
			ReadChanLen:        1,
			WriteChanLen:       1,
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
	if data == nil || len(data) == 0{
		file, err = os.Open(filePath)
		if err != nil {
			panic(err)
		}
		data, err = ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
	}
	if config == nil{
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
func (server *GateServer)RegisterInnerGameServer(serverId int32,session network.SocketSessionInterface){
	session.SetAttribute(network.SERVERID,serverId)
	defer server.lock.Unlock()
	server.lock.Lock()
	_,ok :=server.gameSessions[serverId]
	if ok{
		log4g.Errorf("GameSession[%d]已经存在!",serverId)
	}else{
		server.gameSessions[serverId] = session
		log4g.Infof("网关注册游戏服务器id:%d",serverId)
	}
}
//网关移除内部游戏服务器
func (server *GateServer)RemoveInnerGameServer(serverId int32,session network.SocketSessionInterface){
	server.lock.Lock()
	defer  server.lock.Unlock()
	_,ok := server.gameSessions[serverId]
	if !ok{
		log4g.Errorf("GameSession[%d]不存在，移除失败！",serverId)
	}else{
		delete(server.gameSessions, serverId)
		log4g.Infof("网关移除游戏服务器id:%d",serverId)
	}
}
func (server *GateServer)GetId() int{
	return server.Id
}
func (server *GateServer)GetDBManager() *db.MongoBDManager{
	return server.DBManger
}
//注册玩家通信
func (server *GateServer)RegisterUserSession(serverId int32,userId int64,session network.SocketSessionInterface){
	server.lock.Lock()
	defer  server.lock.Unlock()
	server.userSessions[userId] = session
	session.SetAttribute(network.USERID,userId)
	session.SetAttribute(network.SERVERID,serverId)
}
//移除玩家通信
func (server *GateServer)RemoveUserSession(userId int64){
	server.lock.Lock()
	defer server.lock.Unlock()
	delete(server.userSessions, userId)
}
//取得玩家通信
func (server *GateServer)GetUserSession(userId int64) (session network.SocketSessionInterface){
	return server.userSessions[userId]
}
//注册玩家角色通信
func (server *GateServer)RegisterRoleSession(roleId int64,session network.SocketSessionInterface){
	server.lock.Lock()
	defer  server.lock.Unlock()
	server.roleSessions[roleId] = session
	session.SetAttribute(network.ROLEID, roleId)
}
//移除玩家角色通信
func (server *GateServer)RemoveRoleSession(roleId int64){
	server.lock.Lock()
	defer server.lock.Unlock()
	delete(server.roleSessions, roleId)
}
//取得玩家角色通信
func (server *GateServer)GetRoleSession(roleId int64) (session network.SocketSessionInterface){
	return server.roleSessions[roleId]
}