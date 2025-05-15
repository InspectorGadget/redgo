package main

import (
	"fmt"

	"github.com/InspectorGadget/redgo/client"
	"github.com/InspectorGadget/redgo/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	client := client.GetRedisClient()
	defer client.Close()
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.Index)
	r.POST("/set", controllers.Set)
	r.GET("/get/:key", controllers.Get)

	if err := r.Run(); err != nil {
		fmt.Printf("Error has occurred: %v", err.Error())
	}
}
