package main

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/viper"
	"time"
)

func getEC2InstanceByName(name string) *ec2.DescribeInstancesOutput {
	viper.Set("GetInstanceByIP", map[string]string{
		"root-device-type": "ebs",
		"network-interface.addresses.private-ip-address": getIPByDomainName(name)[0], // TODO
	})

	input := &ec2.DescribeInstancesInput{
		Filters: getEC2Filter(viper.GetStringMapString("GetInstanceByIP")),
	}

	result, err := getEC2svc().DescribeInstances(input)

	if err != nil {
		return nil
	}

	return result
}

func getLatestAMI(filterKey string) *ec2.Image {
	all := filterAMI(filterKey)

	if all == nil {
		return nil
	}

	latest, _ := time.Parse(time.RFC3339, "1970-01-01T01:23:00.000Z")
	var result *ec2.Image

	for _, image := range all.Images {
		_date, _ := time.Parse(time.RFC3339, *image.CreationDate)

		if _date.After(latest) {
			latest = _date
			result = image
		}
	}

	return result
}

func filterAMI(filterKey string) *ec2.DescribeImagesOutput {
	loadEC2Filters()

	input := &ec2.DescribeImagesInput{
		Filters: getEC2Filter(viper.GetStringMapString(filterKey)),
	}

	result, err := getEC2svc().DescribeImages(input)

	if err != nil {
		return nil
	}

	return result
}
