package controller

// BookCreate1 新增存储book书籍信息接口-1
//
//	@Summary		存储book书籍信息
//	@Tags			单本书籍存储
//	@Description	前端传入book书籍相关信息直接保存book对象信息的接口
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result1	"插入数据成功提示"
//	@Failure		400		{object}	ret.Result1	"错误提示"
//	@Router			/create/books1 [post]
//func BookCreate1() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//		var book model.Book
//
//		//err := c.ShouldBindJSON(&book)
//		err := c.ShouldBind(&book)
//		if err != nil {
//			//log.Printf("book param bind json err:%v", err)
//			common.Logger.Error("book param bind json error", zap.String("error", err.Error())) //使用zap包的Error()方法
//			//下面为几种调用zap打印日志方式
//			//logger.Error("book param bind json error", zap.String("error", err.Error()))  //使用自定义封装后zap包的Logger.Error()方法
//			//logger.SugaredLogger.Errorf("book param bind json error:%v", err) //使用带格式化的zap包的Errorf()方法
//			//logger.Logger.Warn("book param bind json error", zap.String("error", err.Error()))  //使用zap包的Warn()方法
//			//logger.Logger.Debug("book param bind json error", zap.String("error", err.Error()))  //使用zap包的Debug()方法
//			//zap.L().Error("book param bind json error", zap.String("error", err.Error()))   //日志包中使用了全局替换zap包中全局的Logger实例，后续在其他包中只需使用zap.L()调用即可
//		}
//
//		//  参数校验
//		if err != nil {
//			//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
//			_, ok := err.(validator.ValidationErrors)
//			if !ok {
//				c.JSON(http.StatusBadRequest, gin.H{
//					//不是validator.validationErrors类型errors的错误直接返回错误信息
//					"msg":  err.Error(),
//					"code": 400,
//				})
//				common.Logger.Error("参数校验失败", zap.String("msg", err.Error()))
//				//zap.L().Error("参数校验失败", zap.String("msg", err.Error()))
//				return
//			}
//		}
//
//		err = controller.CreateBook1(&book) //  插入数据成功返回true，失败返回false
//
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"code": 400,
//				"msg":  fmt.Sprintf("保存数据失败:%v", err),
//			})
//			common.Logger.Warn("保存数据失败", zap.String("msg", fmt.Sprintf("%+v", err.Error()))) //错误信息也在日志中打印出来
//			//zap.L().Warn("保存数据失败", zap.String("msg", fmt.Sprintf("%+v", err.Error())))       //作用同上
//		} else {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 200,
//				"msg":  fmt.Sprintf("保存数据成功"),
//			})
//			common.Logger.Info("保存数据成功", zap.String("msg", "Success")) //保存数据成功信息也在日志中打印出来
//			//zap.L().Info("保存数据成功", zap.String("msg", "Success"))       //作用同上
//		}
//
//	}
//}

// BookCreate2 新增存储book书籍信息接口-2
//
//	@Summary		存储book书籍信息
//	@Tags			单本书籍存储
//	@Description	前端传入book书籍相关信息直接保存book对象信息的接口
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result1	"插入数据成功提示"
//	@Failure		400		{object}	ret.Result1	"错误提示"
//	@Router			/create/books2 [post]
//func BookCreate2(c *gin.Context) {
//
//	var book model.Book
//
//	err := c.ShouldBindJSON(&book)
//	if err != nil {
//		//log.Printf("book param bind json err:%v", err)
//		common.Logger.Error("book param bind json error", zap.String("error", err.Error()))
//	}
//
//	//  参数校验
//	if err != nil {
//		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
//		_, ok := err.(validator.ValidationErrors)
//		if !ok {
//			c.JSON(http.StatusOK, gin.H{
//				//不是validator.validationErrors类型errors的错误直接返回错误信息
//				"msg":  err.Error(),
//				"code": 400,
//			})
//			common.Logger.Error("参数校验失败", zap.String("msg", err.Error()))
//			//zap.L().Error("参数校验失败", zap.String("msg", err.Error()))
//			return
//		}
//	}
//
//	err = mysqlcrud.CreateBook1(&book) //  插入数据成功返回true，失败返回false
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"code": 400,
//			"msg":  fmt.Sprintf("保存数据失败:%v", err),
//		})
//		common.Logger.Warn("保存数据失败", zap.String("msg", fmt.Sprintf("%+v", err.Error())))
//		//zap.L().Warn("保存数据失败", zap.String("msg", fmt.Sprintf("%+v", err.Error())))
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  fmt.Sprintf("保存数据成功"),
//		})
//		common.Logger.Info("保存数据成功", zap.String("msg", "Success"))
//		//zap.L().Info("保存数据成功", zap.String("msg", "Success"))
//	}
//
//}

