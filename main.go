package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sanda0/puller/stucts"
	"github.com/sanda0/puller/util"
)

func handleWebHook(c *gin.Context) {
	secretToken := c.GetHeader("X-Gitlab-Token")
	fmt.Println("secret token: ", secretToken)
	payload := stucts.GitLabWebhookPayload{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(util.PrettyPrintStruct(payload))
}

func main() {
	server := gin.Default()

	server.POST("/puller", handleWebHook)

	server.Run(":8080")
}
