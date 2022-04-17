package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/skozlovtsev/Go-MySQL-CRUD-Bookstore-Management-API/pkg/utils"
	"github.com/skozlovtsev/Go-MySQL-CRUD-Bookstore-Management-API/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()  //Получаем из базы данных db в переменную newBooks типа []Book все записи 
	res, _ := json.Marshal(newBooks)  //В переменную res получаем завчение перменной newBooks в формате json
	w.Header().Set("Content-Type", "pkglication/json")  //устанавливаем json в качестве типа контента в заголовке ответа
	w.WriteHeader(http.StatusOK)  //Добавляем в заголовок статус 200(OK)
	w.Write(res)  //запись res в качестве ответа на запрос
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]  //Из ппеременной vars типа map получаем значение поля с ключем bookId
	ID, err := strconv.ParseInt(bookId, 0, 0)  //Конвертируем bookId в int64
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)  //Получаем из базы данных db запись в виде объекта типа Book по ID в bookDetails
	res, _ := json.Marshal(bookDetails)  //В переменную res получаем завчение перменной bookDetails в формате json
	w.Header().Set("Content-Type", "pkglication/json")  //устанавливаем json в качестве типа контента в заголовке ответа
	w.WriteHeader(http.StatusOK)  //Добавляем в заголовок статус 200(OK)
	w.Write(res)  //запись res в качестве ответа на запрос
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()  //Создаем из объекта CreateBook новую запись в базе данных, с помощью метода CreateBook()
	res, _ := json.Marshal(b)  //В переменную res получаем завчение перменной b в формате json
	w.WriteHeader(http.StatusOK)  //Добавляем в заголовок статус 200(OK)  
	w.Write(res)  //запись res в качестве ответа на запрос
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]  //Из ппеременной vars типа map получаем значение поля с ключем bookId
	ID, err :=	strconv.ParseInt(bookId, 0, 0)  //Конвертируем bookId в int64
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)  //Удаляем из ёбазы данных db запись по ID, получаем удаленную запись
	res, _ := json.Marshal(book)  //В переменную res получаем завчение перменной book в формате json
	w.Header().Set("Content-Type", "pkglication/json")  //устанавливаем json в качестве типа контента в заголовке ответа
	w.WriteHeader(http.StatusOK)  //Добавляем в заголовок статус 200(OK)
	w.Write(res)  //запись res в качестве ответа на запрос
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
    vars := mux.Vars(r)
	bookId := vars["bookId"]  //Из ппеременной vars типа map получаем значение поля с ключем bookId
	ID, err := strconv.ParseInt(bookId, 0, 0)  //Конвертируем bookId в int64
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)   //Получаем из базы данных db запись в виде объекта типа Book по ID в bookDetails 
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication !=  ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)  //Сохраняем обновленное значение в базе данных db
	res, _ := json.Marshal(bookDetails)  //В переменную res получаем завчение перменной book в формате json
	w.Header().Set("Content-Type", "pkglication/json")  //устанавливаем json в качестве типа контента в заголовке ответа
	w.WriteHeader(http.StatusOK)  //Добавляем в заголовок статус 200(OK)
	w.Write(res)  //запись res в качестве ответа на запрос
}