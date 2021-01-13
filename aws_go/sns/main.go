package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.ListSubscriptions(nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, s := range result.Subscriptions {
		vArn := *s.SubscriptionArn
		fmt.Println(ServiceNameFromARN(vArn) + "," + *s.TopicArn)
	}
}

// GetServiceFromArn removes the arn:aws: component string of
// the name and returns the first keyword that appears, svc
func ServiceNameFromARN(arn *string) *string {
	shortArn := strings.Replace(*arn, "arn:aws:", "", -1)
	sliced := strings.Split(shortArn, ":")
	return &sliced[0]
}
