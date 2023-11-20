package main

import (
	"go_demo/middleware/mysql/gorm/operate"
	"go_demo/middleware/mysql/model"
)

/*
*
插入数据表的时候通过 struct 内部定义的字段与标签内的 column 与数据库表字段映射关联
*/
func main() {
	// 插入
	//author := model.Author{Name: "zhangsan"}
	//operate.Insert(&author)
	//authors := []*model.Author{{Name: "lisi"}, {Name: "wangwu"}, {Name: "zhaoliu"}}
	//operate.InsertBatch(authors)

	// 更新
	author := model.Author{Name: "hahaha", ID: 1}
	operate.Update(&author)
}
