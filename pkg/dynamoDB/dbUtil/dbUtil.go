package dbUtil

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetDBInstance() *dynamodb.DynamoDB {
	AnonymousCredentials := credentials.NewStaticCredentials("bdh", "abcd", "")
	sess := session.Must(session.NewSession(&aws.Config{
		//Endpoint: aws.String("http://localhost:8000"),
		Endpoint: aws.String("http://172.19.0.2:8000"),
		//Endpoint: aws.String("http://localhost:8000"),
		Region: aws.String("us-east-1"),
		Credentials: AnonymousCredentials,
	}))
	return dynamodb.New(sess)
}