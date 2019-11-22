# Script sets retention perdiod on cloudwatch log groups where it is not set already
*for ex when you create a lambda function using AWS Console, no expiration is set on corresponding cloudwatch logs*


*Assumes you've AWS Access Keys and Secret ID setup via aws configure command in your runtime*

**Steps to run**

>go run main.go

**Steps to build**

>go build main.go

This generates a binary file main which you can then execute

If you want to build for another OS for ex Linux:

>GOOS=linux go build main.go

*On Line 39 in file is a variable days, default value is 7, you can change it if you want to set some other expiry period*
