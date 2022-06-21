BIN_DIR = bin

project := create_user_role update_user_role delete_user_role get_user_role

all: $(project) ## Generate build

create_user_role: $@ ## Generate build for create user role
update_user_role: $@ ## Generate build for update user role
delete_user_role: $@ ## Generate build for delete user role
get_user_role: $@ ## Generate build for get user role

$(project):
	go mod tidy
	go mod vendor
	env GOOS=linux go build -ldflags="-s -w" -o ${BIN_DIR}/$@ api/$@/main.go

rebuild: clean all ## Rebuild the whole project

clean:
	rm -rf ./bin ./vendor

deploy: rebuild
	sls deploy --verbose
