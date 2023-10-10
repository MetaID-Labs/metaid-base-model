package major

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type Database struct {
	MongoClientMap map[string]*mongo.Client
}

type MGOConfig struct {
	DsName    string
	Addrs     string
	Timeout   int64
	Database  string
	Username  string
	Password  string
	PoolLimit int
}

var _mongoOnce sync.Once
var DB *Database

func InitMongoDB(configPath string) {
	_mongoOnce.Do(func() {
		DB = &Database{MongoClientMap: make(map[string]*mongo.Client, 0)}
		err := setConnect(configPath)
		if err != nil {
			panic("Mongo init error")
		}
		fmt.Println("Mongodb init done")
	})
}

// read config
func readConfig(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// read json config for mongodb
func readJsonConfig(path string, result interface{}) error {
	if data, err := readConfig(path); err != nil {
		return err
	} else {
		err := json.Unmarshal(data, result)
		if err != nil {
			return err
		}
		return nil
	}
}

// connect setting
func setConnect(configPath string) error {
	mongoConfigs := []MGOConfig{}
	if err := readJsonConfig(configPath, &mongoConfigs); err != nil {
		fmt.Println("ReadJsonConfigError ", err)
		return nil
	}
	for _, mongoConfig := range mongoConfigs {
		fmt.Println("mongodb config address : " + mongoConfig.Addrs)
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoConfig.Timeout)*time.Second)
		defer cancel()
		mongoUrl := "mongodb://" + mongoConfig.Username + ":" + mongoConfig.Password + "@" + mongoConfig.Addrs + "/" + mongoConfig.Database
		var err error
		mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl).SetMaxPoolSize(uint64(mongoConfig.PoolLimit)))
		if err != nil {
			return err
		}
		if DB != nil {
			DB.MongoClientMap[mongoConfig.DsName] = mongoClient
		} else {
			return err
		}
	}
	return nil
}

func GetDBWith(dbName string) (*mongo.Client, error) {
	if DB.MongoClientMap[dbName] != nil {
		return DB.MongoClientMap[dbName], nil
	}
	return nil, errors.New("MongoClientMap is nil")
}
