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

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//Data model of user role
	var newUserRole model.UserRole
	//Decode reeuqest body to user role data model
	err := json.Unmarshal([]byte(request.Body), &newUserRole)
	if err != nil {
		return utils.RespondError(400, "Invalid json"), err
	}

	//Check if user role id is empty
	if request.PathParameters["id"] == "" || newUserRole.Name == "" {
		return utils.RespondError(422, "Invalid parameters"), err
	}

	//Get user role id from path parameters
	newUserRole.ID, _ = strconv.Atoi(request.PathParameters["id"])
	//Connect to database
	db := database.InitDB()
	//Close database connection
	defer db.Close()
	//Update user role in database
	newUserRole.UpdatedAt = time.Now()
	_, err = db.Model(&newUserRole).WherePK().UpdateNotZero(&newUserRole)
	if err != nil {
		log.Printf("Error while creating new user role %v", err)
		return utils.InternalServerError(), err
	}
	//Return response
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
