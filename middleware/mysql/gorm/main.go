package main

import "go_demo/middleware/mysql/gorm/operate"

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
	//author := model.Author{Name: "nihao", ID: 1}
	//operate.Update(&author)
	//operate.UpdateMultiColumn(&author)
	//operate.UpdateWithSelectedField(&author)
	//authorBatch := model.Author{Name: "nihao"}
	//operate.UpdateBatch(authorBatch)

	// 删除
	//author := model.Author{ID: 11}
	//operate.Delete(&author)
	//author := model.Author{Name: "nihao"}
	//	//operate.DeleteBatch(&author)

	// 更新或者新增
	//author := &model.Author{
	//	Name: "尼古拉斯",
	//}
	////operate.SaveUpdate(author)
	//authors := []*model.Author{{Name: "张三"}, {ID: 9, Name: "李1四"}}
	//operate.SaveUpdateBatch(authors)

	// 查询
	//operate.GetSingle()
	operate.List()
}
