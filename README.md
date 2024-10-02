# jantools-api-v2
jantools-v2用のAPI

## TeckStack
* Golang
* Echo
* AWS SAM
* AWS Lambda
* AWS APIGateway
* AWS Dynamodb

## Local
```
cd scripts
docker compose up
```
jantools-api-v2
> http://localhost:8080/

dynamodb-admin
> http://localhost:8001/

## Deploy
dev
```
sam build
sam deploy --config-env dev
```
prod
```
sam build
sam deploy --config-env prod
```