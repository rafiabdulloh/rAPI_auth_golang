package models

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnect() {

	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatal("err env")
		return
	}

	// create database manual terlebih dahulu
	dsn := "root:rafi123@tcp(localhost:3308)/go_learning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Users{})

	DB = db

}

func Seeder() {

	password := "mochrafi123"

	// generate hash dari password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	data := Users{
		Username: "moch_rafi02",
		Password: string(hashedPassword),
	}

	// verifikasi password
	// err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	// if err != nil {
	// 	fmt.Println("Password tidak cocok")
	// } else {
	// 	fmt.Println("Password cocok")
	// }

	var count int64
	DB.Model(&Users{}).Count(&count)
	if count == 0 {
		// Jika database kosong, maka insert data awal
		fmt.Println("========seeder running=============")
		DB.Create(&data)
	}

}
