package errmsg

var codeMsg = make(map[int]string)

// var codeMsg = map[int]string{
// 	SUCCSE:                 "OK",
// 	ERROR:                  "FAIL",
// 	ERROR_BIND:             "输入参数错误",
// 	ERROR_USERNAME_USED:    "用户名已存在！",
// 	ERROR_PASSWORD_WRONG:   "密码错误",
// 	ERROR_USER_NOT_EXIST:   "用户不存在",
// 	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
// 	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
// 	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
// 	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
// 	ERROR_USER_NO_RIGHT:    "该用户无权限",
// 	ERROR_UPDATE_USER:      "更新用户失败",
// 	ERROR_LIST_USER:        "查询用户列表失败",
// 	ERROR_DELETE_USER:      "删除用户失败",
// 	ERROR_GET_USER:         "获取用户失败",

// 	ERROR_ART_NOT_EXIST: "文章不存在",

// 	ERROR_CATENAME_USED:   "该分类已存在",
// 	ERROR_CATE_NOT_EXIST:  "该分类不存在",
// 	ERROR_UPDATE_CATENAME: "更新分类错误",
// 	ERROR_DELETE_CATENAME: "删除分类错误",
// 	ERROR_LIST_CATEGORY:   "查询分类列表错误",

// 	ERROR_BONUSE_EXIST:     "奖金已存在",
// 	ERROR_BONUSE_NOT_EXIST: "奖金不存在",
// 	ERROR_LIST_BONUSE:      "查询奖金列表失败",
// 	ERROR_UPDATE_BONUSE:    "更新奖金失败",
// 	ERROR_GET_BONUSE:       "查询奖金失败",
// 	ERROR_DELETE_BONUSE:    "删除奖金失败",
// 	ERROR_CREATE_BONUSE:    "创建奖金失败",

// 	ERROR_UPDATE_TOPIC: "更新话题错误",
// 	ERROR_DELETE_TOPIC: "删除话题错误",
// 	ERROR_LIST_TOPIC:   "查询话题列表错误",
// }
