package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := ec2.New(sess)
	resp, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range resp.Regions {
		region := *v.RegionName
		svc = ec2.New(sess, aws.NewConfig().WithRegion(region))
		resp, err := svc.DescribeAddresses(&ec2.DescribeAddressesInput{})
		if err != nil {
			fmt.Println(err)
		}
		if resp.Addresses != nil {
			for _, v := range resp.Addresses {
				if v.AssociationId == nil {
					fmt.Println(*v.PublicIp)
					svc.ReleaseAddress(&ec2.ReleaseAddressInput{PublicIp: v.PublicIp})
				}
			}
		}
	}

}
