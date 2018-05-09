Stupid easy badge issuer for AWS Beanstalk Evironments

Issue GitHub badges for your Beanstalk evironments that look like this:

![alt text](http://img.shields.io/badge/Your%20Beanstalk%20Environment-Ready_Version_0.0-brightgreen.svg)

![alt text](http://img.shields.io/badge/Your%20Beanstalk%20Environment-Ready_Version_0.0-yellow.svg)

![alt text](http://img.shields.io/badge/Your%20Beanstalk%20Environment-Terminated_Version_0.0-red.svg)

![alt text](http://img.shields.io/badge/Your%20Beanstalk%20Environment-Updating_Version_0.0-lightgrey.svg)


## Running it in AWS
This project uses [serverless framework](https://serverless.com/) to deploy and operate the serverless app.

in order to deploy to your AWS environment, install serverless and run:

```
serverless deploy
```

This will deploy the CloudFormation template created by ```serverless.yml``` and handle provisioning of resources, etc. needed to run the application

After deployment you should see a similar message to the following:

```
Service Information
service: elasticbeanstalk-badger
stage: dev
region: us-east-1
stack: elasticbeanstalk-badger-dev
api keys:
  None
endpoints:
  GET - https://xxxxxxx.execute-api.xx-region-x.amazonaws.com/dev/getBadge/{environmentId}
functions:
  getBadge: elasticbeanstalk-badger-dev-getBadge
```

You can now get a badge for your beanstalk environments by sending a GET request to the following url as listed above:

```
https://xxxxxxx.execute-api.xx-region-x.amazonaws.com/dev/getBadge/{environmentId}
```

and replacing `{environmentId}` with the environment id you wish to get a badge for

## Running it locally

Serverless has some features that allow you to run locally.
