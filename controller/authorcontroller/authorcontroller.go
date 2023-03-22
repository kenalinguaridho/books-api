package authorcontroller

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
	var authors []entities.Author

	if err := config.DB.Find(&authors).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List authors", authors)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var author entities.Author

	// menyimpan data json ke variabel author
	if err := json.NewDecoder(r.Body).Decode(&author) ; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	// tutup request body
	defer r.Body.Close()

	// data yang telah ditampung dalam body
	// dikirim ke database dengan cara berikut
	if err := config.DB.Select("Name", "Gender", "Email", "Age").Create(&author).Error; err != nil {
		helper.Response(w, 500, "Create author failed", nil)
		return
	}

	helper.Response(w, 201, "Create author success", nil)

}

func Detail(w http.ResponseWriter, r *http.Request) {

	idParam := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(idParam)

	var author entities.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail author", author)

}

func Update(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(idParam)

	var author entities.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("id = ?", id).Updates(&author).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Update author success", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var author entities.Author

	result := config.DB.Delete(&author, id)

	if result.Error != nil {
		helper.Response(w, 500, result.Error.Error(), nil)
		return
	}

	if result.RowsAffected == 0 {
		helper.Response(w, 404, "Author not found", nil)
		return
	}

	helper.Response(w, 200, "Delete author success", nil)
}