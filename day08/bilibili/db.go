package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func connect(DBType string, dsn string) (err error) {
	if DBType == "mysql"
	if dsn == "" {
		dsn = "postgres://postgres:123456@127.0.0.1:5432/bilibili"
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}
	fmt.Println(" DB connect Successful: ", DB.Name())
	return nil
}

func InitTables() {
	if DB == nil {
		err := connect("")
		if err != nil {
			return
		}
	}
	err := DB.AutoMigrate(&Bilibili{}, &VideoData{})
	// err := DB.AutoMigrate(&Bilibili{}, &StatInfo{}, &VideoData{}, &Subtitle{}, &VideoRights{}, &VideoPage{}, &UserInfo,
	// 	&DimensionInfo{}, &UserGarb{})
	if err != nil {
		panic(err)
	}
}

func Save(b *Bilibili) {
	res := DB.Create(b)
	fmt.Println(res.Error)
	DB.Commit()
}
