package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-pg/pg/v10"
	"github.com/meetsoni15/go-serverless-aws/database"
	"github.com/meetsoni15/go-serverless-aws/model"
	"github.com/meetsoni15/go-serverless-aws/utils"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//Data model of user role
	var UserRole model.UserRole
	//Decode reeuqest body to user role data model
	if request.PathParameters["id"] == "" {
		return utils.RespondError(422, "Invalid parameters"), fmt.Errorf("Invalid parameters")
	}

	//Get user role id from path parameters
	UserRole.ID, _ = strconv.Atoi(request.PathParameters["id"])
	//Connect to database
	db := database.InitDB()
	//Close database connection
	defer db.Close()
	//Delete user role from database
	err := db.Model(&UserRole).WherePK().Select()
	if err != nil {
		//Check if user role not found
		if err.Error() == pg.ErrNoRows.Error() {
			log.Println("UserRole not found")
			return utils.RespondError(404, "User role not found"), err
		} else {
			log.Printf("Error while getting user role by id %v", err)
			return utils.InternalServerError(), err
		}
	}
	//Return response
	resp := utils.RespondWithData(UserRole, 200)
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
