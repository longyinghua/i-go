// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "龙应华",
            "url": "http://www.swagger.io/support",
            "email": "542791872@qq.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/create/books1": {
            "post": {
                "description": "前端传入book书籍相关信息直接保存book对象信息的接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "单本书籍存储"
                ],
                "summary": "存储book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "插入数据成功提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    }
                }
            }
        },
        "/create/books2": {
            "post": {
                "description": "前端传入book书籍相关信息直接保存book对象信息的接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "单本书籍存储"
                ],
                "summary": "存储book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "插入数据成功提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    }
                }
            }
        },
        "/create/books3": {
            "post": {
                "description": "前端传入book书籍相关信息直接保存book对象信息的接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "多本书籍存储"
                ],
                "summary": "存储book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestPayloadBook"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "插入数据成功提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    }
                }
            }
        },
        "/create/books4": {
            "post": {
                "description": "前端传入book书籍相关信息直接保存book对象信息的接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "多本书籍存储"
                ],
                "summary": "存储book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestPayloadBook"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "插入数据成功提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    }
                }
            }
        },
        "/custom/books1": {
            "post": {
                "description": "前端传入book书籍相关信息查询book对象信息的接口，通过author查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍查询"
                ],
                "summary": "查询book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询数据成功提示",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ret.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Book"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    }
                }
            }
        },
        "/custom/books2": {
            "post": {
                "description": "前端传入book书籍相关信息查询book对象信息的接口,通过id查询，返回模型",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍查询"
                ],
                "summary": "查询book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询数据成功提示",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ret.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Book"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result"
                        }
                    }
                }
            }
        },
        "/custom/books3": {
            "post": {
                "description": "前端传入book书籍相关信息查询book对象信息的接口,通过id查询返回map",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍查询"
                ],
                "summary": "查询book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询数据成功提示",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ret.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Book"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result"
                        }
                    }
                }
            }
        },
        "/custom/books4": {
            "post": {
                "description": "前端传入book书籍相关信息查询book对象信息的接口,通过字段查询返回数据列表集合",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍查询"
                ],
                "summary": "查询book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TFilter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询数据成功提示",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ret.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Book"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result"
                        }
                    }
                }
            }
        },
        "/custom/books5": {
            "post": {
                "description": "前端传入book书籍相关信息查询book对象信息的接口,通过书籍字段查询返回数据列表集合",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍查询"
                ],
                "summary": "查询book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询数据成功提示",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ret.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Book"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result"
                        }
                    }
                }
            }
        },
        "/custom/books6": {
            "post": {
                "description": "前端传入book书籍相关信息查询book对象信息的接口,通过字段查询返回数据列表集合",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍查询"
                ],
                "summary": "查询book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询数据成功提示",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ret.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Book"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result"
                        }
                    }
                }
            }
        },
        "/delete/books1": {
            "post": {
                "description": "前端传入book书籍相关信息删除book对象信息的接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍删除"
                ],
                "summary": "删除book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除数据成功提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result"
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result"
                        }
                    }
                }
            }
        },
        "/update/books1": {
            "post": {
                "description": "前端传入book书籍相关信息直接更新book对象信息的接口，更新单列字段",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍更新"
                ],
                "summary": "更新book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新数据成功提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    }
                }
            }
        },
        "/update/books2": {
            "post": {
                "description": "前端传入book书籍相关信息直接更新book对象信息的接口，更新多列字段，map写法",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍更新"
                ],
                "summary": "更新book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新数据成功提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    }
                }
            }
        },
        "/update/books3": {
            "post": {
                "description": "前端传入book书籍相关信息直接更新book对象信息的接口，更新多列字段，模型写法",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍更新"
                ],
                "summary": "更新book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新数据成功提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    }
                }
            }
        },
        "/update/books4": {
            "post": {
                "description": "前端传入book书籍相关信息直接更新book对象信息的接口，更新多列，更新选定字段",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍更新"
                ],
                "summary": "更新book书籍信息",
                "parameters": [
                    {
                        "description": "book书籍的title，author，price相关信息",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新数据成功提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    },
                    "400": {
                        "description": "错误提示",
                        "schema": {
                            "$ref": "#/definitions/ret.Result1"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "description": "作者",
                    "type": "string"
                },
                "create_at": {
                    "description": "创建时间",
                    "type": "string",
                    "format": "date-time"
                },
                "id": {
                    "description": "主键",
                    "type": "integer"
                },
                "price": {
                    "description": "价格",
                    "type": "integer"
                },
                "publish_date": {
                    "description": "出版日期",
                    "type": "string",
                    "format": "date-time"
                },
                "title": {
                    "description": "书籍名称",
                    "type": "string"
                },
                "update_at": {
                    "description": "更新时间",
                    "type": "string",
                    "format": "date-time"
                }
            }
        },
        "model.RequestPayloadBook": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Book"
                    }
                }
            }
        },
        "model.TFilter": {
            "type": "object",
            "properties": {
                "column": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "ret.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "ret.Result1": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "10.40.3.9:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Book API",
	Description:      "Book API+SQL 增删改查",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}