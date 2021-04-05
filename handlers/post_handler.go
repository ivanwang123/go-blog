package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/ivanwang123/go-blog/models"
	"github.com/ivanwang123/go-blog/stores"
	"github.com/labstack/echo/v4"
)

func ListPosts(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	posts, err := store.PostStore.Posts()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Unable to retrieve posts: %v", err))
	}

	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully retrieved posts",
		"posts":   posts,
	})
}

func PaginatePosts(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	page, pageErr := strconv.Atoi(c.QueryParam("page"))
	if pageErr != nil || page <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Page is not an integer")
	}
	posts, limit, postsErr := store.PostStore.PaginatedPosts(page)
	if postsErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to retrieve posts: %v", postsErr))
	}

	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully retrieved posts",
		"posts":   posts,
		"hasMore": len(posts) == limit,
	})
}

func GetPost(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not an integer")
	}

	post, err := store.PostStore.Post(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Post does not exist")
	}

	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully retrieved post",
		"post":    post,
	})
}

func CreatePost(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	title := c.FormValue("title")
	content := c.FormValue("content")
	userId, err := strconv.Atoi(c.FormValue("user-id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID is not an integer")
	}

	validate := validator.New()
	post := &models.Post{
		Title:   title,
		Content: content,
		UserId:  userId,
	}
	validateErr := validate.Struct(post)
	if validateErr != nil {
		return c.JSON(http.StatusBadRequest, &Res{
			"message": "Unable to create post",
			"errors":  strings.Split(validateErr.(validator.ValidationErrors).Error(), "\n"),
		})
	}

	if err := store.PostStore.CreatePost(post); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to create post: %v", err))
	}
	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully created post",
		"postId":  post.Id,
	})
}

func UpdatePost(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not an integer")
	}
	title := c.FormValue("title")
	content := c.FormValue("content")

	validate := validator.New()
	post := &models.Post{
		Id:      id,
		Title:   title,
		Content: content,
	}
	validateErr := validate.Struct(post)
	if validateErr != nil {
		return c.JSON(http.StatusBadRequest, &Res{
			"message": "Unable to update post",
			"errors":  strings.Split(validateErr.(validator.ValidationErrors).Error(), "\n"),
		})
	}

	if err := store.PostStore.UpdatePost(post); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to update post: %v", err))
	}
	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully updated post",
	})
}

func DeletePost(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not an integer")
	}

	if err := store.PostStore.DeletePost(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to delete post: %v", err))
	}
	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully deleted post",
	})
}
