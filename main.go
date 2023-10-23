package main

import (
	"bwa/golang/auth"
	"bwa/golang/campaign"
	"bwa/golang/handler"
	"bwa/golang/helper"
	"bwa/golang/payment"
	"bwa/golang/transaction"
	"bwa/golang/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(localhost:3306)/bwastarup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//payment service
	paymentService := payment.NewPaymentService()

	//user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewJwtService()
	userHandler := handler.NewUserHandler(userService, authService)

	//campaigns
	campaignRepository := campaign.NewRepositoryCampaign(db)
	campaignService := campaign.NewCampaignService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	//Transaction
	transactionRepository := transaction.NewTransactionRepository(db)
	transactionService := transaction.NewTransactionService(transactionRepository, campaignRepository, paymentService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	userId, _ := userService.GetUserById(24)
	trs := transaction.CreateTransactionInput{
		CampaignId: 1,
		User:       userId,
		Amount:     100,
	}

	t, err := transactionService.CreateTransaction(trs)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t)
	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	//user
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	//campaigns
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)

	//transaction

	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransaction)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransaction)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)

	router.Run()

}
func authMiddleware(serviceAuth auth.ServiceAuth, service user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {

			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := serviceAuth.ValidateToken(tokenString)

		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			if err != nil {
				response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
		}

		userId := int(claim["user_id"].(float64))

		newUser, err := service.GetUserById(userId)

		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", newUser)
	}

}
