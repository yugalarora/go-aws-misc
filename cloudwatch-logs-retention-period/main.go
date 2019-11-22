package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

var listtoupdate []string

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := cloudwatchlogs.New(sess)
	recurse(svc, nil)
}

func recurse(svc *cloudwatchlogs.CloudWatchLogs, next *string) {
	resp, err := svc.DescribeLogGroups(&cloudwatchlogs.DescribeLogGroupsInput{NextToken: next})
	for _, v := range resp.LogGroups {
		if v.RetentionInDays != nil {
			// fmt.Printf("%v\t:%v\n", *v.LogGroupName, *v.RetentionInDays)
		} else {
			listtoupdate = append(listtoupdate, *v.LogGroupName)
			// fmt.Printf("%v\t:%v\n", *v.LogGroupName, nil)
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	if resp.NextToken != nil {
		nt := resp.NextToken
		recurse(svc, nt)
	} else {
		var days int64
		// days to set retention for
		days = 7
		setexpiry(svc, listtoupdate, days)
	}

}

func setexpiry(svc *cloudwatchlogs.CloudWatchLogs, listtoupdate []string, days int64) {
	fmt.Println("Total items without RetentionPeriod:", len(listtoupdate))
	for _, v := range listtoupdate {
		svc.PutRetentionPolicy(&cloudwatchlogs.PutRetentionPolicyInput{LogGroupName: &v, RetentionInDays: &days})
		fmt.Println(v)
	}
}
