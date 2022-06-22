package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/meetsoni15/go-serverless-aws/database"
	"github.com/meetsoni15/go-serverless-aws/model"
	"github.com/meetsoni15/go-serverless-aws/utils"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//Data model of user role
	var newUserRole model.UserRole
	//Decode reeuqest body to user role data model
	err := json.Unmarshal([]byte(request.Body), &newUserRole)
	//If error while decoding, return error
	if err != nil {
		return utils.RespondError(400, "Invalid json"), err
	}

	//Connect to database
	db := database.InitDB()
	//Close database connection
	defer db.Close()
	//Set created at & updated at to current time
	newUserRole.CreatedAt = time.Now()
	newUserRole.UpdatedAt = time.Now()
	//Insert user role to database
	_, err = db.Model(&newUserRole).Insert(&newUserRole)
	//If error while inserting, return error
	if err != nil {
		log.Printf("Error while creating new user role %v", err)
		return utils.InternalServerError(), err
	}

	//Return user role data model in json format
	resp := events.APIGatewayProxyResponse{
		StatusCode:      201,
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
