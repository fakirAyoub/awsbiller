# Description

A set of go scripts allowing you to get actual costs and forecasts from the AWS Cost Explorer API. Doc: https://docs.aws.amazon.com/aws-cost-management/

This program supposes that the credentials of the account you want to get metrics from are already the default ones in your environment, see: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html

# USAGE:
First, you need to ` go get github.com/aws/aws-sdk-go/`
## Cost Explorer to get actual Cost and Usage metrics:
```sh
$ Launcher.go --startDate yyyy-MM-dd --endDate yyyy-MM-dd --granularity=DAILY/MONTHLY/YEARLY --costType costExplorer --metrics AMORTIZED_COST --metrics <another_metric_here>
```

## Cost Forecast to get forecast metrics:
```sh
$ Launcher.go --startDate yyyy-MM-dd --endDate yyyy-MM-dd --granularity=DAILY/MONTHLY/YEARLY --costType costForecast --metric AMORTIZED_COST
```

## List of Metrics can be found here:


* [Cost and Usage](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostAndUsage.html)
* [Cost Forecast](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostForecast.html)
