package controller

import (
	"helloshuya-web-api/entities"
	"helloshuya-web-api/infrastructure"
	"net/http"

	"github.com/labstack/echo"
)

func GetPosts(c echo.Context) error {
	posts, _ := GetRepoPosts()
	return c.JSON(http.StatusOK, posts)
}

func GetRepoPosts() (entities.Posts, error) {
	db := infrastructure.GetDBInstance()
	posts := entities.Posts{}

	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
