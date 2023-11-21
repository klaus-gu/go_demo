package operate

import (
	"fmt"
	"go_demo/middleware/mysql/model"
)

func Update(author *model.Author) {
	update := MyDB.Table("author").Where("id", author.ID).Update("name", author.Name)
	fmt.Println("更新行数：", update.RowsAffected)
}

// 更新的行必须是非空的，且只有真正的变化之后才会被更新
func UpdateMultiColumn(authors *model.Author) {
	updates := MyDB.Table("author").Model(&authors).Updates(authors)
	fmt.Println("更新行数：", updates.RowsAffected)
}

// 需要使用 where 条件进行限制
func UpdateWithSelectedField(author *model.Author) {
	update := MyDB.Table("author").Select("name").Where("id = ?", author.ID).Update("name", author.Name)
	fmt.Println("更新行数：", update.RowsAffected)
}

// 批量更新的 Author 内部不能出现 ID，否则更新失败
func UpdateBatch(author *model.Author) {
	updates := MyDB.Table("author").Where("id  IN ?", []int{10, 11}).Updates(author)
	fmt.Println("更新行数：", updates.RowsAffected)
}
