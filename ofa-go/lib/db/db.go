package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var (
	dbCfg     *dbConfig
	DataGroup map[string]*xorm.EngineGroup
)

type dbGroupConfig struct {
	OpenConns       int `yaml:"openConns"`
	IdleConns       int `yaml:"idleConns"`
	ConnMaxLifetime int `yaml:"maxLifetime"`
	Prefix 			string `yaml:"prefix"`

	Master string   `yaml:"master"`
	Slaves []string `yaml:"slaves"`
}
type dbConfig struct {
	Adapter string                    `yaml:"adapter"`
	Db      map[string]*dbGroupConfig `yaml:"db"`
}

func initDataGroup() map[string]*xorm.EngineGroup {
	var groups = make(map[string]*xorm.EngineGroup)
	if dbCfg == nil {
		fmt.Println("db config setting error")
	}
	for g, e := range dbCfg.Db {
		dataSourceSlice := make([]string, 0)
		dataSourceSlice = append(dataSourceSlice, e.Master)
		for _, sn := range dbCfg.Db[g].Slaves {
			dataSourceSlice = append(dataSourceSlice, sn)
		}
		if len(dataSourceSlice) > 0 {
			group, err := xorm.NewEngineGroup(dbCfg.Adapter, dataSourceSlice)
			if err != nil {
				fmt.Println("创建数据组链接错误：" + err.Error())
				return nil
			}
			group.SetMaxOpenConns(dbCfg.Db[g].OpenConns)
			group.SetMaxIdleConns(dbCfg.Db[g].IdleConns)
			group.SetConnMaxLifetime(time.Duration(dbCfg.Db[g].ConnMaxLifetime) * time.Second)
			tbMapper := names.NewPrefixMapper(names.SnakeMapper{},dbCfg.Db[g].Prefix)
			group.SetTableMapper(tbMapper)
			//group.SetConnMaxLifetime(5*time.Minute)
			groups[g] = group
			fmt.Println(fmt.Sprintf("%s EngineGroup Opened", g))
		}
	}
	return groups
}

func Use(dbName string) *xorm.EngineGroup {
	if DataGroup == nil {
		DataGroup = initDataGroup()
	}
	g, ok := DataGroup[dbName]
	if !ok {
		fmt.Println(dbName + " - Database does not exist.")
		return nil
	}
	return g
}

func Init(dbCfgFile string) {
	buf, err := ioutil.ReadFile(dbCfgFile)
	if err != nil {
		fmt.Println(dbCfgFile + "文件读取失败")
	}
	err = yaml.Unmarshal(buf, &dbCfg)
	if err != nil {
		fmt.Println(dbCfgFile + "解析失败")
	}
	DataGroup = initDataGroup()
}

func Close() {
	for n, db := range DataGroup {
		_ = db.Close()
		fmt.Println(fmt.Sprintf("%s EngineGroup Closed", n))
	}
}
