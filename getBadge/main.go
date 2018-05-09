package main

import (
    "fmt"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/service/elasticbeanstalk"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"

)


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  sess := session.New()
  fmt.Println("region is: ", *sess.Config.Region)
  svc := elasticbeanstalk.New(sess)
  input := &elasticbeanstalk.DescribeEnvironmentsInput{
      EnvironmentIds: []*string{
          aws.String(request.PathParameters["environmentId"]),
      },
  }
  fmt.Println("Received Input: ", input)
  result, err := svc.DescribeEnvironments(input)
  if err != nil {
      if aerr, ok := err.(awserr.Error); ok {
          switch aerr.Code() {
          default:
              fmt.Println(aerr.Error())
          }
      } else {
          // Print the error, cast err to awserr.Error to get the Code and
          // Message from an error.
          fmt.Println(err.Error())
      }
      return events.APIGatewayProxyResponse{Body:  "that didnt work for now", StatusCode: 200}, nil
  }

  fmt.Println("result is:", result)
	return events.APIGatewayProxyResponse{Body:  "thatll work for now", StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
