package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

func costForecast(granularity *string, startDate *string, endDate *string, metric *string) {


	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	svc := costexplorer.New(awsSession)

	result, err := svc.GetCostForecast(&costexplorer.GetCostForecastInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: startDate,
			End:   endDate,
		},
		Granularity: granularity,
		Metric: metric,
	})
	if err != nil {
		fmt.Println("Couldn't retrieve metrics: %v", err)
	}

	fmt.Println("Cost Report:", result.ForecastResultsByTime)
}