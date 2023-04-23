package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lixvyang/betxin-micro/service/user/model"
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
	return &AdminMiddleware{}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// 检查当前用户是否具有管理员权限
		isAdmin := checkUserIsAdmin(r)

		if !isAdmin {
			// 如果当前用户不是管理员，则返回 403 Forbidden 错误
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}

func checkUserIsAdmin(r *http.Request) bool {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return false
	}

	var user model.User
	if err := json.Unmarshal(body, &user); err != nil {
		return false
	}

	if user.Uid == "" {
		return true
	}

	return false
}
