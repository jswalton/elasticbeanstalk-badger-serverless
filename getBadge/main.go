package main

import (
    "fmt"
    "net/http"
    "io/ioutil"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    // "github.com/aws/aws-sdk-go/service/elasticbeanstalk"
    // "github.com/aws/aws-sdk-go/aws/session"
    // "github.com/aws/aws-sdk-go/aws"
    // "github.com/aws/aws-sdk-go/aws/awserr"

)


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  // sess := session.New()
  // fmt.Println("region is: ", *sess.Config.Region)
  // svc := elasticbeanstalk.New(sess)
  // input := &elasticbeanstalk.DescribeEnvironmentsInput{
  //     EnvironmentIds: []*string{
  //         aws.String(request.PathParameters["environmentId"]),
  //     },
  // }
  // fmt.Println("Received Input: ", input)
  // result, err := svc.DescribeEnvironments(input)
  // if err != nil {
  //     if aerr, ok := err.(awserr.Error); ok {
  //         switch aerr.Code() {
  //         default:
  //             fmt.Println(aerr.Error())
  //         }
  //     } else {
  //         // Print the error, cast err to awserr.Error to get the Code and
  //         // Message from an error.
  //         fmt.Println(err.Error())
  //     }
  //
  //     return events.APIGatewayProxyResponse{Body: "Error describing env", StatusCode: 400}, nil
  // }
  // // envName := result.environments[0].EnvironmentName //.gsub! '-', '_'
	// // envStatus := result.environments[0].Status //.downcase.capitalize
	// // envHealth := result.environments[0].Health //.downcase
	// // envAppVersion := result.environments[0].VersionLabel[0,10]

  url := fmt.Sprintf("https://img.shields.io/badge/bloop-bleep_Version_292-green.svg")

  client := &http.Client{}

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return events.APIGatewayProxyResponse{Body: "Error getting badge for env", StatusCode: 400}, nil
  }


  resp, err := client.Do(req)
  if err != nil {
    return events.APIGatewayProxyResponse{Body: "Error getting badge for env", StatusCode: 400}, nil
  }
  defer resp.Body.Close()

  bodyBytes, err2 := ioutil.ReadAll(resp.Body)
  if err2 != nil {
    return events.APIGatewayProxyResponse{Body: "Error getting badge for env", StatusCode: 400}, nil
  }
  headers := map[string]string{ "content-type": "image/svg+xml"}
	return events.APIGatewayProxyResponse{Body: string(bodyBytes), StatusCode: 200, Headers: headers}, nil
}

func main() {
	lambda.Start(Handler)
}
