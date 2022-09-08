package helper

import (
	"log"
)

type Helper struct {
	account   *Account
	Hk4eToken string
	info      *RespInfo
	roles     *RespBirthdayRole
}
type Account struct {
	Server string `json:"server"`
	UID    string `json:"uid"`
	Cookie string `json:"mys-cookie"`
}

// New 初始化账号信息
func New(account Account) (*Helper, error) {
	h := new(Helper)
	h.account = &Account{
		Server: account.Server,
		UID:    account.UID,
		Cookie: account.Cookie,
	}
	err := h.Login()
	if err != nil {
		return &Helper{}, err
	}
	return h, nil
}

func (h *Helper) Do() error {
	//查看今天是否有角色生日
	if err := h.GetBirthdayRole(); err != nil {
		log.Fatalln(err)
	}
	if len(h.roles.Data.Role) == 0 {
		log.Println("今天没有角色生日")
		return nil
	}

	//获取生日贺卡
	if err := h.PostBirthday(); err != nil {
		log.Fatalln(err)
	}
	return nil
}
