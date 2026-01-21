package database

import (
	"fmt"

	"github.com/philippgille/chromem-go"
)

var collections = map[string]*chromem.Collection{}

var db *chromem.DB

func NewChromemDB() (*chromem.DB, error) {
	if db != nil {
		return nil, fmt.Errorf("请勿重复创建 Chromem DB")
	}
	db = chromem.NewDB()
	c, err := db.CreateCollection("knowledge-base", nil, nil)
	if err != nil {
		panic(err)
	}
	collections["knowledge-base"] = c
	return db, nil
}

func GetCollection(name string) (*chromem.Collection, error) {
	if db == nil {
		return nil, fmt.Errorf("Chromem DB 未初始化")
	}
	if c, ok := collections[name]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("未找到 Chromem 集合 %s", name)
}
