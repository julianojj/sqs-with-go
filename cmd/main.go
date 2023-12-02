package main

import (
	"app/internal/adapters"
	"app/internal/application/usecases"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	sqs := &adapters.Sqs{}
	err := sqs.Connect()
	if err != nil {
		panic("error to connect to sqs")
	}
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil)).With(
		"service", os.Getenv("SERVICE_NAME_APPLICATION"),
		"environment", os.Getenv("SERVICE_ENV"),
	)
	createUser := usecases.NewCreateUser(sqs, logger)
	router := gin.Default()
	router.POST("/create_user", func(ctx *gin.Context) {
		var input *usecases.CreateUserInput
		err := ctx.ShouldBindJSON(&input)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
				"message": err.Error(),
				"code":    http.StatusBadRequest,
			})
			return
		}
		err = createUser.Execute(input)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]any{
				"message": err.Error(),
				"code":    http.StatusUnprocessableEntity,
			})
			return
		}
		ctx.JSON(http.StatusCreated, nil)
	})
	router.Run(":8080")
}
