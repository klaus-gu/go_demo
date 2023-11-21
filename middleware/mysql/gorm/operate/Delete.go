package operate

import (
	"fmt"
	"go_demo/middleware/mysql/model"
)

// 实际是根据结构体内部的 主键 id 进行删除的
func Delete(author *model.Author) {
	del := MyDB.Table("author").Delete(&author)
	fmt.Println("影响行数：", del.RowsAffected)
}

// 删除是一样的道理，当使用 where 条件语句的时候， 传入的结构体内不能出现主键，否则会退化到上个 Delete 的例子中
func DeleteBatch(author *model.Author) {
	del := MyDB.Table("author").Where("name LIKE ?", "%"+author.Name+"%").Delete(&author)
	fmt.Println("影响行数：", del.RowsAffected)
	if del.Error != nil {
		fmt.Println("删除异常1：", del.Error)
	}
	del2 := MyDB.Table("author").Delete(&author, "id IN ?", []int{3, 4})
	fmt.Println("影响行数：", del2.RowsAffected)
	if del2.Error != nil {
		fmt.Println("删除异常2：", del2.Error)
	}
}
