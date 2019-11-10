package model

import (
	"context"
	"decode_test/pkg/app"
	"decode_test/pkg/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

type DBWrapper struct {
	mysqlDB    *gorm.DB
	lock       sync.Mutex
	mongoDB    *mongo.Database
	mongoColls map[string]*mongo.Collection
}

func Setup() *DBWrapper {
	// init mongo mongoDB
	var err error
	var dbWrapper = new(DBWrapper)
	dbWrapper.mongoDB, err = buildMongoDB(config.Cfg.MongoDBConfig)
	if err != nil {
		logrus.Error("init mongo mongoDB error")
		panic(err)
	}
	dbWrapper.mysqlDB, err = buildMysqlDB(config.Cfg.DataBaseConfig)
	if err != nil {
		logrus.Error("init mysql mongoDB error")
		panic(err)
	}
	dbWrapper.mongoColls = make(map[string]*mongo.Collection)
	return dbWrapper
}

func buildMongoDB(cfg config.MongoDBCfg) (*mongo.Database, error) {
	opts := options.Client().
		SetAuth(options.Credential{Username: cfg.Username, Password: cfg.Password, PasswordSet: true}).
		SetHosts(cfg.Addresses).
		SetAppName("decode_test").
		SetMaxPoolSize(uint64(cfg.Pool.MaxSize)).
		SetMinPoolSize(uint64(cfg.Pool.InitSize)).
		SetMaxConnIdleTime(cfg.Pool.IdleTimeOut * time.Millisecond).
		SetReplicaSet(cfg.ReplicaSet).
		SetConnectTimeout(app.DBConnectTimeout).
		SetSocketTimeout(app.DBRWTimeout)
	if cfg.UseSecondaryRead {
		opts.SetReadPreference(readpref.SecondaryPreferred())
	}
	newClient, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}
	err = newClient.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return newClient.Database(cfg.DBName), nil
}

func buildMysqlDB(cfg config.DataBaseCfg) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s:%d)/%s?charset=utf8&timeout=%d&readTimeout=%d&writeTimeout=%d"
	connStr := fmt.Sprintf(s, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, )
	db2, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	db2.DB().SetMaxIdleConns(cfg.Pool.InitSize)
	db2.DB().SetConnMaxLifetime(cfg.Pool.IdleTimeOut)
	db2.DB().SetMaxOpenConns(cfg.Pool.MaxSize)
	return db2, err
}

func (db *DBWrapper)getColl(collName string) *mongo.Collection {
	if db.mongoColls == nil {
		logrus.Error("register mongo collection before model setup")
		panic("register mongo collection before model setup")
	}
	res, ok := db.mongoColls[collName]
	if !ok {
		db.lock.Lock()
		defer db.lock.Unlock()
		res = db.mongoDB.Collection(collName)
		db.mongoColls[collName] = res
	}
	return res
}
