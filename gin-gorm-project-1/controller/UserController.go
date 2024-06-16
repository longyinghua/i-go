package controller

import (
	"errors"
	"gin-gorm-app1/common"
	"gin-gorm-app1/dal/model"
	"gin-gorm-app1/response"
	"gin-gorm-app1/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// Register 处理用户注册请求。
// 它从请求体中解析用户信息，验证用户输入，创建新用户，并生成认证令牌。
// @Summary 用户注册
// @Description 用户注册
// @Tags User注册
// @Accept json
// @Produce json
// @Param user body model.User  true "注册用户信息"
// @Success 200 {object} response.result{} "注册成功"
// @Failure 400 {object} response.result{} "请求参数错误"
// @Failure 422 {object} response.result{} "用户已存在"
// @Router /auth/register [post]
func Register(ctx *gin.Context) {
	// 解析请求体中的用户信息。
	var requestUser model.User
	err := ctx.ShouldBind(&requestUser)
	if err != nil {
		// 日志记录绑定过程中的错误。
		common.Logger.Error("user param bind json error", zap.String("error", err.Error()))
		// 对非验证错误的绑定问题返回400错误。
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"code": 400,
		})
		return
	}

	// 验证用户输入是否符合规则。
	// 获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 校验手机号长度。
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	// 校验密码长度。
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 如果用户名为空，生成一个随机用户名。
	if len(name) == 0 {
		name = utils.CreateRandomString(10)
	}
	// 检查电话号码是否已存在。
	if isTelephoneExist(common.DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		common.Logger.Warn("用户已存在", zap.String("telephone", telephone), zap.String("msg", "用户手机号已存在"))
		return
	}

	// 使用bcrypt对密码进行加密。
	hasePassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		common.Logger.Error("用户密码加密错误", zap.String("password", password), zap.String("msg", err.Error()))
		return
	}
	// 创建新用户。
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassowrd),
	}
	common.DB.Create(&newUser)

	// 生成并返回认证令牌。
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		common.Logger.Error("token generate error", zap.String("msg", err.Error()))
		return
	}
	response.Success(ctx, gin.H{"token": token}, "注册成功")
	common.Logger.Info("用户注册成功", zap.String("telephone", telephone), zap.String("msg", "注册成功"))
}

// Login处理用户登录请求
// 它验证用户提供的凭据，并在验证成功后生成并返回一个JWT令牌。
// @Summary 用户登录
// @Description 用户登录
// @Tags User登陆
// @Accept json
// @Produce json
// @Param user body model.User  true "登录用户信息"
// @Success 200 {object} response.result{} "登录成功"
// @Failure 400 {object} response.result{} "请求错误"
// @Failure 500 {object} response.result{} "服务器错误"
// @Router /auth/login [post]
func Login(ctx *gin.Context) {
	// 获取数据库连接
	db := common.DB

	// 解析请求中的用户数据
	var requestUser model.User
	err := ctx.ShouldBind(&requestUser)
	// 日志记录绑定过程中的任何错误
	if err != nil {
		common.Logger.Error("requestuser param bind json error", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数错误" + err.Error(),
			"code": 400,
		})
		return
	}

	// 验证用户输入的数据是否符合预期格式
	//  参数校验
	if err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "code": 400})
			common.Logger.Error("参数校验失败", zap.String("msg", err.Error()))
			return
		}
	}

	// 提取用户输入的姓名、电话和密码
	//从用户请求的json中获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	// 校验电话号码长度
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		common.Logger.Warn("登陆手机号必须为11位", zap.String("telephone", telephone), zap.String("msg", "登陆手机号必须为11位"))
		return
	}
	// 校验密码长度
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		common.Logger.Warn("登陆密码不能少于6位", zap.String("telephone", telephone), zap.String("msg", "登陆密码不能少于6位"))
		return
	}

	// 根据电话号码查询用户
	var foundUser model.User
	//执行查询判断手机号是否存在
	result := db.Where("telephone = ?", telephone).First(&foundUser)

	// 处理查询错误
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			common.Logger.Warn("用户手机号不存在", zap.String("user", name), zap.String("msg", "用户手机号不存在，请先注册再登陆！"))
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
			return
		} else {
			common.Logger.Error("查询用户时发生错误", zap.String("error", result.Error.Error()))
			response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
			return
		}
	}

	// 校验用户是否存在
	//判断密码是否正确
	if foundUser.ID == 0 {
		common.Logger.Warn("用户手机号不存在", zap.String("user", name), zap.String("msg", "用户手机号不存在，请先注册再登陆！"))
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil {
		common.Logger.Warn("用户登陆密码错误", zap.String("user", foundUser.Name), zap.String("password", password), zap.String("msg", "登陆密码错误"))
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 生成并返回JWT令牌
	//发送token
	token, err := common.ReleaseToken(foundUser)
	common.Logger.Info("用户token", zap.String("token", token), zap.String("user", foundUser.Name), zap.String("password", requestUser.Password), zap.String("telephone", foundUser.Telephone))
	if err != nil {
		common.Logger.Error("系统异常", zap.String("system 异常", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		return
	}

	// 返回成功响应和令牌
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登陆成功")
	common.Logger.Info("登陆成功", zap.String("login success", foundUser.Telephone))
}

// 方法：判断电话号码是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	// 查询数据库中是否存在该电话号码的用户
	var user model.User
	db.Where("telephone=?", telephone).First(&user)

	// 如果用户存在，则返回true
	if user.ID != 0 {
		return true
	}
	// 否则返回false
	return false
}

// Info 回复用户相关信息。
//
// 此函数从 gin.Context 中获取用户相关的信息，并以 JSON 格式返回这些信息。
// 它不接受任何参数，但使用 gin.Context 来获取请求上下文中的用户数据。
// 返回的状态码为 http.StatusOK，表示操作成功。
func Info(ctx *gin.Context) {
	// 从上下文中获取用户相关的信息。
	user, _ := ctx.Get("user")
	userIds, _ := ctx.Get("userId")
	token, _ := ctx.Get("token")
	claims, _ := ctx.Get("claims")
	password, _ := ctx.Get("password")
	telephone, _ := ctx.Get("telephone")

	// 构建并返回包含用户信息的 JSON 响应。
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"user":      user,
			"userId":    userIds,
			"telephone": telephone,
			"password":  password,
			"token":     token,
			"claims":    claims,
		},
		"msg": "获取用户信息成功",
	})

	//fmt.Println(user, userIds, token, claims, password, telephone)
}

