package database

import (
	"root/mk/internal/model"
)

func Migration() {

	if err := DB.DB.AutoMigrate(&model.Message{}); err != nil {
		panic(err)
	}
}