package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db * gorm.DB
)

func Connect(){
	//Инициализируем новое подключение к базе данных
	d, err := gorm.Open("postgres", "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Moscow")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	//Возвращает объект типа gorm.DB
	return db
}