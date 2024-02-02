package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//code=1000...用户模块的错误

	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	//code=2000...文章模块的错误

	ERROR_ART_NOT_EXIST = 2001

	//code=3000...分类模块的错误

	ERROR_CATETITLE_USED      = 3001
	ERROR_CATETITLE_NOT_EXIST = 3002

	//code=4000...角色模块的错误

	ERROR_ROLE_USED      = 4001
	ERROR_ROLE_NOT_EXIST = 4002
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "用户密码错误",
	ERROR_USER_NOT_EXIST:   "无法查找到用户",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN错误",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",

	ERROR_ART_NOT_EXIST: "文章不存在",

	ERROR_CATETITLE_USED:      "分类已存在",
	ERROR_CATETITLE_NOT_EXIST: "分类不存在",

	ERROR_ROLE_USED:      "角色已存在",
	ERROR_ROLE_NOT_EXIST: "角色不存在",
}

func GetCode(code int) string {
	return codeMsg[code]
}
