# Description

A set of go scripts allowing you to get actual costs and forecasts from the AWS Cost Explorer API. Doc: https://docs.aws.amazon.com/aws-cost-management/

# USAGE:

## Cost Explorer to get actual Cost and Usage metrics:
```sh
$ Launcher.go --startDate yyyy-MM-dd --endDate yyyy-MM-dd --granularity=DAILY/MONTHLY/YEARLY --costType costExplorer --metrics AMORTIZED_COST --metrics <another_metric_here>
```

## Cost Explorer to get forecast metrics:
```sh
$ Launcher.go --startDate 2020-10-10 --endDate 2020-10-11 --granularity=DAILY --costType costExplorer --metric AMORTIZED_COST
```

## List of Metrics can be found here:


* [Cost and Usage](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostAndUsage.html): Cost and Usage Metric Values
* [Cost Forecast](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostForecast.html): Cost Forecast Metric Values
