package utils

import (
	"github.com/casbin/casbin"
)

// code
const (
	Success            = 200
	LoginSuccess       = 2000
	UserNameError      = 1000
	PasswordError      = 1001
	InValidTokenParams = 1002
	UserNoExist        = 1003
	UserExists         = 1004
	CrypPasswordErr    = 1005
	DBInsertSuccess    = 1006
	TenantExist        = 1007
	TenantNoExist      = 1008
	DBInsertFailed     = 1009
)

// other information
const (
	APPVERSION     = "1.0"
	DBTYPE         = "mysql"
	WebViewPerm    = "user"
	BaseTimeFormat = "2006-01-02 15:04:05"
	TimeZone       = "Asia/Shanghai"
)

var (
	// ProjectPath project path
	ProjectPath string
	// CasbinEnforcer casbinEndorcer object
	CasbinEnforcer *casbin.Enforcer
)

// MsgFlags parse message
var MsgFlags = map[int]string{
	Success:            "ok",
	LoginSuccess:       "登录成功",
	UserNameError:      "用户名错误",
	PasswordError:      "密码错误",
	UserNoExist:        "用户不存在",
	UserExists:         "用户已存在",
	InValidTokenParams: "请求未携带token，无权限访问",
	CrypPasswordErr:    "加密失败",
	DBInsertSuccess:    "数据库插入成功",
	DBInsertFailed:     "数据库插入失败",
	TenantExist:        "租户已存在",
	TenantNoExist:      "租户不存在",
}

// GetStatusText get message base on code
func GetStatusText(code int) string {
	return MsgFlags[code]
}
