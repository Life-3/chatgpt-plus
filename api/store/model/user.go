package model

type User struct {
	BaseModel
	Mobile      string
	Password    string
	Avatar      string
	Salt        string // 密码盐
	TotalTokens int64  // 总消耗 tokens
	Calls       int    // 剩余对话次数
	ImgCalls    int    // 剩余绘图次数
	ChatConfig  string `gorm:"column:chat_config_json"` // 聊天配置 json
	ChatRoles   string `gorm:"column:chat_roles_json"`  // 聊天角色
	ExpiredTime int64  // 账户到期时间
	Status      bool   `gorm:"default:true"` // 当前状态
	LastLoginAt int64  // 最后登录时间
	LastLoginIp string // 最后登录 IP
}
