package buisness

import (
	"fmt"
	"go-gin-docker/constants"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func FetchCondition() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials(constants.AWSAccessKeyID, constants.AWSSecretAccessKey, ""),
	})
	if err != nil {
		fmt.Println("connection issue", err)
	}

	svc := dynamodb.New(sess)

	// Construct the Key condition builder
	keyCond := expression.Key("ClientCode").Equal(expression.Value("12")).And(expression.Key("TransactionID").Equal(expression.Value("12")))

	// Create the project expression builder with a names list.
	proj := expression.NamesList(expression.Name("Transaction"))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).WithProjection(proj).Build()
	if err != nil {
		log.Fatal("error building the expression")
	}

	print(expr.Condition())

	input := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("transaction"),
	}

	fmt.Println(input)

	result, err := svc.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}
