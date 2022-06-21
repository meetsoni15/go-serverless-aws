package main

import (
	"encoding/json"
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
	err := json.Unmarshal([]byte(request.Body), &newUserRole)
	if err != nil {
		return utils.RespondError(400, "Invalid json"), err
	}

	if request.PathParameters["id"] == "" || newUserRole.Name == "" {
		return utils.RespondError(422, "Invalid parameters"), err
	}

	newUserRole.ID, _ = strconv.Atoi(request.PathParameters["id"])
	db := database.InitDB()
	newUserRole.UpdatedAt = time.Now()
	_, err = db.Model(&newUserRole).WherePK().UpdateNotZero(&newUserRole)
	if err != nil {
		log.Printf("Error while creating new user role %v", err)
		return utils.InternalServerError(), err
	}

	resp := events.APIGatewayProxyResponse{
		StatusCode:      204,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
