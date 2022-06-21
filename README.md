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