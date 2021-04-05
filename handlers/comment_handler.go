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

func GetCommentsByPost(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	postId, err := strconv.Atoi(c.Param("post-id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Post ID is not an integer")
	}

	comments, err := store.CommentStore.CommentsByPost(postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Unable to retrieve comments: %v", err))
	}

	return c.JSON(http.StatusOK, &Res{
		"message":  "Successfully retrieved comments",
		"comments": comments,
	})
}

func GetComment(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not an integer")
	}

	comment, err := store.CommentStore.Comment(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Comment does not exist")
	}

	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully retrieved comment",
		"comment": comment,
	})
}

func CreateComment(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	content := c.FormValue("content")
	post_id, err := strconv.Atoi(c.FormValue("post-id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Post ID is not an integer")
	}

	validate := validator.New()
	comment := &models.Comment{
		Content: content,
		PostId:  post_id,
	}
	validateErr := validate.Struct(comment)
	if validateErr != nil {
		return c.JSON(http.StatusBadRequest, &Res{
			"message": "Unable to create comment",
			"errors":  strings.Split(validateErr.(validator.ValidationErrors).Error(), "\n"),
		})
	}

	if err := store.CommentStore.CreateComment(comment); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to create comment: %v", err))
	}
	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully created comment",
	})
}

func UpdateComment(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not an integer")
	}
	content := c.FormValue("content")

	validate := validator.New()
	comment := &models.Comment{
		Id:      id,
		Content: content,
	}
	validateErr := validate.Struct(comment)
	if validateErr != nil {
		return c.JSON(http.StatusBadRequest, &Res{
			"message": "Unable to update comment",
			"errors":  strings.Split(validateErr.(validator.ValidationErrors).Error(), "\n"),
		})
	}

	if err := store.CommentStore.UpdateComment(comment); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to update comment: %v", err))
	}
	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully updated comment",
	})
}

func DeleteComment(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not an integer")
	}

	if err := store.CommentStore.DeleteComment(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to delete comment: %v", err))
	}
	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully deleted comment",
	})
}
