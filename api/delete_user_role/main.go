package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/meetsoni15/go-serverless-aws/database"
	"github.com/meetsoni15/go-serverless-aws/model"
	"github.com/meetsoni15/go-serverless-aws/utils"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var newUserRole model.UserRole
	if request.PathParameters["id"] == "" {
		return utils.RespondError(422, "Invalid parameters"), fmt.Errorf("Invalid parameters")
	}

	newUserRole.ID, _ = strconv.Atoi(request.PathParameters["id"])
	db := database.InitDB()
	newUserRole.UpdatedAt = time.Now()
	_, err := db.Model(&newUserRole).WherePK().Delete()
	if err != nil {
		log.Printf("Error while deleting new user role %v", err)
		return utils.InternalServerError(), err
	}

	resp := events.APIGatewayProxyResponse{
		StatusCode: 204,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
