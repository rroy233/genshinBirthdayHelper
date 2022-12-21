package main

import (
	"encoding/json"
	"errors"
	"genshinBirthdayHelper/helper"
	"github.com/rroy233/logger"
	"log"
	"os"
)

type ConfigStruct struct {
	Accounts []helper.Account `json:"accounts"`
	Logger   struct {
		Enabled   bool   `json:"enabled"`
		Report    bool   `json:"report"`
		ReportUrl string `json:"reportUrl"`
		QueryKey  string `json:"queryKey"`
	} `json:"logger"`
}

var config *ConfigStruct

func main() {
	//读取配置
	err := loadConfig()
	if err != nil {
		log.Fatalln("读取配置发生错误：", err)
	}

	//日志服务
	logger.New(&logger.Config{
		StdOutput:      true,
		StoreLocalFile: config.Logger.Enabled,
		StoreRemote:    config.Logger.Report,
		RemoteConfig: logger.RemoteConfigStruct{
			RequestUrl: config.Logger.ReportUrl,
			QueryKey:   config.Logger.QueryKey,
		},
	})

	for i, account := range config.Accounts {
		//初始化
		h, err := helper.New(account)
		if err != nil {
			logger.FATAL.Fatalf("【账号-%d】初始化失败：%s", i, err.Error())
		}
		err = h.Do()
		if err != nil {
			logger.FATAL.Fatalf("【账号-%d】执行失败：%s", i, err.Error())
		}
	}

}

func loadConfig() error {
	//检查config.json是否存在
	_, err := os.Open("./config.json")
	if err != nil {
		if os.IsNotExist(err) == true {
			os.WriteFile("./config.json", []byte("{\"accounts\":[{\"server\":\"cn_gf01\",\"uid\":\"\",\"mys-cookie\":\"\"}],\"logger\":{\"enabled\":true,\"report\":false,\"reportUrl\":\"http://127.0.0.1:8990/log\",\"queryKey\":\"?key=\"}}"), 0755)
			return errors.New("配置文件不存在，已自动创建，请手动填入cookie。")
		}
	}

	data, err := os.ReadFile("./config.json")
	if err != nil {
		return err
	}

	config = new(ConfigStruct)
	err = json.Unmarshal(data, config)
	if err != nil {
		return err
	}

	return nil
}
