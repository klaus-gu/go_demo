package operate

import (
	"fmt"
	"go_demo/middleware/mysql/model"
)

/**
根据是否有主键判断是更新还是新增操作。
*/

func SaveUpdate(author *model.Author) {
	save := MyDB.Table("author").Save(author)
	fmt.Println("影响行数：", save.RowsAffected)
}

func SaveUpdateBatch(authors []*model.Author) {
	save := MyDB.Table("author").Save(authors)
	fmt.Println("影响行数：", save.RowsAffected)
}