// BookCreate3 新增存储book书籍信息接口-3
//
//	@Summary		存储book书籍信息
//	@Tags			多本书籍存储
//	@Description	前端传入book书籍相关信息直接保存book对象信息的接口
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.RequestPayloadBook	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result1					"插入数据成功提示"
//	@Failure		400		{object}	ret.Result1					"错误提示"
//	@Router			/create/books3 [post]
//func BookCreate3(c *gin.Context) {
//
//	// 解析JSON请求体
//	var requestPayload struct {
//		Books []*model.Book `json:"books"`
//	}
//
//	err := c.ShouldBindJSON(&requestPayload)
//
//	//  参数校验
//	if err != nil {
//		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
//		_, ok := err.(validator.ValidationErrors)
//		if !ok {
//			c.JSON(http.StatusBadRequest, gin.H{
//				//不是validator.validationErrors类型errors的错误直接返回错误信息
//				"error": err.Error(),
//				"code":  400,
//			})
//			common.Logger.Error("book param bind json error", zap.String("error", err.Error()))
//			//zap.L().Error("book param bind json error", zap.String("error", err.Error()))
//			return
//		}
//	}
//
//	err = mysqlcrud.CreateBook2(requestPayload.Books) //  插入数据成功返回true，失败返回false
//
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 400,
//			"msg":  fmt.Sprintf("保存数据失败:%v", err),
//		})
//		common.Logger.Warn("保存数据失败", zap.String("msg", fmt.Sprintf("%+v", err.Error())))
//		//zap.L().Warn("保存数据失败", zap.String("msg", fmt.Sprintf("%+v", err.Error())))
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  fmt.Sprintf("保存数据成功"),
//		})
//		common.Logger.Info("保存数据成功", zap.String("msg", "Success"))
//		//zap.L().Info("保存数据成功", zap.String("msg", "Success"))
//	}
//
//}

//apifox接口传参示例
//{
//"books": [
//{
//"title": "<<C语言之路>>",
//"author": "wawa",
//"price": 100
//},
//{
//"title": "<<C++语言之路>>",
//"author": "wlalaawa",
//"price": 200
//},
//{
//"title": "<<Kubernetes进阶指南>>",
//"author": "kakak",
//"price": 300
//}
//],
//"hello": "world"
//}

// BookCreate4 新增存储book书籍信息接口-4
//
//	@Summary		存储book书籍信息
//	@Tags			多本书籍存储
//	@Description	前端传入book书籍相关信息直接保存book对象信息的接口
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.RequestPayloadBook	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result1					"插入数据成功提示"
//	@Failure		400		{object}	ret.Result1					"错误提示"
//	@Router			/create/books4 [post]
//func BookCreate4(c *gin.Context) {
//
//	//引用全局自定义的模型
//	var requestPlayload model.RequestPayloadBook
//
//	//解析JSON请求体，将json请求体绑定到requestPlayload中，api传入的books的key对应requestPlayload.Books的key名，books的value为列表对应requestPlayload.Books的value，【】*Book
//	err := c.ShouldBindJSON(&requestPlayload)
//	//  注意：参数绑定时，一定要包含对应的key，key对应模型中对应的字段名，如果前端传入的key名和模型中字段名不一致，则不会绑定到对应模型的字段中
//
//	//  参数校验
//	if err != nil {
//		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
//		_, ok := err.(validator.ValidationErrors)
//		if !ok {
//			c.JSON(http.StatusBadRequest, gin.H{
//				//不是validator.validationErrors类型errors的错误直接返回错误信息
//				"error": err.Error(),
//				"code":  400,
//			})
//			common.Logger.Error("book param bind json error", zap.String("error", err.Error()))
//			//zap.L().Error("book param bind json error", zap.String("error", err.Error()))
//			return
//		}
//	}
//
//	err = mysqlcrud.CreateBook2(requestPlayload.Books) //  插入数据成功返回true，失败返回false
//
//	//BookBody := requestPlayload.Books
//	//err = mysqlcrud.CreateBook2(BookBody) //  插入数据成功返回true，失败返回false
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"code": 400,
//			"msg":  fmt.Sprintf("保存数据失败:%v", err),
//		})
//		common.Logger.Warn("保存数据失败", zap.String("msg", fmt.Sprintf("%+v", err.Error())))
//		//zap.L().Warn("保存数据失败", zap.String("msg", fmt.Sprintf("%+v", err.Error())))
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  fmt.Sprintf("保存数据成功"),
//		})
//		common.Logger.Info("保存数据成功", zap.String("msg", "Success"))
//		//zap.L().Info("保存数据成功", zap.String("msg", "Success"))
//	}
//
//}
