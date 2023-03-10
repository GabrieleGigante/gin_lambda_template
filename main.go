package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambdaV2

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	if inCloud() {
		gin.SetMode(gin.ReleaseMode)
		ginLambda = ginadapter.NewV2(SetupRuter())
		lambda.Start(Handler)
		return
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":9999"
	}
	SetupRuter().Run(PORT)
}

func inCloud() bool {
	return os.Getenv("LAMBDA_TASK_ROOT") != ""
}
