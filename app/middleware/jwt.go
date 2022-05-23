/*
 * @Author: alexander.huang
 * @Date:   2022-05-23 21:40:51
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-23 21:40:51
 */
package middleware

import (
	"github.com/Anzz-bot/DouYin_demo/app/common/response"
	"github.com/Anzz-bot/DouYin_demo/app/services"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.URL.Query().Get("token")
		if tokenStr == "" {
			response.TokenFail(c)
			c.Abort()
			return
		}
		global.App.Log.Info(tokenStr)
		tokenStr = tokenStr[:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})

		if err != nil { //  || services.JwtService.IsInBlacklist(tokenStr)   封装redis可实现黑名单校验退出
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}
		// token 续签
		//if claims.ExpiresAt-time.Now().Unix() < global.App.Config.Jwt.RefreshGracePeriod {
		//	lock := global.Lock("refresh_token_lock", global.App.Config.Jwt.JwtBlacklistGracePeriod)
		//	if lock.Get() {
		//		err, user := services.JwtService.GetUserInfo(GuardName, claims.Id)
		//		if err != nil {
		//			global.App.Log.Error(err.Error())
		//			lock.Release()
		//		} else {
		//			tokenData, _, _ := services.JwtService.CreateToken(GuardName, user)
		//			c.Header("new-token", tokenData.AccessToken)
		//			c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
		//			_ = services.JwtService.JoinBlackList(token)
		//		}
		//	}
		//}

		c.Set("token", token)
		c.Set("id", claims.Id)

	}
}