// 该函数接收一个 gin.Context 类型的参数 ctx，用于处理 HTTP 请求上下文。
// 函数主要通过解析请求中的令牌，验证其有效性并获取对应用户的信息。
// 用户信息查询
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags token 信息
// @Accept json
// @Produce json
// @Param tokenkey body model.Token  true "用户token信息"
// @Success 200 {object} response.result{} "获取用户信息成功"
// @Failure 400 {object} response.result{} "获取用户信息失败"
// @Failure 500 {object} response.result{} "服务器错误"
// @Router /auth/information [post]
func GetUserInfo(ctx *gin.Context) {
	// 定义一个 model.Token 类型的变量，用于存储解析出的令牌信息。
	var tokenstring model.Token

	// 尝试从请求体中解析 JSON 格式的令牌信息。
	err := ctx.BindJSON(&tokenstring)
	if err != nil {
		// 如果解析失败，记录错误日志并返回错误信息。
		common.Logger.Error("token bind param error", zap.String("token error", err.Error()))
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"code": http.StatusBadRequest,
				"data": nil,
				"msg":  "token参数错误" + err.Error(),
			},
		)
		return
	}

	// 如果解析成功，但存在验证错误，处理验证错误。
	// 参数校验
	if err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg":  err.Error(),
				"code": 400,
				"data": nil,
			})
			common.Logger.Error("参数校验失败", zap.String("msg", err.Error()))
			return
		}
	}

	// 解析令牌，获取令牌的有效性、声明等信息。
	token, claims, err := common.ParseToken(tokenstring.Token)
	if err != nil || !token.Valid {
		// 如果令牌解析失败或无效，返回未授权错误。
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token无效，请检查token",
			"data": nil,
		})
		ctx.Abort()
		return
	}

	// 检查令牌是否过期。
	if claims.ExpiresAt < time.Now().Unix() {
		// 如果令牌过期，返回未授权错误。
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token已过期，请使用新token",
			"data": nil,
		})
		ctx.Abort()
		return
	}

	// 从令牌声明中获取用户ID。
	userId := claims.UserId

	// 获取数据库实例。
	db := common.DB

	// 根据用户ID查询用户信息。
	var userinfo model.User
	db.Where("id = ?", userId).First(&userinfo)

	// 如果用户ID为0，表示未找到用户，返回未授权错误。
	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "查无此用户",
			"data": nil,
		})
	}

	// 如果用户信息查询成功，返回用户信息。
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户信息查询成功",
		"data": gin.H{
			"userId":    userId,
			"name":      userinfo.Name,
			"telephone": userinfo.Telephone,
			"password":  userinfo.Password,
			"token":     tokenstring.Token,
			"claims":    claims,
		},
	})

	common.Logger.Info("用户信息查询成功", zap.String("userId", strconv.Itoa(int(userId))), zap.String("name", userinfo.Name), zap.String("telephone", userinfo.Telephone), zap.String("password", userinfo.Password), zap.String("token", tokenstring.Token), zap.Any("claims", claims))

	return
}

// 用户删除
// @Summary 删除用户信息
// @Description 删除用户信息
// @Tags user信息删除
// @Accept json
// @Produce json
// @Param user body model.User  true "用户token信息"
// @Success 200 {object} response.result{} "获取用户信息成功"
// @Failure 400 {object} response.result{} "获取用户信息失败"
// @Failure 500 {object} response.result{} "服务器错误"
// @Router /user/delete [post]
func DeleteUser(ctx *gin.Context) {
	db := common.DB

	// 解析请求中的用户数据
	var requestUser model.User
	err := ctx.ShouldBind(&requestUser)
	// 日志记录绑定过程中的任何错误
	if err != nil {
		common.Logger.Error("requestuser param bind json error", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数错误" + err.Error(),
			"code": 400,
			"data": nil,
		})
		return
	}

	// 验证用户输入的数据是否符合预期格式
	//  参数校验
	if err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg":  err.Error(),
				"code": 400,
				"data": nil,
			})
			common.Logger.Error("参数校验失败", zap.String("msg", err.Error()))
			return
		}
	}

	// 检查电话号码是否已存在
	isExist := isTelephoneExist(common.DB, requestUser.Telephone)

	if isExist {
		db.Where("telephone = ?", requestUser.Telephone).Delete(&requestUser)
		common.Logger.Info("删除用户成功", zap.String("telephone", requestUser.Telephone))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除用户成功" + requestUser.Telephone,
			"data": nil,
		})
	} else {
		common.Logger.Info("删除用户失败", zap.String("telephone", requestUser.Telephone))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "删除用户失败" + requestUser.Telephone,
			"data": nil,
		})
	}

	return

}
