package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/service/elasticbeanstalk"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"

)


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  //create new session for ElasticBeanstalk SDK and get environment descriptions for requested env
  sess := session.New()
  fmt.Println("region is: ", *sess.Config.Region)
  svc := elasticbeanstalk.New(sess)
  input := &elasticbeanstalk.DescribeEnvironmentsInput{
      EnvironmentIds: []*string{
          aws.String(request.PathParameters["environmentId"]),
      },
  }
  //check the result for errors
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

      return events.APIGatewayProxyResponse{Body: "Error describing env", StatusCode: 400}, nil
  }

  //defaults for EB env variables
  envName := "<nil>"
  envStatus := "<nil>"
  envHealth := "<nil>"
  envAppVersion := "<nil>"

  //if there are environments returned, go ahead and assign the information to our EB env variables
  if len(result.Environments) > 0 {
    if result.Environments[0].EnvironmentName  != nil { // Check for nil!
      envName = strings.Replace(*result.Environments[0].EnvironmentName, "-", "_", -1)
    }
    if result.Environments[0].Status  != nil { // Check for nil!
      envStatus = *result.Environments[0].Status
    }
    if result.Environments[0].Health  != nil { // Check for nil!
      envHealth = strings.ToLower(*result.Environments[0].Health)
    }
    if result.Environments[0].VersionLabel  != nil { // Check for nil!
      envAppVersion = *result.Environments[0].VersionLabel
    }
  } else {
    return events.APIGatewayProxyResponse{Body: "No Environment Found", StatusCode: 400}, nil
  }

  //Go get the badge from img.shields.io
  url := fmt.Sprintf("http://img.shields.io/badge/" + envName + "-" +envStatus+ "_Version_" + envAppVersion + "-" + envHealth + ".svg")
  fmt.Println("url will be: ", url)

  client := &http.Client{}

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return events.APIGatewayProxyResponse{Body: "Error getting badge for env", StatusCode: 400}, nil
  }

  //check the img.shields.io response
  resp, err := client.Do(req)
  if err != nil {
    return events.APIGatewayProxyResponse{Body: "Error getting badge for env", StatusCode: 400}, nil
  }
  defer resp.Body.Close()

  //get the body bytes (svg) to be returned
  bodyBytes, err2 := ioutil.ReadAll(resp.Body)
  if err2 != nil {
    return events.APIGatewayProxyResponse{Body: "Error getting badge for env", StatusCode: 400}, nil
  }

  //create headers for content type
  headers := map[string]string{ "content-type": "image/svg+xml"}

  //return the APIGatewayProxyRequest with the svg badge image as the body
	return events.APIGatewayProxyResponse{Body: string(bodyBytes), StatusCode: 200, Headers: headers}, nil
}

func main() {
	lambda.Start(Handler)
}
