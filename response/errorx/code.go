package errorx

import "errors"

var ErrCode = struct {
	S_OK                           int //成功
	E_FAILED                       int //400 失败
	E_WEB_FAILURE                  int //500 故障
	E_UNKNOWN                      int //2147745797 未知
	E_INVALID_ACCESS_TOKEN         int //401 无效的访问令牌(token)
	E_INVALID_ACCESS_SIGN          int //2164260866 无效的访问令牌(sign)
	E_INVALID_ACCESS_SIGNEXPIRE    int //2164260867 无效的访问令牌(sign expire)
	E_BAD_USERNAME_OR_PASSWORD     int //2164260868 用户名或密码错误
	E_INVALID_COMMAND              int //2164260869 e_无效的_命令
	E_NOT_IMPLEMENTED              int //2164260870 未实施
	E_ACCESS_DENIED                int //2164260871 访问被拒绝
	E_INVALID_PARAM_VALUE          int //402 无效的参数值
	E_INVALID_APP_ID               int //403 无效的APPID
	E_INVALID_TENANT_ID            int //406 无效的TENANTID
	E_INVALID_TOKEN                int //407 无效的TOKEN
	E_OBJECT_NOT_FOUND             int //2164260873 找不到对象
	W_DATA_EXISTS                  int //1879048193 数据已存在
	E_NO_VALUE_FOR_UPDATE          int //2164260874 没有用于更新的值
	E_ACTION_NOT_ALLOWED_BY_STATUS int //2164260875 状态不允许的操作
	E_REQUEST_ALREAY_EXISTS        int //2164260876 请求一直存在
	E_USER_ACCOUNT_LOCKED          int //2164260877 用户帐户已锁定
	E_SESSION_TIME_OUT             int //2164260879 会话超时,请退出重新登陆
	E_SESSION_USER_ERROR           int //2164260881 用户不匹配
	E_USER_STATE_ERROR             int //2164260882 用户状态错误
	E_CERT_STATE_ERROR             int //2164260883 客户资质状态错误
	//E_AUTH_NOTENOUGH                     int     //2147745798 权限不足
	E_INSUFFICIENT_QUOTA   int //2147680257 额度不足
	E_REQUEST_TOO_FREQUENT int //2181042177 请求太频繁
	E_ACCOUNT_CENTER       int
}{
	S_OK:                           0,          //成功
	E_FAILED:                       400,        //2164260864 失败
	E_WEB_FAILURE:                  500,        //2013265920 故障
	E_UNKNOWN:                      0x80040005, //2147745797 未知
	E_INVALID_ACCESS_TOKEN:         401,        //2164260865 无效的访问令牌(token)
	E_INVALID_ACCESS_SIGN:          0x81000002, //2164260866 无效的访问令牌(sign)
	E_INVALID_ACCESS_SIGNEXPIRE:    0x81000003, //2164260867 无效的访问令牌(sign expire)
	E_BAD_USERNAME_OR_PASSWORD:     0x81000004, //2164260868 用户名或密码错误
	E_INVALID_COMMAND:              0x81000005, //2164260869 e_无效的_命令
	E_NOT_IMPLEMENTED:              0x81000006, //2164260870 未实施
	E_ACCESS_DENIED:                0x81000007, //2164260871 访问被拒绝
	E_INVALID_PARAM_VALUE:          402,        //2164260872 无效的参数值
	E_INVALID_APP_ID:               403,        //2164260865 无效的APPID
	E_OBJECT_NOT_FOUND:             0x81000009, //2164260873 找不到对象
	W_DATA_EXISTS:                  0x70000001, //1879048193 数据已存在
	E_NO_VALUE_FOR_UPDATE:          0x8100000A, //2164260874 没有用于更新的值
	E_ACTION_NOT_ALLOWED_BY_STATUS: 0x8100000B, //2164260875 状态不允许的操作
	E_REQUEST_ALREAY_EXISTS:        0x8100000C, //2164260876 请求一直存在
	E_USER_ACCOUNT_LOCKED:          0x8100000D, //2164260877 用户帐户已锁定
	E_SESSION_TIME_OUT:             0x8100000F, //2164260879 会话超时,请退出重新登陆
	E_SESSION_USER_ERROR:           0x81000011, //2164260881 用户不匹配
	E_USER_STATE_ERROR:             0x81000012, //2164260882 用户状态错误
	E_CERT_STATE_ERROR:             0x81000013, //2164260883 客户资质状态错误
	//E_AUTH_NOTENOUGH                     : 0x80040006,     //2147745798 权限不足
	E_REQUEST_TOO_FREQUENT: 0x82001001, //2181042177 请求太频繁
}

var (
	NoData = errors.New("无数据")
)
