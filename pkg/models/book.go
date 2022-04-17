package models

import(
	"github.com/jinzhu/gorm"
	"github.com/skozlovtsev/Go-MySQL-CRUD-Bookstore-Management-API/pkg/config"
)

var db * gorm.DB

type Book struct{
	//Структура данных репрезентующая книгу
	//имеет поля
	//ID, CreateAt, UpdateAt, DeleteAt, Name, Author, Publication
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	//При инициации запускает автомиграцию
	//Добавляет в db пропущеные поля из структуры Book
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	//Создание новой записи в базе данных db
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	//Получение из базы данных db всех записей
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	//Получение записи из базы данных db по полю ID
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	//Удаление записи из базы данных db по полю ID
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}