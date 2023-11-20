package operate

import (
	"fmt"
	"go_demo/middleware/mysql/model"
)

func Insert(author *model.Author) {
	result := MyDB.Table("author").Create(author)
	fmt.Println("插入条数：", result.RowsAffected)
	fmt.Println("插入返回主键ID：", author.ID)
}

func InsertBatch(authors []*model.Author) {
	result := MyDB.Table("author").Create(authors)
	fmt.Println("插入条数：", result.RowsAffected)
	for _, id := range authors {
		fmt.Println("插入返回主键ID：", id.ID)
	}
}
