package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
  )
  

type Tarefa struct {
	gorm.Model
	Titulo string
	Descricao string
	Concluida bool
}

var DB *gorm.DB
func InitDb() {
	db, err := gorm.Open(sqlite.Open("./internal/models/db.db"), &gorm.Config{})
	if err != nil {
		panic("Probelma na conexão com o Banco!")
	}

	db.AutoMigrate(&Tarefa{})

	DB = db
}