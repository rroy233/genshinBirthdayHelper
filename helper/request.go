package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// urlIndex 和 urlPost 只需e_hk4e_token，需要从 urlApiLogin 中获取
const (
	urlApiLogin = "https://api-takumi.mihoyo.com/common/badge/v1/login/account"
	urlIndex    = "https://hk4e-api.mihoyo.com/event/birthdaystar/account/index?badge_uid=%s&badge_region=%s&game_biz=hk4e_cn&lang=zh-cn&activity_id=20220301153521"
	urlPost     = "https://hk4e-api.mihoyo.com/event/birthdaystar/account/post_my_draw?badge_uid=%s&badge_region=%s&game_biz=hk4e_cn&lang=zh-cn&activity_id=20220301153521"
)

func (h *Helper) Login() error {
	type Payload struct {
		GameBiz string `json:"game_biz"`
		Lang    string `json:"lang"`
		Region  string `json:"region"`
		Uid     string `json:"uid"`
	}
	data := Payload{
		"hk4e_cn",
		"zh-cn",
		h.account.Server,
		h.account.UID,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", urlApiLogin, body)
	if err != nil {
		return err
	}
	req.Host = "api-takumi.mihoyo.com"

	h.setHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respData, _ := ioutil.ReadAll(resp.Body)

	//log.Println(string(respData))
	h.info = new(RespInfo)
	if err = json.Unmarshal(respData, h.info); err != nil {
		return err
	}

	if h.info.Retcode != 0 {
		return errors.New("请求返回错误：" + h.info.Message)
	}

	//提取e_hk4e_token，更新现有cookie中的
	for _, sc := range resp.Header.Values("Set-Cookie") {
		eHk4eToken, _ := MatchSingle(regexp.MustCompile(`e_hk4e_token=(.+?);`), sc)
		if eHk4eToken != "" {
			h.account.Cookie = strings.Replace(h.account.Cookie, " ", "", -1) //去空格
			if strings.Contains(h.account.Cookie, "e_hk4e_token") == true {
				old, _ := MatchSingle(regexp.MustCompile(`e_hk4e_token=(.+?);`), sc)
				h.account.Cookie = strings.Replace(h.account.Cookie, old, eHk4eToken, -1)
			} else {
				if strings.HasSuffix(h.account.Cookie, ";") == false {
					h.account.Cookie += ";"
				}
				h.account.Cookie += fmt.Sprintf("e_hk4e_token=%s", eHk4eToken)
			}
			break
		}
	}

	return nil
}

func (h *Helper) GetBirthdayRole() error {
	req, err := http.NewRequest("GET", fmt.Sprintf(urlIndex, h.info.Data.GameUid, h.info.Data.Region), nil)
	if err != nil {
		return err
	}
	req.Host = "hk4e-api.mihoyo.com"
	h.setHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respData, _ := ioutil.ReadAll(resp.Body)

	//log.Println(string(respData))

	h.roles = new(RespBirthdayRole)
	if err = json.Unmarshal(respData, h.roles); err != nil {
		return err
	}

	if h.info.Retcode != 0 {
		return errors.New("请求返回错误：" + h.roles.Message)
	}
	return nil
}

func (h *Helper) PostBirthday() error {

	type Payload struct {
		RoleID int `json:"role_id"`
	}

	for i, r := range h.roles.Data.Role {
		data := Payload{
			r.RoleId,
		}
		payloadBytes, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body := bytes.NewReader(payloadBytes)

		req, err := http.NewRequest("POST", fmt.Sprintf(urlPost, h.info.Data.GameUid, h.info.Data.Region), body)
		if err != nil {
			return err
		}
		req.Host = "hk4e-api.mihoyo.com"

		h.setHeader(req)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		respData, _ := ioutil.ReadAll(resp.Body)

		//log.Println(string(respData))

		res := new(RespPost)
		if err = json.Unmarshal(respData, res); err != nil {
			return err
		}

		if res.Retcode == 0 {
			log.Printf("[%d] - 今天是【%s】的生日！已成功获取生日贺卡【%s】！\n", i, r.Name, r.TakePicture)
		} else if res.Retcode == -512009 {
			log.Printf("[%d] - 今天是【%s】的生日！你之前已领取！\n", i, r.Name)
		} else {
			return errors.New("请求返回错误：" + h.roles.Message)
		}

	}

	return nil
}

func (h *Helper) setHeader(req *http.Request) {
	req.Header.Set("Origin", "https://webstatic.mihoyo.com")
	req.Header.Set("Cookie", h.account.Cookie)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) miHoYoBBS/2.36.1")
	req.Header.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Set("Referer", "https://webstatic.mihoyo.com/")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
}
