# Script deleted free Elastic IP's across all AWS Regions for your account
*An ElasticIP which is not associated will cost you*

*Assumes you've AWS Access Keys and Secret ID setup via aws configure command in your runtime*

**Steps to run**

>go run main.go

**Steps to build**

>go build main.go

This generates a binary file main which you can then execute

If you want to build for another OS for ex Linux:

>GOOS=linux go build main.go

**Schedule on Lambda**

Can be run/scheduled as a Lambda function using instructions here:
https://docs.aws.amazon.com/lambda/latest/dg/lambda-go-how-to-create-deployment-package.html