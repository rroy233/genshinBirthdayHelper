package helper

type RespInfo struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Game       string `json:"game"`
		Region     string `json:"region"`
		GameBiz    string `json:"game_biz"`
		Nickname   string `json:"nickname"`
		Level      int    `json:"level"`
		RegionName string `json:"region_name"`
		GameUid    string `json:"game_uid"`
	} `json:"data"`
}

type RespBirthdayRole struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		NickName string `json:"nick_name"`
		Uid      int    `json:"uid"`
		Region   string `json:"region"`
		Gender   int    `json:"gender"`
		Role     []struct {
			JumpTarget    string `json:"jump_target"`
			IsPartake     bool   `json:"is_partake"`
			GalResource   string `json:"gal_resource"`
			RoleGender    int    `json:"role_gender"`
			JumpType      string `json:"jump_type"`
			GalXml        string `json:"gal_xml"`
			Bgm           string `json:"bgm"`
			JumpEndTime   string `json:"jump_end_time"`
			JumpStartTime string `json:"jump_start_time"`
			RoleId        int    `json:"role_id"`
			Name          string `json:"name"`
			TakePicture   string `json:"take_picture"`
		} `json:"role"`
		DrawNotice   bool   `json:"draw_notice"`
		CurrentTime  string `json:"CurrentTime"`
		IsShowRemind bool   `json:"is_show_remind"`
	} `json:"data"`
}

type RespPost struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
	} `json:"data"`
}
