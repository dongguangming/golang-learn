package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres

	"csdn/middlewares"
	"csdn/models"
	"csdn/responses"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// 初始化数据库连接，添加路由表
func (a *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	a.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\n 不能连接到数据库： %s", DbName)
		log.Fatal("发生了连接数据库错误:", err)
	} else {
		fmt.Printf("恭喜连接到数据库： %s", DbName)
	}

	a.DB.Debug().AutoMigrate(&models.Article{}) //database migration

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json

	a.Router.HandleFunc("/", home).Methods("GET")
	
	s := a.Router.PathPrefix("/api").Subrouter() // subrouter to add auth middleware

	s.HandleFunc("/article", a.CreateArticle).Methods("POST")
	s.HandleFunc("/article", a.GetArticles).Methods("GET")
	s.HandleFunc("/article/{id:[0-9]+}", a.GetArticle).Methods("GET")
	s.HandleFunc("/article/{id:[0-9]+}", a.UpdateArticle).Methods("PUT")
	s.HandleFunc("/article/{id:[0-9]+}", a.DeleteArticle).Methods("DELETE")
}

func (a *App) RunServer() {
	log.Printf("\nServer starting on port 9999")
	log.Fatal(http.ListenAndServe("192.168.8.200:9999", a.Router))
}

func home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Golang学习系列第五天： Golang和PostgreSQL开发 RESTful API")
}
