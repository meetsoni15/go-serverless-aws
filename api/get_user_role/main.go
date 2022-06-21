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

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var UserRole model.UserRole
	if request.PathParameters["id"] == "" {
		return utils.RespondError(422, "Invalid parameters"), fmt.Errorf("Invalid parameters")
	}

	UserRole.ID, _ = strconv.Atoi(request.PathParameters["id"])
	db := database.InitDB()
	err := db.Model(&UserRole).WherePK().Select()
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			log.Println("UserRole not found")
			return utils.RespondError(404, "User role not found"), err
		} else {
			log.Printf("Error while getting user role by id %v", err)
			return utils.InternalServerError(), err
		}
	}

	resp := utils.RespondWithData(UserRole, 200)
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
