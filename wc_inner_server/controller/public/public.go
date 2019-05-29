package public

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wc_inner_server/common"
	"wc_inner_server/common/code"
	"wc_inner_server/model/behavior/public"
	"wc_inner_server/model/table"
)

var sendError = common.SendError
var admin model.Admins
var md5compute = common.Md5Compute

func LoginToSystem(c *gin.Context) {

	mobile := c.PostForm("username")
	password := c.PostForm("password")
    if err := public.GetPostUserInfo("mobile", mobile, &admin); err != nil {
    	sendError("未录入该员工信息", c)
		return
	}

    if admin.Password == md5compute(password + md5compute(admin.No)) {
    	timeNow := time.Now().Format("2006-01-02")
		if err := public.UpdateUserInfo("last_login",  timeNow, admin.No, &admin);err != nil {
			sendError("登陆发生未知错误", c)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": code.SuccessCode.SUCCESS,
		"msg": "success",
		"data": admin,
	})

}

func GetAuthInfo(c *gin.Context)  {
	tokenString,err := createToken(admin.No)
	if err != nil {
		sendError("登陆发生未知错误", c)
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": code.SuccessCode.SUCCESS,
		"msg": "success",
		"data": tokenString,
	})
}

func createToken(userNo string) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()
	claims["no"] = userNo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(code.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString , nil
}


