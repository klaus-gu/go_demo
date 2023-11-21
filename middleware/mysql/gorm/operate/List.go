package operate

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go_demo/middleware/mysql/model"
)

func GetSingle() {
	first := model.Author{}
	MyDB.Table("author").First(&first)
	fmt.Println(jsoniter.MarshalToString(&first))

	last := model.Author{}
	MyDB.Table("author").Last(&last)
	fmt.Println(jsoniter.MarshalToString(&last))

}

func List() {
	fmt.Println("查询所有")
	authors := []*model.Author{}
	MyDB.Table("author").Find(&authors)
	for _, author := range authors {
		marshalToString, _ := jsoniter.MarshalToString(author)
		fmt.Println(marshalToString)
	}
	fmt.Printf("\n")
	fmt.Println("条件查询")
	MyDB.Table("author").Where("name LIKE ? and id = ?", "%张三%", 23).Find(&authors)
	for _, author := range authors {
		marshalToString, _ := jsoniter.MarshalToString(author)
		fmt.Println(marshalToString)
	}
}

func Page(offset, limit int) {
	fmt.Println("分页查询所有")
	authors := []*model.Author{}
	MyDB.Table("author").Offset(offset).Limit(limit).Find(&authors)
	for _, author := range authors {
		marshalToString, _ := jsoniter.MarshalToString(author)
		fmt.Println(marshalToString)
	}
}
