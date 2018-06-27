package db

import (
	"gopkg.in/mgo.v2"
	"github.com/bianchengxiaobei/cmgo/log4g"
	"sync/atomic"
)

type MongoBDManager struct {
	DbSession *mgo.Session
	DbList    chan *mgo.Session
	used      int32
}

func NewMongoBD(url string,poolSize int)  *MongoBDManager{
	db := MongoBDManager{}
	db.Connect(url)
	db.NewSessionPool(poolSize)
	return &db
}
func (db *MongoBDManager)NewSessionPool(size int){
	if db.DbSession != nil{
		db.DbList = make(chan *mgo.Session,size)
		db.used = 0
	}
}
func (db *MongoBDManager)Connect(url string){
	var err error
	db.DbSession, err = mgo.Dial(url)
	if err != nil{
		log4g.Error(err.Error())
	}
	db.DbSession.SetMode(mgo.Eventual,true)
}
func (db *MongoBDManager)Get() *mgo.Session {
	atomic.AddInt32(&db.used,1)
	select {
	case session,more:=<-db.DbList:
		if session == nil || !more{
			log4g.Error("MongoDBPool已经关闭!")
		}
		return session
	default:
		return db.DbSession.Copy()
	}
}
func (db *MongoBDManager)Put(session *mgo.Session)  {
	atomic.AddInt32(&db.used,-1)
	if session == nil{
		return
	}
	select {
	case db.DbList<-session:
	default:
		//如果db.DbList满了的话，就关闭
		session.Close()
	}
}
func (db *MongoBDManager)Close()  {
	close(db.DbList)
	for session := range db.DbList{
		session.Close()
	}
	db.DbSession.Close()
}
func (db *MongoBDManager)Used() int {
	return int(db.used)
}