package middleware

import (
	"gin-gorm-app1/common"
	"gin-gorm-app1/dal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// AuthMiddleware 创建一个中间件函数，用于验证 JWT 令牌并解析其声明。
// 如果令牌无效或不存在，将返回未授权的响应。
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头中获取 Authorization 信息
		tokenString := ctx.GetHeader("Authorization")

		// 检查 Authorization 信息是否存在且以 Bearer 开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 提取 token 字符串
		tokenString = tokenString[7:]
		common.Logger.Info("tokenString:", zap.String("tokenKey", tokenString))

		// 解析 token 并检查其有效性
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token无效，权限不足"})
			ctx.Abort()
			return
		}

		// 从声明中获取用户 ID
		//验证通过后获取Claiim中的userId
		userId := claims.UserId

		// 使用全局数据库实例
		DB := common.DB

		// 根据用户 ID 查询用户信息
		var user model.User
		DB.First(&user, userId)

		// 如果用户不存在，返回未授权响应
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户不存在"})
			ctx.Abort()
			return
		}

		// 将用户相关信息设置到上下文中，供后续处理使用
		ctx.Set("user", user.Name)
		ctx.Set("userId", userId)
		ctx.Set("token", tokenString)
		ctx.Set("claims", claims)
		ctx.Set("password", user.Password)
		ctx.Set("telephone", user.Telephone)

		// 继续处理后续请求
		ctx.Next()
	}
}
