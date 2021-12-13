package param

type UserSignup struct {
	Type     int    `json:"type"`
	Account  string `json:"account"`
	PassWord string `json:"passWord"`
}
