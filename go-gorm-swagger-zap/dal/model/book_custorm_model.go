package model

// 自定义结构体，可以基于json复制粘贴生产struct
type RequestPayloadBook struct {
	Books []*Book `json:"books"`
}

type TFilter struct {
	Column string `json:"column"`
	Value  string `json:"value"`
}
