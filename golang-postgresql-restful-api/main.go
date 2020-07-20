package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"csdn/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("不能加载属性.env文件")
	}

	app := api.App{}
	app.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	app.RunServer()
}

