package response

type User struct {
	Model
	UserName   string `json:"user_name"`
	PassWord   string `json:"pass_word"`
	Email      string `json:"email"`
	NickName   string `json:"nick_name"`
	Avatar     string `json:"avatar"`
	Introduce  string `json:"introduce"`
	CreateTime string `json:"create_time"`
}
