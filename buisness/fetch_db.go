package buisness

import (
	"fmt"
	"go-gin-docker/constants"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ReadDB() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials(constants.AWSAccessKeyID, constants.AWSSecretAccessKey, ""),
	})
	if err != nil {
		fmt.Println("connection issue", err)
	}

	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("transaction"),
		Key: map[string]*dynamodb.AttributeValue{
			"ClientCode": {
				S: aws.String("12"),
			},
			"TransactionID": {
				S: aws.String("12"),
			},
		},
	})

	if err != nil {
		fmt.Println("error fetching items", err)
	}

	fmt.Println(result)
}
