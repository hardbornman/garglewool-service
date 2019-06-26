package gwUser

import (
	"fmt"
	myjwt "gitee.com/gbat/utils/middleware"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"github.com/hardbornman/garglewool-service/service/wechat"
	"net/http"
	"time"
)

type LoginResult struct {
	Token string `json:"token"`
	model.GwUser
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func GetDataByTime(c *gin.Context) {
	isPass := c.GetBool("isPass")
	if !isPass {
		return
	}
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

// 登录
func Auth(c *gin.Context) {
	account := c.Request.PostFormValue("loginname")
	pwd := c.Request.PostFormValue("loginpwd")

	if account != "" && pwd != "" {
		user, err := service.Login(account, pwd)
		if err == nil && user.Userid > 0 {
			generateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败" + err.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "json 解析失败",
		})
		return
	}

}

// 生成令牌
func generateToken(c *gin.Context, user model.GwUser) {
	j := &myjwt.JWT{
		[]byte("newtrekWang"),
	}

	claims := myjwt.CustomClaims{
		user.Userid,
		0,
		user.UserName,
		user.UserPhone,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),    // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 48*3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                      //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}
	fmt.Println(token)
	data := LoginResult{
		GwUser: user,
		Token:  token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}

// 微信登录
func WeChatLogin(c *gin.Context) {
	secret := c.Request.PostFormValue("secret")
	appid := c.Request.PostFormValue("appid")
	js_code := c.Request.PostFormValue("js_code")
	if secret != "" && appid != "" && js_code != "" {
		m, err := wechat.WechatLogin(js_code, appid, secret)
		if err == nil && len(m) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    m,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败" + err.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "json 解析失败",
		})
		return
	}

}
