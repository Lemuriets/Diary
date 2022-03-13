package http

type AuthReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
