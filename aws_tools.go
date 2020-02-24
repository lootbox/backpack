package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getEC2Filter(filterConfig map[string]string) []*ec2.Filter {
	filters := make([]*ec2.Filter, 0)
	for k, v := range filterConfig {
		filters = append(filters, &ec2.Filter{
			Name: aws.String(k),
			Values: []*string{
				aws.String(v),
			},
		})
	}
	return filters
}
