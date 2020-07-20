package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"csdn/models"
	"csdn/responses"
)

// 创建文章
func (a *App) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "文章创建成功"}

	article := &models.Article{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &article)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	article.Prepare()

	if err = article.Validate(); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if existed, _ := article.GetArticle(a.DB); existed != nil {
		resp["status"] = "failed"
		resp["message"] = "文章已存在"
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}
	
	articleCreated, err := article.SaveArticle(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["article"] = articleCreated
	responses.JSON(w, http.StatusCreated, resp)
	return
}

//获取所有文章
func (a *App) GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := models.GetArticles(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, articles)
	return
}

//获取单篇文章
func (a *App) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	article, err := models.GetArticleById(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, article)
	return
}

//更改文章
func (a *App) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "修改文章成功！！！"}

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	//article, err := models.GetArticleById(id, a.DB)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	articleUpdate := models.Article{}
	if err = json.Unmarshal(body, &articleUpdate); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	articleUpdate.Prepare()

	_, err = articleUpdate.UpdateArticle(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, resp)
	return
}

//删除文章
func (a *App) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "文章删除成功！！！"}

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	//article, err := models.GetArticleById(id, a.DB)

	err := models.DeleteArticle(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, resp)
	return
}
