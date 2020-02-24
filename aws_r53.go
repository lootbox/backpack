package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"log"
)

func getIPByDomainName(address string) []string {
	cnames := listCNAMES()

	if cnames == nil {
		log.Println("NO CNAMES")
		return nil
	}

	for _, all := range cnames.ResourceRecordSets {
		if *all.Name == address {
			var res []string
			for _, addr := range all.ResourceRecords {
				res = append(res, *addr.Value)
			}
			return res
		}
	}

	log.Println("NOT FOUND")
	return nil
}

func getDomainNameByIP(address string) string {
	cnames := listCNAMES()

	if cnames == nil {
		log.Println("NO CNAMES")
		return ""
	}

	for _, all := range cnames.ResourceRecordSets {
		for _, cname := range all.ResourceRecords {
			if address == *cname.Value {
				return *all.Name
			}
		}
	}

	log.Println("NOT FOUND")
	return ""
}

func listCNAMES() *route53.ListResourceRecordSetsOutput {
	awsConfig := newAWSCfg()
	listParams := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(awsConfig.R53HostedZoneId),
	}
	result, err := getR53svc().ListResourceRecordSets(listParams)
	if err != nil {
		return nil
	}
	return result
}
