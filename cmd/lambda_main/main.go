package main

import (
	"context"

	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	fiberApp := app.Build()

	fiberLambda = fiberadapter.New(fiberApp)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {	
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
