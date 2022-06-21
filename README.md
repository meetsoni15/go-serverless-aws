# go-serverless-aws
Serverless AWS go implementation

Make sure `serverless` is installed. [See installation guide](https://serverless.com/framework/docs/providers/openwhisk/guide/installation/).

## 1. Clone this repository
`git clone https://github.com/meetsoni15/go-serverless-aws`

## 2. Change Database credentials in serverless.yml
```
DB_HOST: AWS_DB_HOST
DB_PORT: AWS_DB_PORT
DB_NAME: AWS_DB_NAME
DB_USER: AWS_DB_USER
DB_PASSWORD: AWS_DB_PASSWORD
```

## 3. Deploy
`make deploy`

```
Serverless: Packaging service...
Serverless: Compiling Functions...
Serverless: Compiling API Gateway definitions...
Serverless: Compiling Rules...
Serverless: Compiling Triggers & Feeds...
Serverless: Deploying Functions...
Serverless: Deployment successful!
```

## 4. Testing endpoints

```
sls invoke -f getUserRole --data '{ "pathParameters": {"id":"11"}}'  
```
## 5. Run endpoint locally

```
sls invoke local -f getUserRole --data '{ "pathParameters": {"id":"11"}}'                                                                        took 6s at 12:31:27

START RequestId: e8521206-8c52-1724-fb39-1e76c65b5a72 Version: $LATEST
2022/06/21 07:01:57 Connection to database successful
END RequestId: e8521206-8c52-1724-fb39-1e76c65b5a72
REPORT RequestId: e8521206-8c52-1724-fb39-1e76c65b5a72  Init Duration: 78.80 ms Duration: 3643.27 ms    Billed Duration: 3644 ms        Memory Size: 1024 MB    Max Memory Used: 23 MB

{"statusCode":200,"headers":{"Content-Type":"application/json"},"multiValueHeaders":null,"body":"{\"id\":11,\"role_name\":\"Admin New\",\"created_at\":\"2022-06-20T12:18:04.427649Z\",\"updated_at\":\"2022-06-20T12:32:40.968316Z\"}"}

```