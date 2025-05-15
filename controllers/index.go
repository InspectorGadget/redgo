package controllers

import (
	"fmt"
	"net/http"

	"github.com/InspectorGadget/redgo/client"
	"github.com/InspectorGadget/redgo/structs"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		structs.Response{
			"message": "All good!",
		},
	)
}

func Set(c *gin.Context) {
	if c.ContentType() != "application/json" {
		c.JSON(
			http.StatusBadRequest,
			structs.Response{
				"error": "Content-Type must be application/json",
			},
		)
		return
	}

	redisValue := &structs.RedisValue{}
	err := c.ShouldBind(redisValue)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			structs.Response{
				"error": err.Error(),
			},
		)
		return
	}

	client := client.GetRedisClient()
	err = SetValue(client, redisValue.GetKey(), redisValue.GetValue())
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			structs.Response{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		structs.Response{
			"message": "Stored in Redis",
			"key":     redisValue.GetKey(),
			"value":   redisValue.GetValue(),
		},
	)
}

func Get(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(
			http.StatusBadRequest,
			structs.Response{
				"error": "Key is required",
			},
		)
		return
	}

	client := client.GetRedisClient()
	value, err := GetValue(client, key)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			structs.Response{
				"error": fmt.Sprintf("Key '%v' is not found", key),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		structs.Response{
			"message": "Successfully fetched the key + value from Redis",
			"key":     key,
			"value":   value,
		},
	)
}
