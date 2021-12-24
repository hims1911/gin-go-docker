package buisness

import (
	"fmt"
	"go-gin-docker/constants"
	"go-gin-docker/models"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func CreateTable() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials(constants.AWSAccessKeyID, constants.AWSSecretAccessKey, ""),
	})
	if err != nil {
		fmt.Println("connection issue", err)
	}

	svc := dynamodb.New(sess)

	item := &models.Item{
		Year:   2015,
		Title:  "The Big New Movie",
		Plot:   "Nothing happens at all.",
		Rating: 0.0,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got Error While Marshalling")
	}

	tableName := "Movies"

	// input := &dynamodb.PutItemInput{
	// 	Item:      av,
	// 	TableName: aws.String(tableName),
	// }

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	operation, err := svc.PutItem(input)
	if err != nil {
		log.Fatal("Error Occured", err)
	}
	log.Fatal(operation)

	year := strconv.Itoa(item.Year)
	fmt.Println(year)

	fmt.Println("Successfully added '" + item.Title + "' (" + year + ") to table " + tableName)
}
