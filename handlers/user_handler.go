package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
	"github.com/ivanwang123/go-blog/models"
	"github.com/ivanwang123/go-blog/stores"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Res map[string]interface{}

func Login(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	email := c.FormValue("email")
	password := c.FormValue("password")

	user, userErr := store.UserStore.UserByEmail(email)
	if userErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect email")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	signedToken, err := token.SignedString([]byte("SECRET123"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to login: %v", err))
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = signedToken
	cookie.MaxAge = 3600 * 24
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully logged in",
	})
}

func Logout(c echo.Context) error {
	cookie, err := c.Cookie("auth_token")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to logout: %v", err))
	}
	cookie.MaxAge = -1

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully logged out",
	})
}

func Me(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))

	me, err := store.UserStore.User(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to get current user: %v", err))
	}

	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully retrieved current user",
		"user":    me,
	})
}

func ListUsers(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	users, err := store.UserStore.Users()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Unable to retrieve users: %v", err))
	}

	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully retrieved users",
		"users":   users,
	})
}

func GetUser(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not an integer")
	}

	user, err := store.UserStore.User(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
	}

	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully retrieved user",
		"user":    user,
	})
}

func CreateUser(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm-password")

	if err := validation.Validate(password, validation.Required, validation.Length(8, 255)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid password: %v", err))
	}

	if password != confirmPassword {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Confirm password does not match password"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to create user: %v", err))
	}

	validate := validator.New()
	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hash),
	}
	validateErr := validate.Struct(user)
	if validateErr != nil {
		return c.JSON(http.StatusBadRequest, &Res{
			"message": "Unable to create user",
			"errors":  strings.Split(validateErr.(validator.ValidationErrors).Error(), "\n"),
		})
	}

	if err := store.UserStore.CreateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to create user: %v", err))
	}
	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully created user",
	})
}

func UpdateUser(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not an integer")
	}
	username := c.FormValue("username")
	email := c.FormValue("email")

	validate := validator.New()
	user := &models.User{
		Id:       id,
		Username: username,
		Email:    email,
	}
	validateErr := validate.Struct(user)
	if validateErr != nil {
		return c.JSON(http.StatusBadRequest, &Res{
			"message": "Unable to update user",
			"errors":  strings.Split(validateErr.(validator.ValidationErrors).Error(), "\n"),
		})
	}

	if err := store.UserStore.UpdateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to update user: %v", err))
	}
	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully updated user",
	})
}

func DeleteUser(c echo.Context) error {
	store := c.Get("store").(*stores.Store)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not an integer")
	}

	if err := store.UserStore.DeleteUser(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to delete user: %v", err))
	}
	return c.JSON(http.StatusOK, &Res{
		"message": "Successfully deleted user",
	})
}
