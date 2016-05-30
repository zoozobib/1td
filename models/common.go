package models

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	//"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	//"time"
)

/*数据库默认注册*/
func init() {
	// set default database
	//orm.RegisterDataBase("default", "mysql", "yaoqianshu:yaoqianshu@/yaoqianshu?charset=utf8", 30)
	//orm.Debug = true
}

/*end by zoozobib*/

/*md5加密方法*/
func Md5(str string) string {
	if str != "" {
		s := md5.New()
		s.Write([]byte(str))
		ss := s.Sum(nil)
		return hex.EncodeToString(ss)
	} else {
		return ""
	}
}

/*end by zoozobib*/

/**/
func Change_strings(str string) string {
	var left_p, right_p string

	if len(str) != 11 {
		return ""
	}

	if str != "" {
		for i := 0; i < 3; i++ {
			left_p += string(str[i])
		}

		for i := 7; i < 11; i++ {
			right_p += string(str[i])
		}

		return left_p + fmt.Sprintf("%s", "%%") + right_p

	} else {
		return ""
	}
}

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

/*base64加密*/

func Base64_encode(src []byte) []byte {
	var coder = base64.NewEncoding(base64Table)
	return []byte(coder.EncodeToString(src))
}

/*end*/

/*base64解密*/
func Base64_decode(src []byte) ([]byte, error) {
	var coder = base64.NewEncoding(base64Table)
	return coder.DecodeString(string(src))
}

/*end*/
