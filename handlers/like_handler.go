package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivanwang123/go-blog/models"
	"github.com/ivanwang123/go-blog/stores"
	"github.com/labstack/echo/v4"
)

func GetLikesByPost(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	postId, err := strconv.Atoi(c.Param("post-id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Post ID is not an integer")
	}

	likes, err := store.LikeStore.LikesByPost(postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to get likes: %v", err))
	}

	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully retrieved likes",
		"likes":   likes,
	})
}

func ToggleLikePost(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	userId, userErr := strconv.Atoi(c.FormValue("user-id"))
	postId, postErr := strconv.Atoi(c.FormValue("post-id"))
	if userErr != nil || postErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User or Post ID is not an integer")
	}

	liked := store.LikeStore.ToggleLike(&models.Like{
		UserId: userId,
		PostId: postId,
	})
	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully toggled post like",
		"liked":   liked,
	})
}

func LikedByUser(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	userId, userErr := strconv.Atoi(c.Param("user-id"))
	postId, postErr := strconv.Atoi(c.Param("post-id"))
	if userErr != nil || postErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User or Post ID is not an integer")
	}

	liked := store.LikeStore.LikedByUser(userId, postId)
	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully retrieved liked by user",
		"liked":   liked,
	})
}
