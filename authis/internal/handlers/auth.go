package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"authis/internal/database"
	"authis/internal/models"
	"authis/internal/shared"
)

var (
	ctx = context.Background()
)

func Register(c *fiber.Ctx) error {
	data := map[string]string{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Username: data["username"],
		Password: shared.HashPassword(data["password"]),
	}

	if err := user.Create(); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	data := map[string]string{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{}
	if err := user.FindByUsername(data["username"]); err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := shared.ComparePassword(user.Password, data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "wrong password",
		})
	}

	access, err := shared.AccessToken(user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	refresh, err := shared.RefreshToken(user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	rdb := database.GetRedisClient()

	rdb.Set(strconv.Itoa(int(user.ID)), *access, time.Hour*24)
	rdb.Set("refresh_"+strconv.Itoa(int(user.ID)), *refresh, time.Hour*24*7)

	return c.JSON(fiber.Map{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func User(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	rdb := database.GetRedisClient()

	rdb.Del(strconv.Itoa(int(user.ID)))
	rdb.Del("refresh_" + strconv.Itoa(int(user.ID)))

	return c.JSON(fiber.Map{
		"message": "Successful",
	})
}

func Refresh(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	access, err := shared.AccessToken(user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	refresh, err := shared.RefreshToken(user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	rdb := database.GetRedisClient()

	rdb.Set(strconv.Itoa(int(user.ID)), *access, time.Hour*24)
	rdb.Set("refresh_"+strconv.Itoa(int(user.ID)), *refresh, time.Hour*24*7)

	return c.JSON(fiber.Map{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
