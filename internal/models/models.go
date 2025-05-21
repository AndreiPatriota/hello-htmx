package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
  

type Tarefa struct {
	gorm.Model
	Titulo string
	Descricao string
	Concluida bool
}

var db *gorm.DB
func InitDb() {
	database, err := gorm.Open(sqlite.Open("./internal/models/db.db"), &gorm.Config{})
	if err != nil {
		panic("Probelma na conexão com o Banco!")
	}

	database.AutoMigrate(&Tarefa{})

	db = database
}