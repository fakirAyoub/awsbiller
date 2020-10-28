package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

func costExplorer(granularity *string, startDate *string, endDate *string, metrics []string) {

	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	svc := costexplorer.New(awsSession)

	result, err := svc.GetCostAndUsage(&costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: startDate,
			End:   endDate,
		},
		Granularity: granularity,
		GroupBy: []*costexplorer.GroupDefinition{
			&costexplorer.GroupDefinition{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
		},
		Metrics: aws.StringSlice(metrics),
	})
	if err != nil {
		fmt.Println("Couldn't retrieve metrics: %v", err)
	}

	fmt.Println("Cost Report:", result.ResultsByTime)
}