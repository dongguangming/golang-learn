package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type Article struct {
	gorm.Model
	Title        string `gorm:"size:100;not null;unique" json:"title"`
	Content      string `gorm:"not null"                 json:"content"`
	Category     string `gorm:"size:50;not null"         json:"category"`
	Tag          string `gorm:"size:50;not null"         json:"tag"`
  Author       string `gorm:"size:50;not null"         json:"author"`
}

func (article *Article) Prepare() {
	article.Title = strings.TrimSpace(article.Title)
	article.Content = strings.TrimSpace(article.Content)
	article.Category = strings.TrimSpace(article.Category)
	article.Tag = strings.TrimSpace(article.Tag)
	article.Author = strings.TrimSpace(article.Author)
}

func (article *Article) Validate() error {
	if article.Title == "" {
		return errors.New("Name is required")
	}
	if article.Content == "" {
		return errors.New("Description about venue is required")
	}
	
	if article.Category == "" {
		return errors.New("Category of venue is required")
	}
	if article.Tag == "" {
		return errors.New("Category of venue is required")
	}
	if article.Author == "" {
		return errors.New("Category of venue is required")
	}
	return nil
}

func (article *Article) SaveArticle(db *gorm.DB) (*Article, error) {
	var err error

	// Debug a single operation, show detailed log for this operation
	err = db.Debug().Create(&article).Error
	if err != nil {
		return &Article{}, err
	}
	return article, nil
}

func (article *Article) GetArticle(db *gorm.DB) (*Article, error) {
	newArticle := &Article{}
	if err := db.Debug().Table("articles").Where("title = ?", article.Title).First(newArticle).Error; err != nil {
		return nil, err
	}
	return newArticle, nil
}

func GetArticles(db *gorm.DB) (*[]Article, error) {
	articles := []Article{}
	if err := db.Debug().Table("articles").Find(&articles).Error; err != nil {
		return &[]Article{}, err
	}
	return &articles, nil
}

func GetArticleById(id int, db *gorm.DB) (*Article, error) {
	article := &Article{}
	if err := db.Debug().Table("articles").Where("id = ?", id).First(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (article *Article) UpdateArticle(id int, db *gorm.DB) (*Article, error) {
	if err := db.Debug().Table("articles").Where("id = ?", id).Updates(Article{
		Title : article.Title,
	  Content : article.Content,
	  Category : article.Category,
	  Tag : article.Tag,
	  Author : article.Author}).Error; err != nil {
		return &Article{}, err
	}
	return article, nil
}

func DeleteArticle(id int, db *gorm.DB) error {
	if err := db.Debug().Table("articles").Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}

