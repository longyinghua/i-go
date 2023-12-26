package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"io"
	"strconv"
	"strings"
	"time"
)

// 全局变量 es6Client es客户端
var es6Client *elasticsearch6.Client

var es7Client *elasticsearch7.Client

var es8Client *elasticsearch8.Client

// initElasticsearch6
//
//	@Description: 初始化创建elasticsearch6客户端
//	@return client *elasticsearch6.Client
func initElasticsearch6() (client *elasticsearch6.Client) {
	config := elasticsearch6.Config{
		Addresses: []string{"192.168.2.250:9200"},
		Username:  "elastic",
		Password:  "elastic",
	}

	newClient, err := elasticsearch6.NewClient(config)
	if err != nil {
		fmt.Printf("Error creating the client: %s\n", err)
		return
	}

	return newClient
}

// initelasticsearch7
//
//	@Description: 初始化创建elasticsearch7客户端
//	@return client *elasticsearch7.Client
func initelasticsearch7() (client *elasticsearch7.Client) {
	config := elasticsearch7.Config{
		Addresses: []string{"http://10.40.3.29:9200"},
	}

	newClient, err := elasticsearch7.NewClient(config)
	if err != nil {
		fmt.Printf("Error creating the client: %s\n", err)
		return
	}

	return newClient
}

// initElasticsearch8
//
//	@Description: 初始化创建elasticsearch8客户端
//	@return client *elasticsearch8.TypedClient
func initElasticsearch8() (client *elasticsearch8.TypedClient) {
	config := elasticsearch8.Config{
		Addresses: []string{"http://10.40.3.29:9200"},
	}

	newClient, err := elasticsearch8.NewTypedClient(config)
	if err != nil {
		fmt.Printf("Error creating the client: %s\n", err)
		return
	}
	return newClient
}

// indexIsExist
//
//	@Description: 判断索引是否存在
//	@param es7Client *elasticsearch7.Client es7客户端
//	@param indexName ...string 索引名称
func indexIsExist1(es7Client *elasticsearch7.Client, indexName ...string) {
	for _, index := range indexName {
		response, err := es7Client.Indices.Exists([]string{index})
		if err != nil {
			fmt.Printf("Error checking the index: %s\n", err)
		}

		defer response.Body.Close()

		////处理响应，其实可以根据响应的状态码来判断是否存在，当响应状态指示失败时，IsError 返回 true，也就是索引不存在响应结果状态码为404
		//if response.IsError() {
		//	log.Fatalf("Error checking the index: %s\n", response.Status()) //  log.Fatalf()日志打印后会回调os.Exit(1)导致程序退出
		//	//fmt.Printf("Error checking the index: %s\n", response.Status())
		//}

		if response.StatusCode == 404 {
			fmt.Printf("Index %s does not exist\n", index)
		} else {
			fmt.Printf("Index %s exists\n", index)

		}
	}
}

// indexIsExist2
//
//	@Description: 判断索引是否存在
//	@param es7Client *elasticsearch7.Client es7客户端
//	@param indexName string 索引名称
//	@return bool 索引是否存在，true存在，false不存在
func indexIsExist2(es7Client *elasticsearch7.Client, indexName string) bool {
	response, err := es7Client.Indices.Exists([]string{indexName})
	if err != nil {
		fmt.Printf("Error checking the index: %s\n", err)
	}

	defer response.Body.Close()

	////处理响应，其实可以根据响应的状态码来判断是否存在，当响应状态指示失败时，IsError 返回 true，也就是索引不存在响应结果状态码为404
	//if response.IsError() {
	//	log.Fatalf("Error checking the index: %s\n", response.Status()) //  log.Fatalf()日志打印后会回调os.Exit(1)导致程序退出
	//	//fmt.Printf("Error checking the index: %s\n", response.Status())
	//}

	if response.StatusCode == 404 {
		fmt.Printf("Index %s does not exist\n", indexName)
		return false
	} else {
		fmt.Printf("Index %s exists\n", indexName)
		return true
	}
}

