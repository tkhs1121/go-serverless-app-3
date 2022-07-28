package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type db struct {
	ddb *dynamodb.DynamoDB
}

func newDynamoDB() (*db, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	return &db{ddb: dynamodb.New(sess)}, nil
}

func (db *db) putAmazonURL(url string) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("table"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(getRandID()),
			},
			"time": {
				N: aws.String(getEpochTime()),
			},
			"url": {
				S: aws.String(url),
			},
		},
	}
	if _, err := db.ddb.PutItem(input); err != nil {
		return err
	}
	return nil
}
