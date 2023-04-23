// Code generated by goctl. DO NOT EDIT.
package types

type Errmsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type User struct {
	IdentityNumber string `json:"identity_number"`
	Uid            string `json:"uid"`
	FullName       string `json:"full_name"`
	AvatarUrl      string `json:"avatar_url"`
	Biography      string `json:"biography"`
}

type CreateUserReq struct {
	IdentityNumber string `json:"identity_number"`
	Uid            string `json:"uid"`
	FullName       string `json:"full_name"`
	AvatarUrl      string `json:"avatar_url"`
	Biography      string `json:"biography"`
}

type CreateUserResp struct {
	Errmsg
	Data string `json:"data"`
}