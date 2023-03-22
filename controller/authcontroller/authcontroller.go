package authcontroller

import (
	"encoding/json"
	"errors"
	"github.com/kenalinguaridho/books-api/config"
	"github.com/kenalinguaridho/books-api/entities"
	"github.com/kenalinguaridho/books-api/helper"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {

	// mengambil input json
	var userInput entities.User

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		helper.Response(w, 400, "Gagal decode json", nil)
		return
	}

	defer r.Body.Close()

	// cek username
	var user entities.User
	if err := config.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Username atau password salah", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	// cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		helper.Response(w, 404, "Username atau password salah", nil)
		return
	}

	// proses generate JWT(JSON Web Token)
	expTime := time.Now().Add(time.Minute * 1)
	claim := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer : "go-jwt-mux",
			ExpiresAt : jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasikan algoritma untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	// set token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name : "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	helper.Response(w, 200, "Login berhasil", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	
	// mengambil input json
	var userInput entities.User

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		helper.Response(w, 500, "Gagal decode json", nil)
		return
	}

	defer r.Body.Close()

	// hash password menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert ke database
	if err := config.DB.Create(&userInput).Error; err != nil {
		helper.Response(w, 500, "Register user gagal", nil)
		return
	}

	helper.Response(w, 201, "Register user berhasil", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	// hapus token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name : "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
	})

	helper.Response(w, 200, "Logout berhasil", nil)
}