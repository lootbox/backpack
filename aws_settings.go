package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/spf13/viper"
)

type awsCfg struct {
	AwsAccessKey        string
	AwsSecretKey        string
	AwsRegion           string
	AwsAvailabilityZone string
	//	ROUTE 53
	R53HostedZoneId  string
	R53BaseNamespace string
	R53BaseDomain    string
	// EC 2
	Ec2SubnetId string
	Ec2SgIds    string
	//	S3
	S3BucketName   string
	S3MetadataPath string
}

func loadEC2Filters() {
	viper.SetDefault("LatestRedHatAMIFilter", map[string]string{
		"state":               "available",
		"root-device-type":    "ebs",
		"image-type":          "machine",
		"virtualization-type": "hvm",
		"name":                "amzn1*",
	})

}

func newAWSCfg() *awsCfg {
	cfg := &awsCfg{
		AwsAccessKey:        getEnv("AWS_ACCESS_KEY", ""),
		AwsSecretKey:        getEnv("AWS_SECRET_KEY", ""),
		AwsRegion:           getEnv("AWS_REGION", "eu-central-1"),
		AwsAvailabilityZone: getEnv("AWS_AVAILABILITY_ZONE", "eu-central-1a"),
		R53HostedZoneId:     getEnv("R53_HOSTED_ZONE_ID", ""),
		R53BaseNamespace:    getEnv("R53_BASE_NAMESPACE", ""),
		R53BaseDomain:       getEnv("R53_BASE_DOMAIN", ""),
		Ec2SubnetId:         getEnv("EC2_SUBNET_ID", ""),
		Ec2SgIds:            getEnv("EC2_SG_IDs", ""),
		S3BucketName:        getEnv("S3_BUCKET_NAME", ""),
		S3MetadataPath:      getEnv("S3_METADATA_PATH", ""),
	}
	cfg.R53BaseDomain = cfg.R53BaseNamespace + cfg.R53BaseDomain
	return cfg
}

func getEC2svc() *ec2.EC2 {
	awsConfig := newAWSCfg()
	Session := session.Must(session.NewSession())
	return ec2.New(
		Session,
		aws.NewConfig().WithRegion(awsConfig.AwsRegion).WithCredentials(
			credentials.NewStaticCredentials(
				awsConfig.AwsAccessKey,
				awsConfig.AwsSecretKey,
				"",
			),
		),
	)
}

func getR53svc() *route53.Route53 {
	awsConfig := newAWSCfg()
	Session := session.Must(session.NewSession())
	return route53.New(
		Session,
		aws.NewConfig().WithRegion(awsConfig.AwsRegion).WithCredentials(
			credentials.NewStaticCredentials(
				awsConfig.AwsAccessKey,
				awsConfig.AwsSecretKey,
				"",
			),
		),
	)
}