// 创建索引成功的响应体模型
type T struct {
	Acknowledged       bool   `json:"acknowledged"`
	ShardsAcknowledged bool   `json:"shards_acknowledged"`
	Index              string `json:"index"`
}

// 创建索引 document
// Review 评价数据
type Review struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userID"`
	Score       int64     `json:"score"`
	Content     string    `json:"content"`
	Tags        []Tag     `json:"tags"`
	Status      int       `json:"status"`
	PublishTime time.Time `json:"PublishTime"`
}

// Tag 评价标签
type Tag struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

// createIndex7
//
//	@Description: 创建索引
//	@param es7Client *elasticsearch7.Client es7客户端
//	@param indexName
func createIndex7(es7Client *elasticsearch7.Client, indexName string) {

	isExist := indexIsExist2(es7Client, indexName)

	if isExist {
		fmt.Printf("%s 索引已经存在,不能重复创建\n", indexName)
	} else if isExist == false {
		response, err := es7Client.Indices.Create(indexName)
		if err != nil {
			fmt.Printf("Error creating the index: %s\n", err)
			return
		}

		defer response.Body.Close()

		s := make([]byte, 1024)
		_, err = response.Body.Read(s)
		if err != nil {
			fmt.Printf("Error reading the response body: %s\n", err)
			return
		}

		//索引创建成功显示响应体
		fmt.Printf("Response body: %s\n", string(s))
		fmt.Println("----------------")

	}

}

func createIndex8(es8client *elasticsearch8.TypedClient) {
	response, err := es8client.Indices.Create("my-review-1").Do(context.Background())

	if err != nil {
		fmt.Printf("Error creating the index: %s\n", err)
		return
	}

	fmt.Printf("success create index :%#v", response.Index)

}

// 创建一条document 并添加到my-review-1的索引中
func indexDocument(es7Client *elasticsearch7.Client, indexName string) {
	//	定义 document结构体对象
	document1 := Review{
		ID:      2,
		UserID:  147,
		Score:   5,
		Content: "这是一个好评",
		Tags: []Tag{
			{Code: 1000, Title: "好评", Name: "aa"},
			{Code: 1001, Title: "物超所值", Name: "bb"},
			{Code: 1002, Title: "有图", Name: "cc"},
		},
		Status:      2,
		PublishTime: time.Now(),
	}

	//	将结构体对象转换成json格式,序列化
	data, _ := json.Marshal(document1)

	//	添加文档
	//  在索引中创建或更新文档，如果文档不存在，则创建文档，如果文档存在，则更新文档
	//第一个参数指定索引名称，第二参数指定需要在索引中创建的文档
	response, err := es7Client.Index(indexName, bytes.NewReader(data))
	if err != nil {
		fmt.Printf("Error indexing document: %s\n", err)
	}

	defer response.Body.Close()

	str := make([]byte, 1024)

	_, err = response.Body.Read(str) //  将响应体读取到字节数组中
	if err != nil {
		fmt.Printf("Error reading the response body: %s\n", err)
	}

	fmt.Printf("创建Document Response body: %s\n", string(str))

}

// 获取 document
// 根据id获取document
func getDocument(es7Client *elasticsearch7.Client, indexName string, id int) {
	response, err := es7Client.Get(indexName, strconv.Itoa(id))
	//response, err := es7Client.Get(indexName, id)
	if err != nil {
		fmt.Printf("Error getting document: %s\n", err)
	}

	defer response.Body.Close()

	str := make([]byte, 1024)
	_, err = response.Body.Read(str)
	if err != nil {
		fmt.Printf("Error reading the response body: %s\n", err)
		return
	}

	if response.StatusCode == 404 {
		fmt.Printf("%s document id=%d Not Found\n", indexName, id)
		return
	} else {
		fmt.Printf("Get Document Response body: %s\n", string(str))
	}

}

