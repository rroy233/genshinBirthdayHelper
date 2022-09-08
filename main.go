package main

import (
	"encoding/json"
	"errors"
	"genshinBirthdayHelper/helper"
	"io/ioutil"
	"log"
	"os"
)

type ConfigStruct struct {
	Accounts []helper.Account `json:"accounts"`
}

var config *ConfigStruct

func main() {
	//读取配置
	err := loadConfig()
	if err != nil {
		log.Fatalln("读取配置发生错误：", err)
	}

	for i, account := range config.Accounts {
		//初始化
		h, err := helper.New(account)
		if err != nil {
			log.Fatalf("【账号-%d】初始化失败：%s", i, err.Error())
		}
		err = h.Do()
		if err != nil {
			log.Fatalf("【账号-%d】执行失败：%s", i, err.Error())
		}
	}

}

func loadConfig() error {
	//检查config.json是否存在
	_, err := os.Open("./config.json")
	if err != nil {
		if os.IsNotExist(err) == true {
			ioutil.WriteFile("./config.json", []byte("{\n\t\"accounts\": [\n\t\t{\n\t\t\t\"server\": \"cn_gf01\",\n\t\t\t\"uid\": \"\",\n\t\t\t\"mys-cookie\": \"\"\n\t\t}\n\t]\n}"), 0755)
			return errors.New("配置文件不存在，已自动创建，请手动填入cookie。")
		}
	}

	data, err := ioutil.ReadFile("./config.json")
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
