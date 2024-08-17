package entry

type User struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	Avatar           string `json:"avatar"`
	Bot              bool   `json:"bot"`
	MemberOpenId     string `json:"member_openid"`      //群成员id
	UserOpenId       string `json:"user_openid"`        //用户id
	UnionOpenID      string `json:"union_openid"`       // 特殊关联应用的 openid
	UnionUserAccount string `json:"union_user_account"` // 机器人关联的用户信息，与union_openid关联的应用是同一个
}

// Member 群成员
type Member struct {
	GuildID  string    `json:"guild_id"`
	JoinedAt Timestamp `json:"joined_at"`
	Nick     string    `json:"nick"`
	User     *User     `json:"user"`
	Roles    []string  `json:"roles"`
	OpUserID string    `json:"op_user_id,omitempty"`
}

// DeleteHistoryMsgDay 消息撤回天数
type DeleteHistoryMsgDay = int

// 支持的消息撤回天数，除这些天数之外，传递其他值将不会撤回任何消息
const (
	NoDelete                              = 0  // 不删除任何消息
	DeleteThreeDays   DeleteHistoryMsgDay = 3  // 3天
	DeleteSevenDays   DeleteHistoryMsgDay = 7  // 7天
	DeleteFifteenDays DeleteHistoryMsgDay = 15 // 15天
	DeleteThirtyDays  DeleteHistoryMsgDay = 30 // 30天
	DeleteAll         DeleteHistoryMsgDay = -1 // 删除所有消息
)
