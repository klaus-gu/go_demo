package operate

import (
	"fmt"
	"go_demo/middleware/mysql/model"
)

func Update(author *model.Author) {
	update := MyDB.Table("author").Where("id", author.ID).Update("name", author.Name)
	fmt.Println("更新行数", update.RowsAffected)
}

func UpdateBatch(authors []*model.Author) {

}
