// package main

// import (
// 	"encoding/json"
// 	"log"
// )

// type Information struct {
// 	Name string `json:"name"`
// 	Addr string `json:"addr"`
// }

// type DataItem struct {
// 	ArticleId   string `json:"articleId"`
// 	ArticleName string `json:"articleName"`
// }

// func TestStructure() {
// 	var inf Information
// 	inf.Name = "Alice"
// 	inf.Addr = "Green Street"
// 	data, err := json.Marshal(inf)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(string(data))

// 	var data1 DataItem
// 	data1.ArticleId = "1003"
// 	data1.ArticleName = "以中国新发展为世界提供新机遇"
// 	data1_str, _ := json.Marshal(data1)
// 	log.Println(string(data1_str))
// }

// func main() {

// 	TestStructure()
// }

package test
