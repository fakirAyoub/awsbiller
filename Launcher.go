package main

import (
	"flag"
	"fmt"
	"os"
)
type arrayFlags []string

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}
func (i *arrayFlags) String() string {
	return ""
}
var metricsArgArray arrayFlags

func main() {
	typeArg := flag.String("costType", "", "Choose between the actual cost of your services or a forecast")
	metricArg := flag.String("metric", "", "Metrics to get from the report")
	flag.Var(&metricsArgArray, "metrics", "Some description for this param.")

	granularityArg := flag.String("granularity", "", "Granularity - MONTHLY, DAILY or YEARLY")
	startDateArg := flag.String("startDate", "", "Start Date to be used for the date range to get costs")
	endDateArg := flag.String("endDate", "", "End Date to be used for the date range to get costs")

	flag.Parse()

	if *typeArg == "costExplorer" {
		fmt.Println("Cost Explorer: ")
		costExplorer(granularityArg, startDateArg, endDateArg, metricsArgArray)
	} else if *typeArg == "costForecast" {
		fmt.Println("Cost Forecast: ")
		costForecast(granularityArg, startDateArg, endDateArg, metricArg)
	} else {
		exitErrorf("Bad Argument!, %v", *typeArg)
	}

}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
