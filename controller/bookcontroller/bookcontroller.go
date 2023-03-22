package bookcontroller

import (
	"encoding/json"
	"errors"
	"github.com/kenalinguaridho/books-api/config"
	"github.com/kenalinguaridho/books-api/entities"
	"github.com/kenalinguaridho/books-api/helper"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var books []entities.Book
	var bookResponse []entities.BookResponse
	
	if err := config.DB.Joins("Author").Find(&books).Find(&bookResponse).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List books", bookResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var book entities.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// validasi author

	var author entities.Author

	if err  := config.DB.First(&author, book.AuthorId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Create book success", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)
	
	var book entities.Book
	var bookResponse entities.BookResponse

	if err := config.DB.Joins("Author").First(&book, id).First(&bookResponse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Book not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail book", bookResponse)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var book entities.Book

	idParam := mux.Vars(r)["id"]

	id,_:= strconv.Atoi(idParam)

	if err := config.DB.First(&book, id).Error;err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Book not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var bookPayLoad entities.Book

	if err := json.NewDecoder(r.Body).Decode(&bookPayLoad); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var author entities.Author
	if bookPayLoad.AuthorId != 0 {
		if err := config.DB.First(&author, bookPayLoad.AuthorId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Author not found", nil)
				return
			}
	
			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err := config.DB.Where("id = ?", id).Updates(&bookPayLoad).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Update book success", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var book entities.Book
	
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	result := config.DB.Delete(&book, id)

	if result.Error != nil {
		helper.Response(w, 500, result.Error.Error(), nil)
		return
	}

	if result.RowsAffected == 0 {
		helper.Response(w, 404, "Book not found", nil)
		return
	}

	helper.Response(w, 200, "Book delete success", nil)

}