package message

const (
	LoginMesType           = "LoginMesType"
	LoginResMesType        = "LoginResMesType"
	RegisterMesType        = "RegisterMesType"
	RegisterResMesType     = "RegisterResMesType"
	NotifyUserMesType      = "NotifyUserMesType"
	NotifyUserMesUp        = "NotifyUserMesUp"
	NotifyUserMesDown      = "NotifyUserMesDown"
	SmsMesType             = "SmsMesType"
	SmsMesSimpleMesType    = "SmsMesSimpleMesType"
	SmsMesSimpleResMesType = "SmsMesSimpleResMesType"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	User User `json:"user"`
}
type LoginResMes struct {
	ResMes ResMes `json:"resmes"`
	User   User   `json:"user"`
	Users  []int  `json:"users"`
}
type ResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
type RegisterMes struct {
	User User `json:"user"`
}
type RegisterResMes struct {
	ResMes `json:"resmes"`
}
type NotifyUserMes struct {
	UserId  int    `json:"userId"`
	MesType string `json:"mesType"`
}
type SmsMes struct {
	Content string `json:"content"`
	User    User   `json:"user"`
}
type SmsMesSimpleMes struct {
	SmsMes   SmsMes `json:"smsMes"`
	ToUserId int    `json:"toUserId"`
}
type SmsMesSimpleResMes struct {
	ResMes ResMes
}
