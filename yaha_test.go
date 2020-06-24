package yoha

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpenConnect(t *testing.T) {
	db, err := Open("mysql", "root:19931003@/Yearning?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_ = db.DB().Ping()
}