// 检索document
// 构建搜索查询可以使用结构化的查询条件
func searchDocument(es7Client *elasticsearch7.Client, indexName string) {
	//	搜索文档
	query := `{
		"query": {
			"match": {
				"content": "好评"
			}
		}
	}`
	response, err := es7Client.Search(
		es7Client.Search.WithIndex(indexName),
		es7Client.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		fmt.Printf("Error searching document: %s\n", err)
		return
	}

	defer response.Body.Close()

	//str := make([]byte, 4096) //  注意：如果响应体过大，需要设置一个较大的字节数组，否则会出现存放的数据不完整
	//
	//_, err = response.Body.Read(str)
	//if err != nil {
	//	fmt.Printf("Error reading the response body: %s\n", err)
	//	return
	//}
	//
	//if response.IsError() {
	//	fmt.Printf("Error searching document: %s\n", response.String())
	//} else {
	//	fmt.Printf("%s\n", string(str))
	//}

	if response.IsError() {
		fmt.Printf("Error searching document: %s\n", response.String())
	} else {
		data, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error reading the response body: %s\n", err)
			return
		}

		fmt.Printf("Response body: %s\n", string(data))
	}

}

// 更新document
func updateDocument(es7Client *elasticsearch7.Client, indexName string) {

	//这个updte会报400请求错误，原因是请求体传的不对，不符合elasticsearch的请求体格式
	d1 := Review{
		ID:      1,
		UserID:  06666666,
		Score:   90,
		Content: "这是一个修改后的好评",
		Tags: []Tag{
			{Code: 1000, Title: "好评"},
			{Code: 1001, Title: "物超所值"},
			{Code: 1002, Title: "有图有真相"},
		},
		Status:      0,
		PublishTime: time.Now(),
	}

	data, err := json.Marshal(d1)
	if err != nil {
		fmt.Printf("Error marshaling document: %s\n", err)
		return
	}

	response, err := es7Client.Update(indexName, "ITD5nowB18ZzUZmVFxZX", strings.NewReader(string(data))) //  此id为elasticsearch的document的_id,而不是数据中id
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	defer response.Body.Close()

	str := make([]byte, 4096)

	_, err = response.Body.Read(str)
	if err != nil {
		fmt.Printf("Error reading the response body: %s\n", err)
		return
	}

	if response.IsError() {
		fmt.Printf("Error updating document: %s\n", response.String())
	} else {
		fmt.Printf("更新Document Response body: %s\n", string(str))
	}

}

// 删除document
func deleteDocument(es7Client *elasticsearch7.Client, indexName string) {
	response, err := es7Client.Delete(indexName, "HzDinowB18ZzU1ZmVbxZf") //  此id为elasticsearch的document的_id,而不是数据中id
	if err != nil {
		fmt.Printf("Error deleting document: %s\n", err)
		return
	}

	defer response.Body.Close()

	if response.IsError() {
		fmt.Printf("%s\n", response.String())
	} else {
		fmt.Printf("%s\n", response.String())
	}
}

// 删除索引
func deleteIndex(es7Client *elasticsearch7.Client, indexName string) {
	response, err := es7Client.Indices.Delete([]string{indexName})
	if err != nil {
		fmt.Printf("Error deleting index: %s\n", err)
		return
	}

	defer response.Body.Close()

	if response.IsError() {
		if response.StatusCode == 404 {
			fmt.Printf("索引不存在,删除失败-----\n")
		}
		fmt.Printf("%s\n", response.String())
	} else {
		s := make([]byte, 4096)
		_, err = response.Body.Read(s)
		if err != nil {
			fmt.Printf("Error reading the response body: %s\n", err)
			return
		}
		fmt.Printf("删除索引Response body: %s\n", string(s))
	}
}

func main() {
	//es6Client = initElasticsearch6()
	es7Client = initelasticsearch7()

	//createIndex7(es7Client, "my-review-2") //  索引不存在时创建索引，存在时不创建
	//indexIsExist2(es7Client, "my-review-1") //  判断单个索引是否存在
	//indexDocument(es7Client, "my-review-1")
	//getDocument(es7Client, "my-review-1", 1)
	//searchDocument(es7Client, "my-review-1")
	updateDocument(es7Client, "my-review-1")
	//deleteDocument(es7Client, "my-review-1")
	//deleteIndex(es7Client, "my-review-2")
}
