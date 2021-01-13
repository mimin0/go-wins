package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
)

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	type Expence struct {
		RecordID             string `json:"recordId"`
		OperationCatigory    string `json:"operationCatigory"`
		OperationDate        string `json:"operationDate"`
		OperationDescription string `json:"operationDescription"`
		OperationSumm        string `json:"operationSumm"`
		// Info              interface{} `json:"info,omitempty"`
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	input := &dynamodb.ScanInput{
		// Key:       key,
		TableName: aws.String("budgetBot"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var obj []Expence
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &obj)
	if err != nil {
		fmt.Errorf("failed to unmarshal Query result items, %v", err)
	}

	// fmt.Printf("record %+v", obj)
	for _, i := range obj {
		fmt.Println(i.RecordID + "," + i.OperationDate + "," + i.OperationCatigory + "," + i.OperationSumm + "," + i.OperationDescription)
	}

}
